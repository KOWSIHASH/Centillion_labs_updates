package main
import "fmt"
func main() {
    arr1 := []int{1, 3, 5}
    arr2 := []int{2, 4, 6}
    merged := append(arr1, arr2...)
    fmt.Println("Merged array:", merged)
}
