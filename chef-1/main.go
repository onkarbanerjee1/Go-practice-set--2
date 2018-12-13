package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inr := bufio.NewReader(os.Stdin)

	firstLine, err := inr.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read the first line", err)
		os.Exit(1)
	}

	inputs := strings.Split(firstLine, " ")
	if len(inputs) != 2 {
		fmt.Println("First line should contain two space sepaarted values")
		os.Exit(1)
	}

	n, err := strconv.Atoi(strings.TrimSpace(inputs[0]))
	if err != nil {
		fmt.Println("Could not retrieve N ", err)
		os.Exit(1)
	}

	r, err := strconv.Atoi(strings.TrimSpace(inputs[1]))
	if err != nil {
		fmt.Println("Could not retrieve r ", err)
		os.Exit(1)
	}

	if n < 1 || n > 1000 {
		fmt.Println("Constraint: 1≤N≤1,000")
		os.Exit(1)
	}

	if r < 1300 {
		fmt.Println("Constraint: 1,300≤r")
		os.Exit(1)
	}

	for ; n > 0; n-- {
		ins, err := inr.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read input", err)
			os.Exit(1)
		}

		R, err := strconv.Atoi(strings.TrimSpace(ins))
		if err != nil {
			fmt.Println("Could not read input", err)
			os.Exit(1)
		}

		if R > 1501 {
			fmt.Println("R≤1,501")
			os.Exit(1)
		}

		if R < r {
			fmt.Println("Bad boi")
		} else {
			fmt.Println("Good boi")
		}
	}

}
