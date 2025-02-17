package main

import (
	"fmt"
	"sort"
)

// Function to sort an array
func sortArray(arr []int) []int {
	sort.Ints(arr) // Sorts the array in ascending order
	return arr
}

func main() {
	arr := []int{5, 3, 8, 1, 2, 7, 4, 6}
	fmt.Println("Original array:", arr)

	sortedArr := sortArray(arr)
	fmt.Println("Sorted array:", sortedArr)
}
