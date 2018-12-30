package main

import "fmt"

type stack2 []interface{}
type queue2 struct {
	inStack  stack2
	outStack stack2
}

func (s *stack2) push(x interface{}) {
	(*s) = append(*s, x)
}
func (s *stack2) pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}
func (s *stack2) size() int {
	return len(*s)
}
func (q *queue2) enqueue(x interface{}) {
	q.inStack.push(x)
}
func (q *queue2) dequeue() interface{} {
	if q.outStack.size() == 0 {
		for x := q.inStack.pop(); x != nil; {
			q.outStack.push(x)
			x = q.inStack.pop()
		}
	}
	return q.outStack.pop()
}
func main() {
	q := queue2{}
	for i := 0; i < 5; i++ {
		q.enqueue(i)
	}
	fmt.Println(q.dequeue(), q.outStack)
	fmt.Println(q.dequeue(), q.outStack)
	fmt.Println(q.dequeue(), q.outStack)
	for i := 5; i < 10; i++ {
		q.enqueue(i)
	}
	fmt.Println(q.dequeue(), q.outStack)
	fmt.Println(q.dequeue(), q.outStack)
	fmt.Println(q.dequeue(), q.outStack)
	fmt.Println(q.dequeue(), q.outStack)
	fmt.Println(q.dequeue(), q.outStack)
	fmt.Println(q.dequeue(), q.outStack)
}
