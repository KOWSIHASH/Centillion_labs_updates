package main

import (
    "fmt"
    "math"
)

func main() {
    var base, exponent float64
    fmt.Print("Enter base and exponent: ")
    fmt.Scan(&base, &exponent)

    result := math.Pow(base, exponent)
    fmt.Println("Result:", result)
}
