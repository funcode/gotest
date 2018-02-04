package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int)
var wg = &sync.WaitGroup{}

func f(id int) {
	defer wg.Done()
	select {
	case ch <- id:
		fmt.Println("send from : ", id)
	case c := <-ch:
		fmt.Println("get from : ", c)
		// default:
		// fmt.Println("default : ", id)
	}
}

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go f(i)
	}
	wg.Wait()

}
