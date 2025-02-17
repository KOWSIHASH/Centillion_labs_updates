package main
import (
    "fmt"
    "sort"
)

func main() {
    var n1, n2 int
    fmt.Print("Enter size of first array: ")
    fmt.Scan(&n1)
    arr1 := make([]int, n1)

    fmt.Println("Enter",n1,"elements of first array:")
    for i := 0; i < n1; i++ {
        fmt.Scan(&arr1[i])
    }

    fmt.Print("Enter size of second array: ")
    fmt.Scan(&n2)
    arr2 := make([]int, n2)

    fmt.Println("Enter",n2,"elements of second array:")
    for i := 0; i < n2; i++ {
        fmt.Scan(&arr2[i])
    }

    merged := append(arr1, arr2...)

    sort.Ints(merged)

    fmt.Println("Merged and sorted array:",merged)
}
