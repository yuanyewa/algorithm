package main

import (
	"fmt"
)

type myStack []rune

func (m *myStack) isEmpty() bool {
	return len(*m) == 0
}

func (m *myStack) push(r rune) {
	(*m) = append(*m, r)
}

func (m *myStack) pop() rune {
	if len(*m) == 0 {
		return ' '
	}
	v := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return v
}

func (m *myStack) size() int {
	return len(*m)
}

func main() {
	str := "{{{{{[[[[[((((({{{{{}}}}})))))]]]]]}}}}}"
	left := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	right := map[rune]struct{}{
		'}': struct{}{},
		']': struct{}{},
		')': struct{}{},
	}
	m := myStack{}
	for _, r := range str {
		if _, ok := left[r]; ok {
			m.push(r)
		} else {
			if _, ok := right[r]; ok {
				if m.isEmpty() {
					m.push(r)
					break
				}
				if left[m.pop()] != r {
					break
				}
			}
		}
	}
	fmt.Println(m.isEmpty())
}
