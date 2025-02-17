package main

import "fmt"

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func main() {
    var num1, num2 int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&num1, &num2)

    fmt.Println("GCD:", gcd(num1, num2))
}
