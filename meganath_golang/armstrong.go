package main
import (
    "fmt"
    "math"
)
func main() {
    var num, temp, remainder, result int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    temp = num
    for temp != 0 {
        remainder = temp % 10
        result += int(math.Pow(float64(remainder), 3))
        temp /= 10
    }
    if result == num {
        fmt.Println(num, "is an Armstrong Number")
    } else {
        fmt.Println(num, "is Not an Armstrong Number")
    }
}
