package main

import (
	"fmt"
	"time"
)

func checkU(str string) bool {
	tmp := make(map[rune]struct{})
	for _, r := range str {
		_, ok := tmp[r]
		if ok {
			return false
		}
		tmp[r] = struct{}{}
	}
	return true
}
func main() {
	fmt.Println(checkU("abcde"))
	fmt.Println(checkU("abcade"))

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		checkU("abcade")
	}
	fmt.Println(time.Since(start))
}
