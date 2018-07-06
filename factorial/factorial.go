package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Println("Enter the number for factorial")
	fmt.Scanln(&n)
	fmt.Println("Factorial of given number is , ", factorial(n))
}
