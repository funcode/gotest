package main

import "fmt"

func main() {
	a := [...]string{"a", "b"}
	for i := 0; i < len(a); i++ {
		fmt.Printf(a[i])
	}

	s := []int{1, 2, 3}
	for _, n := range s {
		fmt.Println(n)
	}

	var t []int
	t = append(t, s...)

	for _, n := range t {
		fmt.Println(n)
	}

}
