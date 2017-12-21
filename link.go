package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

func buildLink(head *Node) *Node {
	var current *Node
	current = nil
	var x *Node
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		x = new(Node)
		x.Value = i
		x.Next = current
		current = x
		fmt.Println(x)
	}

	head = current
	for v := head; v != nil; v = v.Next {
		fmt.Println(v.Value)
	}
	fmt.Println("head is ", head)
	return head
}

func linkLen(head *Node) int {
	i := 0
	for v := head; v != nil; v = v.Next {
		i++
		if i == 3 {
			v.Next = head
			break
		}
	}
	return i
}

func main() {
	var first *Node

	first = buildLink(first)
	fmt.Println("fist is ", first)
	fmt.Println("Link lenth: ", linkLen(first))

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected!!!")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}
}
