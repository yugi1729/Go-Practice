package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter an integer: ")
	fmt.Scanf("%d", &input)

	temp := input
	var reverse int = 0
	for temp != 0 {
		reverse = reverse * 10
		reverse = reverse + temp%10
		temp = temp / 10

	}
	fmt.Println(reverse)

}
