package main

import (
	"fmt"
)

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {

	arr := []int{4, 3, 2, 1}
	fmt.Println(BubbleSort(arr))
}
