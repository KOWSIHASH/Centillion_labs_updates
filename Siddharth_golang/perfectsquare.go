package main
import (
	"fmt"
	"math"
)

func isPerfectSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	fmt.Println("Perfect Square:", isPerfectSquare(num))
}
