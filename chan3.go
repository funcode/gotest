package main

import (
	"fmt"
	"time"
)

type T int

func IsClosed(ch <-chan T) bool {
	return len(ch) > 0
	// select {
	// case <-ch:
	// 	return true
	// default:
	// }

	// return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	go func() {
		c <- 1

		c <- 2
	}()
	fmt.Println(IsClosed(c)) // false
	fmt.Println(IsClosed(c)) // false
	fmt.Println(IsClosed(c)) // false
	fmt.Println(IsClosed(c)) // false
	fmt.Println(IsClosed(c)) // false
	time.Sleep(2 * time.Second)
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
	fmt.Println(IsClosed(c)) // true
}
