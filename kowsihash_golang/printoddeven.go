package main

import "fmt"

func printOddEvenNumbers(start, end int) {
    fmt.Println("Even numbers:")
    for i := start; i <= end; i++ {
        if i%2 == 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println()

    fmt.Println("Odd numbers:")
    for i := start; i <= end; i++ {
        if i%2 != 0 {
            fmt.Print(i, " ")
        }
    }
    fmt.Println()
}

func main() {
    var start, end int
    fmt.Print("Enter start and end values: ")
    fmt.Scan(&start, &end)

    printOddEvenNumbers(start, end)
}
