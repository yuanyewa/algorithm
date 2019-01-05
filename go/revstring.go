package main

import (
	"fmt"
	"time"
)

func reverse(str string) string { // only works for ascii
	if len(str) <= 1 {
		return str
	}
	// return strings.Join([]string{string(str[len(str)-1]), reverse(str[:len(str)-1])}, "")
	// return string(str[len(str)-1]) + reverse(str[:len(str)-1])
	return reverse(str[1:]) + string(str[0])
}

func perm(pre string, s []string, str string) []string {
	if len(str) <= 1 {
		return append(s, pre+str)
	}
	for i, v := range str {
		s = perm(pre+string(v), s, str[:i]+str[i+1:])
	}
	return s
}

func perm2(str string) []string {
	if len(str) <= 1 {
		return []string{str}
	}
	var s []string
	for i, v := range str {
		for _, p := range perm2(str[:i] + str[i+1:]) {
			s = append(s, string(v)+p)
		}
	}
	return s
}
func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		reverse("abcdefghijk")
	}
	fmt.Println(time.Since(start))

	start = time.Now()
	for i := 0; i < 100000; i++ {
		perm("", []string{}, "abcde")
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 100000; i++ {
		perm2("abcde")
	}
	fmt.Println(time.Since(start))
}
