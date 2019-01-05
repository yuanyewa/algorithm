package main

import (
	"fmt"
	"time"
)

func rec1(target int, coins map[int]bool) int {
	minCoins := target
	if _, ok := coins[target]; ok {
		return 1
	}
	for k := range coins {
		if k > target {
			continue
		}
		numCoins := 1 + rec1(target-k, coins)
		if numCoins < minCoins {
			minCoins = numCoins
		}
	}
	return minCoins
}

var m = make(map[int]int)

func rec2(target int, coins []int) int {
	minCoins := target
	for _, v := range coins {
		if v == target {
			return 1
		}
		if v > target {
			continue
		}
		numCoins := 1
		if c, ok := m[target-v]; ok {
			numCoins += c
		} else {
			c := rec2(target-v, coins)
			m[target-v] = c
			numCoins += c
		}
		if numCoins < minCoins {
			minCoins = numCoins
		}
	}
	return minCoins
}
func main() {
	{
		start := time.Now()
		// fmt.Println(rec1(23, map[int]bool{1: true, 5: true, 10: true, 25: true}))
		fmt.Println(time.Since(start))
	}
	{
		start := time.Now()
		fmt.Println(rec2(1234, []int{1, 5, 10, 25}))
		fmt.Println(time.Since(start))
	}
}
