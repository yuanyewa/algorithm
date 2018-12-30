package main

import "fmt"

type list3 struct {
	value interface{}
	next  *list3
}

func (l *list3) link(n *list3) *list3 {
	l.next = n
	return n
}
func reverse(head *list3) *list3 {
	curr := head
	var prev, next *list3
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev, curr = curr, next
	}
	return prev
}

func main() {
	l0, l1, l2, l3 := list3{0, nil}, list3{1, nil}, list3{2, nil}, list3{3, nil}
	l0.link(&l1).link(&l2).link(&l3)
	head := reverse(&l0)
	fmt.Println(head, head.next, head.next.next, head.next.next.next, head.next.next.next.next)
}
