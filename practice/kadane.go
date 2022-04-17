package main

import (
	"fmt"
)

func maxInt(x, y int) int{
	if(x > y){
		return x
	}else{
		return y
	}
}

func main(){
	var length int
	fmt.Scanf("%d\n",&length)
	arr := make([]int,length)
	for i, _ := range arr{
		fmt.Scanf("%d",&arr[i])
	}
	sum := 0
	maxSoFar := 0
	for _, v :=  range arr{
		sum += v
		maxSoFar = maxInt(maxSoFar, sum)
		sum = maxInt(sum, 0)
	}
	fmt.Println(maxSoFar)
}

 