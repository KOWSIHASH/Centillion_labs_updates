package main

import "fmt"

// Function to reverse the digits of a number
func reverseNumber(num int) int {
    reversed := 0
    for num != 0 {
        digit := num % 10          // Get the last digit
        reversed = reversed*10 + digit  // Append the digit to the reversed number
        num /= 10                  // Remove the last digit
    }
    return reversed
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    reversedNum := reverseNumber(num)
    fmt.Println("Reversed number:", reversedNum)
}
