package main

import "fmt"

// Function to swap two numbers
func swap(a, b int) (int, int) {
    return b, a
}

func main() {
    var num1, num2 int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&num1, &num2)

    fmt.Println("Before swapping: num1 =", num1, "num2 =", num2)
    
    // Swap the numbers
    num1, num2 = swap(num1, num2)

    fmt.Println("After swapping: num1 =", num1, "num2 =", num2)
}
