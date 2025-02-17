package main
import (
	"fmt"
	"math"
)
func secondLargest(arr []int) int {
	largest, second := math.MinInt, math.MinInt

	for _, num := range arr {
		if num > largest {
			second = largest
			largest = num
		} else if num > second && num != largest {
			second = num
		}
	}
	return second
}
func main() {
	arr := []int{12, 45, 23, 89, 56, 90, 78}
	fmt.Println("Second Largest:", secondLargest(arr))
}
