package main

import "fmt"
import "time"

func func1() {
	fmt.Println("fun1 started")
	defer func() {
		fmt.Println("recover in func1 : ", recover())
	}()
	//c := make(chan int)
	go fun2()

	for {
		//c <- 1
		fmt.Println("fun1 is running")

		time.Sleep(1 * time.Second)
	}
	fmt.Println("fun1 ended")
}

func fun2() {
	//defer func() { fmt.Println("func2 is recoverd : ", recover()) }()
	fmt.Println("fun2 started")
	//_ = <-c
	time.Sleep(1 * time.Second)
	panic("func2 booom!")

	fmt.Println("fun2 ended")
}

func main() {
	defer func() { fmt.Println("recover in main : ", recover()) }()
	func1()
}
