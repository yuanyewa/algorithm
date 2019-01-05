package main

import (
	"fmt"
	"time"
)

type list4 struct {
	value interface{}
	next  *list4
}

func (l *list4) link(n *list4) *list4 {
	l.next = n
	return n
}

type stack4 []interface{}

func (s *stack4) push(v interface{}) {
	(*s) = append(*s, v)
}
func (s *stack4) pop() interface{} {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}
func (s *stack4) isEmpty() bool {
	return len(*s) == 0
}
func (s *stack4) size() int {
	return len(*s)
}

func n2Last(head *list4, n int) interface{} {
	s := stack4{}
	curr := head
	for curr != nil {
		s.push(curr.value)
		curr = curr.next
	}
	if n > s.size() {
		return nil
	}
	// if n == s.size() {
	// 	return head.value
	// }
	// var v interface{}
	// for ; n > 0; n-- {
	// 	v = s.pop()
	// }
	// return v
	return s[len(s)-n]
}

func n2Last2(head *list4, n int) interface{} {
	left, right := head, head
	for ; n > 0; n-- {
		right = right.next
	}
	for right != nil {
		left = left.next
		right = right.next
	}
	return left.value
}

func main() {
	l0, l1, l2, l3 := list4{0, nil}, list4{1, nil}, list4{2, nil}, list4{3, nil}
	l0.link(&l1).link(&l2).link(&l3)
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		n2Last(&l0, 3)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		n2Last2(&l0, 3)
	}
	fmt.Println(time.Since(start))
}
