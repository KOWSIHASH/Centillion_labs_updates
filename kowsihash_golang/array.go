package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array elements:")
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}
}
