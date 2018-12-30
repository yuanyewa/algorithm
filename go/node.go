package main

import "fmt"

type node struct {
	value interface{}
	next  *node
}

func main() {
	var n1, n2, n3, n4 node
	head := &n1
	n1.value = 1
	n2.value = "A"
	n3.value = 'B'
	n4.value = "This is me"
	n1.next = &n2
	n2.next = &n3
	n3.next = &n4

	// insert in beginning
	n0 := node{0, &n1}
	head = &n0
	fmt.Println(head.next.next)

}
