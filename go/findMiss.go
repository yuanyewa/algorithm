package main

import (
	"fmt"
	"sort"
	"time"
)

func miss(x, y []int) int {
	// find the element missing from x
	for i, v := range x {
		hasV := false
		for m, n := range y {
			if v == n {
				x = append(x[:i], x[i+1:]...)
				y = append(y[:m], y[m+1:]...)
				hasV = true
				break
			}
		}
		if !hasV {
			return v
		}
	}
	return -1
}
func miss2(x, y []int) int {
	sort.IntSlice(x).Sort()
	sort.IntSlice(y).Sort()
	for i := 0; i < len(y)-1; i++ {
		if x[i] != y[i] {
			return x[i]
		}
	}
	return x[len(x)-1]
}

func miss3(x, y []int) int {
	d := make(map[int]int)
	for _, v := range y {
		d[v]++
	}
	for _, v := range x {
		if d[v] == 0 {
			return v
		}
		d[v]--
	}
	return -1
}

func miss4(x, y []int) int {
	r := 0
	for _, v := range append(x, y...) {
		r ^= v
	}
	return r
}
func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		miss([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 7, 2, 1, 4, 6})
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		miss2([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 7, 2, 1, 4, 6})
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		miss3([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 7, 2, 1, 4, 6})
	}
	fmt.Println(time.Since(start))
	fmt.Println(miss4([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 5, 2, 1, 4, 6}))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		miss4([]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 7, 2, 1, 4, 6})
	}
	fmt.Println(time.Since(start))
}
