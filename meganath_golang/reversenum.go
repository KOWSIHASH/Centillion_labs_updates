package main

import "fmt"

func main() {
    var num, rev int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    for num > 0 {
        rev = rev*10 + num%10
        num /= 10
    }
    fmt.Println("Reversed Number:", rev)
}
