package main

import "fmt"

// Function to find GCD using Euclidean Algorithm
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    var num1, num2 int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&num1, &num2)

    fmt.Println("GCD of", num1, "and", num2, "is:", gcd(num1, num2))
}
