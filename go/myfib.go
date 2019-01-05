package main

import (
	"fmt"
	"time"
)

func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

var m = map[int]int{}

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	if v, ok := m[n]; ok {
		return v
	}
	m[n] = fib2(n-1) + fib2(n-2)
	return m[n]
}
func fib3(n int) int {
	// f(n) = f(n-1) + f(n-2)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	fmt.Println(fib2(10))
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			fib1(10)
		}
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			fib2(10)
		}
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		for i := 0; i < 1000000; i++ {
			fib3(10)
		}
		fmt.Println(time.Since(start))
	}
}
