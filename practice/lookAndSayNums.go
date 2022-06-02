package main

import (
	"fmt"
	"strconv"
)

func lookAndSay(num string) string {
	tmp := ""
	count := 1
	s := num + "$"
	for i := 1; i <= len(s)-1; i++ {

		if s[i-1] != s[i] {
			tmp += strconv.Itoa(count)
			tmp += string(s[i-1])
			count = 1
		} else {
			count++
		}
	}
	return tmp
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	var iters int
	fmt.Scanf("%d", &iters)
	for iters > 0 {
		fmt.Println(lookAndSay(s))
		s = lookAndSay(s)
		iters--
	}
}
