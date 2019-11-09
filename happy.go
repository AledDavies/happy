package main

import (
	"fmt"
	"os"
	"strconv"
)

func Digits(number int) []int {
	digits := make([]int, 0)

	for ; number > 0; number /= 10 {
		digits = append(digits, number%10)
	}

	return digits
}

func SumSquare(digits []int) int {
	sum := 0
	for _, digit := range digits {
		sum = sum + (digit * digit)
	}
	return sum
}

func Happy(number int) bool {
	for seen := map[int]bool{}; !seen[number]; {
		fmt.Printf("%v -> ", number)
		seen[number] = true
		number = SumSquare(Digits(number))
	}
	fmt.Printf("%v\n", number)
	return number == 1
}

func main() {
	if len(os.Args) > 1 {
		number, err := strconv.Atoi(os.Args[1])

		if err == nil && number > 0 {
			fmt.Printf("%v\n", Happy(number))
		} else {
			fmt.Printf("Bad number format %v", os.Args[1])
		}
	} else {
		fmt.Printf("Usage: happy <num>")
	}
}
