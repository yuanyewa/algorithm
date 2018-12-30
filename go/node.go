package main

import "fmt"

type node struct {
	value interface{}
	next  *node
}

func main() {
	var n1, n2, n3, n4 node
	head := &n1
	tail := &n4
	n1.value = 1
	n2.value = "A"
	n3.value = 'B'
	n4.value = "This is me"
	n1.next = &n2
	n2.next = &n3
	n3.next = &n4
	fmt.Println(head)

	// insert in beginning
	n0 := node{0, &n1}
	head = &n0

	// insert after n2
	nn := node{"a new node", nil}
	nn.next = n2.next
	n2.next = &nn

	// insert in end
	ne := node{"This is the new end", nil}
	tail.next = &ne
	tail = &ne

	// remove head
	head = head.next

	// remove tail
	for next := head.next; next != nil; {
		if next.next == tail {
			tail = next
			break
		}
		next = next.next
	}
	fmt.Println(tail)

}
