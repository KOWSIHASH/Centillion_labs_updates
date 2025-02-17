package main

import "fmt"

func main() {
	var num, count int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	for num != 0 {
		num /= 10
		count++
	}
	fmt.Println("Number of digits:", count)
}