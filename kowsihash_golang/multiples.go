package main

import "fmt"


func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}


func lcm(a, b int) int {
    return (a * b) / gcd(a, b)
}

func main() {
    var num1, num2 int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&num1, &num2)

    fmt.Println("LCM of", num1, "and", num2, "is:", lcm(num1, num2))
}
