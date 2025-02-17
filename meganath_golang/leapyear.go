package main

import "fmt"

func main() {
    var year int
    fmt.Print("Enter a year: ")
    fmt.Scan(&year)

    if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
        fmt.Println(year, "is a Leap Year")
    } else {
        fmt.Println(year, "is Not a Leap Year")
    }
}
