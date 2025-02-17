package main
import "fmt"
func main() {
    arr := []int{5, 3, 8, 1, 2}
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j] 
            }
        }
    }
    fmt.Println("Sorted array:", arr)
}
