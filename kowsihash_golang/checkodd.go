package main

import "fmt"

func isOdd(num int) bool {
    return num%2 != 0
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if isOdd(num) {
        fmt.Println(num, "is an odd number.")
    } else {
        fmt.Println(num, "is not an odd number.")
    }
}
