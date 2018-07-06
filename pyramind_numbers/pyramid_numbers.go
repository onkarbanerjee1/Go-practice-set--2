package main

import (
	"fmt"
)

// Output:
// enter the number of rows(+ve)
// 3
//   1
//  234
// 56789

func main() {
	fmt.Println("enter the number of rows(+ve)")
	var n int
	fmt.Scanln(&n)
	c := 1
	for i := 0; i < n; i++ {
		for j := 0; j < (2*n)-1; j++ {
			if j >= (n-1-i) && j <= (n-1+i) {
				fmt.Print(c)
				c++
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
