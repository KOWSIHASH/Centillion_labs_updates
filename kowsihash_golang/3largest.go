package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{10, 5, 20, 8, 25, 18, 30}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println("The three largest numbers are:", arr[0], arr[1], arr[2])
}