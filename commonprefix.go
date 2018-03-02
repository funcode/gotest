package main

import "fmt"

func findPrefix(s1, s2 string) int {
	l := 0
	if len(s1) > len(s2) {
		l = len(s2)
	} else {
		l = len(s1)
	}
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			return i - 1
		}
	}
	return l - 1
}

func main() {
	names := []string{}
	v := 0
	m := 100
	if len(names) < 1 {
		fmt.Println("empty")
		return
	}
	for _, s := range names[1:] {
		v = findPrefix(s, names[0])
		if v < m {
			m = v
		}
		if m < 0 {
			break
		}
	}
	fmt.Println(m)
	var s3 string
	s3 = names[0][0 : m+1]
	fmt.Println(s3)
}
