package main

import "fmt"

func isEven(num int) bool {
    return num%2 == 0
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if isEven(num) {
        fmt.Println(num, "is an even number.")
    } else {
        fmt.Println(num, "is not an even number.")
    }
}
