package main 

import "fmt"

func isPrime(n int) bool{
	p := true
	for i := 2; i < n; i++ {
		if n % i == 0 {
			p = false
		}
	}
	return p
}

func printPrimes(num int) {
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}	

func main() {
	printPrimes(20)
}