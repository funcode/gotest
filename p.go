package main

import (
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func gen(logger *log.Logger, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			logger.Println("generating numer...", n)
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(done <-chan int, logger *log.Logger, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		logger.Println(".............sq entering loop ", time.Now().UTC())
		for n := range in {
			logger.Println("..............sq inside loop before select: ", time.Now().UTC())
			select {
			case out <- n * n:
				logger.Println("..............sq case out: ", n)
			case <-done:
				logger.Println("..............sq case done: ", time.Now().UTC())
				return
			}
			logger.Println("..............sq inside loop after select: ", time.Now().UTC())
		}
		logger.Println(".............sq out of loop ", time.Now().UTC())
	}()
	return out
}

func merge(done <-chan int, logger *log.Logger, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int, tag int) {
		defer wg.Done()
		logger.Println("**** entering loop: ", time.Now().UTC())
		for n := range c {
			logger.Println("**** inside for loop, before select : ", time.Now().UTC())
			select {
			case out <- n:
				logger.Println("****  case out: ", time.Now().UTC())
			case <-done:
				logger.Println("**** case done: ", time.Now().UTC())
				return
				//			default:
				//				logger.Println("--------select default: ", tag)
			}
			logger.Println("**** inside for loop, after select : ", time.Now().UTC())
		}
		logger.Println("**** out of loop: ", time.Now().UTC())
		//	wg.Done()
		//logger.Println("I'm done ----", tag)
	}
	wg.Add(len(cs))
	index := 0
	for _, c := range cs {
		index = index + 1
		logger.Println("routine : ", index)
		go output(c, index)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	logger := log.New(os.Stdout, "DEBUG", log.Lmicroseconds)
	numbers := make([]int, 100)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}
	in := gen(logger, numbers...)

	done := make(chan int)
	//defer close(done)

	c1 := sq(done, logger, in)

	_ = merge(done, logger, c1)
	//out := merge(done, c1)
	//logger.Println(<-out)

	time.Sleep(1 * time.Second)
	close(done)

	time.Sleep(1 * time.Second)
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	return
}
