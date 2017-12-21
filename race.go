package main

import "fmt"
import "time"

func main() {
	i := 0
	ch := make(chan int)
	go func() {
		i++
		time.Sleep(1 * time.Second)
		ch <- i
	}()
	go func() {
		i++
		ch <- i
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
