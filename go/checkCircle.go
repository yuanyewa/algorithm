package main

import "fmt"

type slist struct {
	value interface{}
	next  *slist
}

func (s *slist) link(n *slist) *slist {
	s.next = n
	return n
}
func main() {
	var s0, s1, s2, s3, s4, s5 slist
	s0.link(&s1).link(&s2).link(&s3).link(&s4).link(&s5)
	p0, p1 := &s0, &s0
	loop := false
	for {
		p0 = p0.next
		if p0 == nil {
			break
		}
		p1 = p1.next.next
		if p1 == nil {
			break
		}
		if p0 == p1 {
			loop = true
			break
		}
	}
	fmt.Println("Looped: ", loop)
	loop = false
	for p1 != nil && p1.next != nil {
		p0 = p0.next
		p1 = p1.next.next
		if p0 == p1 {
			loop = true
			break
		}
	}
	fmt.Println("Looped: ", loop)
}
