package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	ss := stack.New()
	ss.Push(3)
	fmt.Println(ss)
}
