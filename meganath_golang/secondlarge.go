package main

import "fmt"

func main() {
    arr := []int{10, 20, 4, 45, 99}
    first, second := 0, 0

    for _, num := range arr {
        if num > first {
            second = first
            first = num
        } else if num > second && num != first {
            second = num
        }
    }

    fmt.Println("Second largest number:", second)
}
