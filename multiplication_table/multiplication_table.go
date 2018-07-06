package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number for whose multiplication table to be generated")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Enter length of table")
	var m int
	fmt.Scanln(&m)
	for i := 1; i <= m; i++ {
		fmt.Printf(" %d X %d = %d \n", n, i, n*i)
	}
}
