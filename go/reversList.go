package main

import "fmt"

type list2 struct {
	value interface{}
	next  *list2
}

type stack2 []interface{}

func (s *stack2) push(v interface{}) {
	(*s) = append(*s, v)
}

func (s *stack2) pop() interface{} {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func (s *stack2) isEmpty() bool {
	return len(*s) == 0
}

func (l *list2) link(n *list2) *list2 {
	l.next = n
	return n
}
func (l *list2) reverse() *list2 {
	s := stack2{}
	c := l // head also in stack
	for c != nil && c.next != nil {
		s.push(c)
		c = c.next
	} // c.next == nil
	l = c
	for !s.isEmpty() {
		if x, ok := s.pop().(*list2); ok {
			c.next = x
		}
		c = c.next
	}
	return l
}

func main() {
	l0, l1, l2, l3 := list2{0, nil}, list2{1, nil}, list2{2, nil}, list2{3, nil}
	l0.link(&l1).link(&l2).link(&l3)
	l0 = (*l0.reverse())
	fmt.Println(l0.next.next)
}
