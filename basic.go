package main

import "fmt"

func my_min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func power(a, b int) int {
	x := 1
	for i := 0; i < b; i++ {
		x = x * a
	}
	return x
}

func main() {
	fmt.Println(power(3, 4))
}
