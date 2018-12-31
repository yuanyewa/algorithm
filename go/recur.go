package main

import (
	"fmt"
	"strings"
)

func cums1(n uint64) uint64 {
	if n < 1 {
		return 0
	}
	return n + cums1(n-1)
}
func cums2(n uint64) uint64 {
	if n < 10 {
		return n
	}
	return n%10 + cums2(n/10)
}

var words = map[string]struct{}{
	"the": struct{}{},
	"run": struct{}{},
	"man": struct{}{},
}

var output []string

func wordSplit(str string) bool {
	if len(str) == 0 {
		return true
	}
	for word := range words {
		if strings.HasPrefix(str, word) {
			str = strings.TrimPrefix(str, word)
			output = append(output, word)
			return wordSplit(str)
		}
	}
	return false
}

func main() {
	fmt.Println(cums1(100))
	fmt.Println(cums2(987654321))
	fmt.Println(wordSplit("themanruntheman"), output)
}
