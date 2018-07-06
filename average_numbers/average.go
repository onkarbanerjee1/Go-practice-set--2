package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter number of elements")
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Println("Enter a number ")
		fmt.Scanln(&arr[i])
		sum += arr[i]
	}
	fmt.Print("Avg is ", float64(sum)/float64(n))
}
