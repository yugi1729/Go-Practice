package main

import "fmt"

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}

}

func rainWater(heights []int) int {
	length := len(heights)
	if length <= 2 {
		return 0
	}
	left, right := make([]int, length), make([]int, length)
	left[0], right[length-1] = heights[0], heights[length-1]

	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], heights[i])
	}

	for i := length - 2; i > -1; i-- {
		right[i] = max(right[i+1], heights[i])
	}

	ans := 0
	for i := 0; i < length; i++ {
		ans += min(left[i], right[i]) - heights[i]
	}

	return ans
}

func main() {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(rainWater(heights))

}
