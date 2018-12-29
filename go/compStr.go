package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func compS(str string) string {
	arr := strings.Split(str, "")
	start := arr[0]
	n := 1
	out := []string{}
	for i, s := range arr[1:] {
		if s != start {
			out = append(out, start, strconv.Itoa(n))
			start = s
			n = 1
		} else {
			n++
		}
		if i == len(arr)-2 {
			out = append(out, start, strconv.Itoa(n))
		}
	}
	return strings.Join(out, "")
}

func compS2(str string) string {
	arr := strings.Split(str, "")
	cnt := 1
	out := []string{}
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			cnt++
		} else {
			out = append(out, arr[i-1], strconv.Itoa(cnt))
			cnt = 1
		}
	}
	out = append(out, arr[len(arr)-1], strconv.Itoa(cnt))
	return strings.Join(out, "")
}
func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		compS("AAAABBBBCCDDE")
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		compS2("AAAABBBBCCDDE")
	}
	fmt.Println(time.Since(start))
}
