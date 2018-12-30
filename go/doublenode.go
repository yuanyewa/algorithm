package main

import "fmt"

type doubleNode struct {
	value interface{}
	prev  *doubleNode
	next  *doubleNode
}

func (d *doubleNode) link(n *doubleNode) *doubleNode {
	// d, n: d -->n
	d.next = n
	n.prev = d
	return n
}
func (d *doubleNode) insert(n *doubleNode) *doubleNode {
	// prev --> d --> next
	// prev --> d --> n --> next
	n.next = d.next
	n.prev = d
	d.next.prev = n
	d.next = n
	return n
}
func (d *doubleNode) delete() {
	// p --> d --> n
	// p --> n
	d.prev.next = d.next
	d.next.prev = d.prev
}

func main() {
	var head, d0, d1, d2, tail doubleNode
	d0.value = "0"
	d1.value = "1"
	d2.value = "2"
	head.link(&d0).link(&d1).link(&d2).link(&tail)
	n := doubleNode{"new node", nil, nil}
	d1.insert(&n)
	d1.delete()
	fmt.Println(head.next.next.next.prev)
}
