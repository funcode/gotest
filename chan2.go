package main

import "fmt"

var done = make(chan int)
var msg = make(chan int)

func main() {
	i := 0
	go func() {
		defer func() {
			fmt.Println("defer")
		}()
		for {
			m := <-msg
			if m >= 100 {
				fmt.Println("Stop!")
				close(done)
				// for s := range msg {
				// 	fmt.Println(s)
				// }
				return
			}
		}
	}()
	for {
		select {
		case <-done:
			fmt.Println("i am done!")
			fmt.Println(i)
			return
		case msg <- i:
			i = i + 1
		}
	}
}
