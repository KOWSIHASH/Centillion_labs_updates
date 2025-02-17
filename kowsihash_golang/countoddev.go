package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd, even := 0, 0
	for _, num := range arr {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println("Odd count:", odd, "Even count:", even)
}