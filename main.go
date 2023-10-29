package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a []int
var b []int
var ans []string

// ss execute sa and sb

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
		a = append(a, num)
	}
	length := len(a)

	fmt.Println(a, b)

	for i := 0; !(SortInteger(a) && len(a) == length) && i < 20; i++ {

		if len(a) > 1 && a[1] < a[0] {
			a = s(a) //sa swap first 2 elements of stack a
			ans = append(ans, "sa")

		} else if len(b) > 1 && b[1] < b[0] {
			b = s(b) //sb swap first 2 elements of stack b
			ans = append(ans, "sb")

		} else if (len(b) > 1 && len(b) > 1 && a[0] > b[0]) || i < 2 {
			p("b")

		} else if (len(a) > 1 && len(b) > 1 && b[0] > a[0]) || i < 2 {
			p("a")

		} else if len(a) > 1 && a[0] > a[len(a)-1] {
			a = rotate(a)
			ans = append(ans, "ra")

		} else if len(b) > 1 && b[0] > b[len(b)-1] {
			b = rotate(b)
			ans = append(ans, "rb")

		} else if len(a) > 1 && a[0] > a[len(a)-1] {
			a = reverse(a)
			ans = append(ans, "rra")

		} else if len(b) > 1 && b[0] > b[len(b)-1] {
			b = reverse(b)
			ans = append(ans, "rrb")

		} 

		fmt.Println(a, b)
	}

	fmt.Println(ans)
}


// else if len(b) > 1 && len(a) > 1 {
// 	a = rotate(a)
// 	b = rotate(b)
// 	ans = append(ans, "rr")

// } 

// else if len(b) > 1 && len(a) > 1 {
// 	a = reverse(a)
// 	b = reverse(b)
// 	ans = append(ans, "rrr")

// }

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

func s(numbers []int) []int {
	numbers[1], numbers[0] = numbers[0], numbers[1]
	return numbers
}

func p(s string) {
	if s == "a" && len(b) != 0 {
		b, a = pushing(b, a)
	} else if s == "b" && len(a) != 0 {
		a, b = pushing(a, b)
	}

	ans = append(ans, "p"+s)
}

func pushing(arr1, arr2 []int) ([]int, []int) {
	new := make([]int, len(arr2)+1)
	new[0] = arr1[0]
	copy(new[1:], arr2)
	arr2 = new

	arr1 = append(arr1[:0], arr1[1:]...)
	return arr1, arr2
}

func rotate(arr1 []int) []int {
	// Shift up the elements
	new := arr1[0]
	copy(arr1, arr1[1:])
	arr1[len(arr1)-1] = new

	// Print the updated stack
	fmt.Println(arr1)
	return arr1
}

func reverse(arr1 []int) []int {
	// Shift down the elements
	lastElement := arr1[len(arr1)-1]
	copy(arr1[1:], arr1[:len(arr1)-1])
	arr1[0] = lastElement

	// Print the updated arr1
	fmt.Println(arr1)
	return arr1
}
