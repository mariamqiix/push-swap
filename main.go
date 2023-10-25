package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers []int
var B []int
var ans []string

func main() {
	if len(os.Args) != 2 {
		Exit()
	}

	array := strings.Fields(os.Args[1])

	for i := 0; i < len(array); i++ {
		num, err := strconv.Atoi(array[i])
		if err != nil {
			Exit()
		}
		numbers = append(numbers, num)
	}

	fmt.Println(numbers)

	for !SortInteger {
		sb(numbers, "a") //sa swap first 2 elements of stack a
		sb(numbers, "b") //sb swap first 2 elements of stack b

		if numbers[len(numbers)-1] < numbers[0] {
			numbers[0], numbers[len(numbers)-1] = numbers[len(numbers)-1], numbers[0]
			ans = append(ans, "pa")
		}

	}
}

func Exit() {
	fmt.Println("Error")
	os.Exit(0)
}

func SortInteger(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				return false
			}
		}
	}
	return true
}

func sb(numbers []int, s string) bool {
	if numbers[1] < numbers[0] {
		numbers[1], numbers[0] = numbers[0], numbers[1]
		ans = append(ans, "s"+s)
		return true
	}
	return false
}
