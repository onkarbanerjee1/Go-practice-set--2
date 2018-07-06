package main

import (
	"fmt"
)

// Output: Enter number of rows
// 5
//     1
//    232
//   34543
//  4567654
// 567898765

func main() {
	fmt.Println("Enter number of rows")
	var rows int
	fmt.Scanln(&rows)
	for i := 0; i < rows; i++ {
		rowStart := i
		for j := 0; j < (2*rows)-1; j++ {
			if j >= rows-1-i && j <= rows-1+i {
				if j <= rows-1 {
					rowStart++
				} else {
					rowStart--
				}
				fmt.Print(rowStart)
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
