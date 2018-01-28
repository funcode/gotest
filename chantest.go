package main

import "fmt"
import "time"
import "sync"
import "math/rand"

func main() {
	ws := make(chan int)
	done := make(chan int)
	var wg, wg2 sync.WaitGroup
	defer close(done)

	wg2.Add(1)
	go func(done chan int, ws chan int) {
		defer close(ws)
		rand.Seed(100)
		for {
			select {
			case <-done:
				fmt.Println("   !!! Done received!!!   ")
				wg.Wait()
				wg2.Done()
				return
			default:
				wg.Add(1)
				msg := rand.Intn(50) + 1
				fmt.Println("  Sending data : ", msg)
				ws <- msg
			}
		}
	}(done, ws)

	var m int
	timer := time.NewTimer(time.Duration(1) * time.Millisecond)

READING:
	for {
		select {
		case <-timer.C:
			fmt.Println("   !!! time out !!! ")
			go func(done chan int) { done <- 1 }(done)
			for m = range ws {
				fmt.Println("clearing : ", m)
				wg.Done()
			}
		case m = <-ws:
			fmt.Print(" default get : ", m)
			if m != 0 {
				wg.Done()
			} else if m == 0 {
				break READING
			}
		default:
			// fmt.Println("waiting..")
		}
	}

	wg2.Wait()
	panic(1)
}
