package main 

import (
	"fmt"
	
)



func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}else if n < 0 {
		return 0
	}
	return n * fact(n-1)
}

func memoizedFact(n int) int {
	memo := make(map[int]int)
	if n == 0 || n == 1 {
		return 1
	}
	if val := memo[n]; val != 0 {
		return val
	}
	memo[n] = n * memoizedFact(n-1)
	return memo[n]
}

func main() {
	fmt.Println(fact(-5))
	fmt.Println(memoizedFact(10))
}