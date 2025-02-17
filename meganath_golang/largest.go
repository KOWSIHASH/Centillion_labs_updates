package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Print("Enter three numbers: ")
    fmt.Scan(&a, &b, &c)

    largest := a
    if b > largest {
        largest = b
    }
    if c > largest {
        largest = c
    }

    fmt.Println("Largest number:", largest)
}
