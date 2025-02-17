package main

import (
	"fmt"
	"math"
)

func findTwoLargest(arr []int) (int, int) {
	if len(arr) < 2 {
		fmt.Println("Array should have at least two elements")
		return math.MinInt, math.MinInt
	}

	first, second := math.MinInt, math.MinInt

	for _, num := range arr {
		if num > first {
			second = first
			first = num
		} else if num > second && num != first {
			second = num
		}
	}

	return first, second
}

func main() {
	arr := []int{10, 20, 4, 45, 99, 99, 56, 89}
	largest, secondLargest := findTwoLargest(arr)
	fmt.Println("Largest:", largest)
	fmt.Println("Second Largest:", secondLargest)
}