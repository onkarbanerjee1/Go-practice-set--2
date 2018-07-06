package main

import "fmt"
import "os"

func main() {
	if len(os.Args) != 4 {
		fmt.Println("3 numbers needed to be input")
		os.Exit(101)
	}
	largest := os.Args[1]
	if os.Args[2] > largest {
		largest = os.Args[2]
	}
	if os.Args[3] > largest {
		largest = os.Args[3]
	}
	fmt.Println("Largest is ", largest)
}
