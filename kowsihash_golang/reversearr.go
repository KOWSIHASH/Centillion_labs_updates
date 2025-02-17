package main

import "fmt"

// Function to reverse an array
func reverseArray(arr []int) []int {
    n := len(arr)
    for i := 0; i < n/2; i++ {
        arr[i], arr[n-1-i] = arr[n-1-i], arr[i]  // Swap elements
    }
    return arr
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println("Original array:", arr)

    reversedArr := reverseArray(arr)
    fmt.Println("Reversed array:", reversedArr)
}
