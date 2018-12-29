package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
)

func anagram1(x string, y string) bool {
	x = strings.ToLower(strings.Replace(x, " ", "", -1))
	y = strings.ToLower(strings.Replace(y, " ", "", -1))
	{
		x := strings.Split(x, "")
		sort.Strings(x)
		y := strings.Split(y, "")
		sort.Strings(y)
		if len(x) != len(y) {
			return false
		}
		for i, v := range x {
			if v != y[i] {
				return false
			}
		}
		return true
	}
}

func anagram2(x, y string) bool {
	x = strings.ToLower(strings.Replace(x, " ", "", -1))
	y = strings.ToLower(strings.Replace(y, " ", "", -1))
	{
		x := strings.Split(x, "")
		y := strings.Split(y, "")
		for _, v := range x {
			if v == " " {
				continue
			}
			found := false
			for m, n := range y {
				if v == n {
					y = append(y[:m], y[m+1:len(y)]...)
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		if len(strings.Replace(strings.Join(y, ""), " ", "", -1)) > 0 {
			return false
		}
		return true
	}
}
func anagram3(x, y string) bool {
	x = strings.ToLower(strings.Replace(x, " ", "", -1))
	y = strings.ToLower(strings.Replace(y, " ", "", -1))
	{
		x := []rune(x)
		y := []rune(y)
		// x := strings.Split(x, "")
		// y := strings.Split(y, "")
		if len(x) != len(y) {
			return false
		}
		sort.SliceStable(x, func(i, j int) bool { return x[i] < x[j] })
		sort.SliceStable(y, func(i, j int) bool { return y[i] < y[j] })
		for i := range x {
			if x[i] != y[i] {
				return false
			}
		}
	}
	return false
}

func anagram4(x, y string) bool {
	d := make(map[rune]int)
	{
		x := []rune(strings.ToLower(strings.Replace(x, " ", "", -1)))
		y := []rune(strings.ToLower(strings.Replace(y, " ", "", -1)))
		for _, v := range x {
			d[v]++
		}
		for _, v := range y {
			if _, ok := d[v]; !ok {
				return false
			}
			d[v]--
		}
	}
	for _, v := range d {
		if v > 0 {
			return false
		}
	}
	return true
}

func anagram5(x, y string) bool {
	d := make(map[rune]int)
	for _, v := range x {
		if v != ' ' {
			d[unicode.ToLower(v)]++
		}
	}
	for _, v := range y {
		if v != ' ' {
			v := unicode.ToLower(v)
			if _, ok := d[v]; !ok {
				return false
			}
			d[v]--
		}
	}
	for _, v := range d {
		if v > 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	s1 := "abc def gh"
	s2 := "hg fed cba"
	for i := 0; i < 1000000; i++ {
		anagram1(s1, s2)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		anagram2(s1, s2)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		anagram3(s1, s2)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		anagram4(s1, s2)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		anagram5(s1, s2)
	}
	fmt.Println(time.Since(start))
	fmt.Println(string(unicode.ToLower('A')), string(unicode.ToLower('2')), string(unicode.ToLower('世')))
}
