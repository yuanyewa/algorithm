package main

import (
	"fmt"
	"time"
)

func pairSum(x []int, k int) {
	pairs := make(map[[2]int]struct{})
	for i, v := range x {
		for _, n := range x[i+1:] {
			if v+n == k {
				if v > n {
					v, n = n, v
				}
				pairs[[2]int{v, n}] = struct{}{}
			}
		}
	}
}

func pairSum2(x []int, k int) {
	pairs := make(map[[2]int]struct{})
	seen := make(map[int]struct{})
	for _, v := range x {
		target := k - v
		if _, ok := seen[target]; ok {
			// delete(seen, target)
			if v > target {
				v, target = target, v
			}
			pairs[[2]int{v, target}] = struct{}{}
		} else {
			seen[v] = struct{}{}
		}
	}
}

func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		pairSum([]int{1, 3, 2, 2, -1, 5, -2, 6}, 4)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		pairSum2([]int{1, 3, 2, 2, -1, 5, -2, 6}, 4)
	}
	fmt.Println(time.Since(start))
}
