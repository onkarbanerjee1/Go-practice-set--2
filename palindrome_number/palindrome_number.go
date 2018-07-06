package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter a number")
	var n int
	fmt.Scanln(&n)
	if isPalindrome(n) {
		fmt.Println("It is a palindrome")
	} else {
		fmt.Println("It is NOT a palindrome")
	}

}

func isPalindrome(n int) bool {
	l := int(math.Log10(float64(n)) + 1)
	if l <= 1 {
		return true
	}
	lsb := n % 10
	msb := n / int(math.Pow10(l-1))
	if lsb == msb {
		return isPalindrome((n - (msb * int(math.Pow10(l-1))) - lsb) / 10)
	}
	return false
}
