package main

import (
	"fmt"
	"strings"
	"time"
)

func revword(str string) string {
	s := strings.Fields(str)
	// for i := 0; i < len(s)/2; i++ {
	// 	s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	// }
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, " ")
}

func main() {
	fmt.Println(revword("  This  is the    best    "))
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		revword(" This is the best ")
	}
	fmt.Println(time.Since(start))
}
