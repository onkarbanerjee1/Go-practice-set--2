package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	doorOptions := []int{1, 2, 3}
	rand.Seed(time.Now().Unix())
	x := rand.Intn(len(doorOptions))
	fmt.Println(doorOptions[x])

	inr := bufio.NewReader(os.Stdin)
	line, err := inr.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read Y")
		os.Exit(1)
	}

	y, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		fmt.Println("Could not read Y")
		os.Exit(1)
	}

	if y < 1 || y > 3 || y == doorOptions[x] {
		fmt.Println("Constraint 1≤Y≤3 and X≠Y", y)
		os.Exit(1)
	}

	newOptions := []int{}
	for index, door := range doorOptions {
		if index == y {
			continue
		}
		newOptions = append(newOptions, door)
	}
	z := rand.Intn(len(newOptions))
	fmt.Println(newOptions[z])
}
