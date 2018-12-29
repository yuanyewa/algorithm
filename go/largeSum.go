package main

import (
	"fmt"
	"time"
)

func large(x []int) int {
	if len(x) == 0 {
		return 0
	}
	m, c := x[0], x[0]
	for _, n := range x[1:] {
		// c = max(c+n, n)
		if c > 0 {
			c += n
		} else {
			c = n
		}
		// m = max(m, c)
		if m < c {
			m = c
		}
	}
	return m
}
func main() {
	fmt.Println(large([]int{1, 2, -1, 3, 4, 10, 10, -10, -1}))
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		large([]int{1, 2, -1, 3, 4, 10, 10, -10, -1})
	}
	fmt.Println(time.Since(start))
}
