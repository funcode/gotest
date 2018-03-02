package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	v := "{}[]()"
	ss := make([]rune, 100)
	t := -1
	match := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, x := range s {
		fmt.Printf("%c", x)
		if !strings.ContainsRune(v, x) {
			return false
		}
		if strings.ContainsRune("{[(", x) {
			t = t + 1
			ss[t] = match[x]
		} else {
			if t < 0 {
				return false
			}
			y := ss[t]
			t = t - 1
			if y != x {
				fmt.Printf("\n%c%c\n", y, x)
				return false
			}
		}
	}
	return t == -1
}

func main() {
	test := "({[]})"
	fmt.Println(isValid(test))
	// test = "({}[])"
	// fmt.Println(isValid(test))
	// test = "({[]})"
	// fmt.Println(isValid(test))
	// test = "({[}])"
	// fmt.Println(isValid(test))
	// test = "(){}[]"
	// fmt.Println(isValid(test))
	// test = "({}[]"
	// fmt.Println(isValid(test))
	test = "({})[]]"
	fmt.Println(isValid(test))
	test = "([)]"
	fmt.Println(isValid(test))
}
