package main

import "fmt"

func main() {
    arr := []int{10, 20, 30, 40, 50}
    indexToRemove := 2 

    arr = append(arr[:indexToRemove], arr[indexToRemove+1:]...)

    fmt.Println("Array after removal:", arr)
}
