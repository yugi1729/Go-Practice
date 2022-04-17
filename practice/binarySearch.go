package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}


func main(){
	arr := []int{1, 2, 3, 5, 6}
	fmt.Println(binarySearch(arr, 5))
}