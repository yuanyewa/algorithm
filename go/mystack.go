package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type deque []interface{}

func (d *deque) isEmpty() bool {
	return len(*d) == 0
}

func (d *deque) addFront(x interface{}) {
	(*d) = append(deque{x}, (*d)...)
}

func (d *deque) addRear(x interface{}) {
	(*d) = append(*d, x)
}

func (d *deque) popFront() interface{} {
	var v interface{}
	v = (*d)[0]
	(*d) = (*d)[1:]
	return v
}
func (d *deque) popRear() interface{} {
	var v interface{}
	v = (*d)[len(*d)-1]
	(*d) = (*d)[:len(*d)-1]
	return v
}
func (d *deque) String() string {
	var b bytes.Buffer
	for _, r := range *d {
		switch x := r.(type) {
		case int:
			b.WriteString(strconv.Itoa(x))
		case rune:
			b.WriteString(string(x))
		}
	}
	return b.String()
}

type queue []interface{}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) enQueue(x interface{}) {
	(*q) = append(*q, x)

}

func (q *queue) DeQueue() interface{} {
	var v interface{}
	if len(*q) == 0 {
		return v
	}
	v = (*q)[0]
	(*q) = (*q)[1:len(*q)]
	return v
}
func (q *queue) String() string {
	var b bytes.Buffer
	for _, r := range *q {
		switch x := r.(type) {
		case int:
			b.WriteString(strconv.Itoa(x))
		case rune:
			b.WriteString(string(x))
		}
	}
	return b.String()
}

type stack []interface{}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}
func (s *stack) push(x interface{}) {
	*s = append(*s, x)
}
func (s *stack) pop() interface{} {
	var v interface{}
	if len(*s) == 0 {
		fmt.Println("s is empty")
		return v
	}
	v = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
func (s stack) len() int {
	return len(s)
}

func (s stack) String() string {
	var b bytes.Buffer
	for _, r := range s {
		switch x := r.(type) {
		case int:
			b.WriteString(strconv.Itoa(x))
		case rune:
			b.WriteString(string(x))
		}
	}
	return b.String()
}
func main() {
	mys := stack{}
	fmt.Println(mys.isEmpty())
	mys.push(5)
	fmt.Println(mys)
	mys.push('a')
	fmt.Println(mys)
	fmt.Println(mys.pop())
	fmt.Println(mys)
	fmt.Println(mys.len())
	sentence := "this is me!!"
	s2 := stack{}
	for _, v := range sentence {
		s2.push(v)
	}
	b := bytes.Buffer{}
	ll := len(s2)
	for i := 0; i < ll; i++ {
		switch x := s2.pop().(type) {
		case int:
			b.WriteString(strconv.Itoa(x))
		case rune:
			b.WriteString(string(x))
		}
	}
	fmt.Println(b.String())

	q := queue{}
	q.enQueue("a")
	q.enQueue("b")
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q)

	d := deque{}
	d.addFront("A")
	d.addFront("B")
	d.addRear("C")
	fmt.Println(d)
	fmt.Println(d.popFront(), d)
	fmt.Println(d.popRear(), d)
}
