package main

import "fmt"

func main() {
    arr := []int{10, 21, 32, 43, 54, 65}
    evenCount, oddCount := 0, 0

    for _, num := range arr {
        if num%2 == 0 {
            evenCount++
        } else {
            oddCount++
        }
    }

    fmt.Println("Even numbers count:", evenCount)
    fmt.Println("Odd numbers count:", oddCount)
}
