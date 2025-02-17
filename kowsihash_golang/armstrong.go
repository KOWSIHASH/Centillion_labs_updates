package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	sum, temp := 0, n
	digits := int(math.Log10(float64(n))) + 1
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if isArmstrong(num) {
		fmt.Println(num, "is an Armstrong number")
	} else {
		fmt.Println(num, "is not an Armstrong number")
	}
}