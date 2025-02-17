package main

import "fmt"

func main() {
    var n, sum int
    fmt.Print("Enter number of elements: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Enter the elements:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
        sum += arr[i]
    }

    fmt.Println("Sum of array elements:", sum)
}
