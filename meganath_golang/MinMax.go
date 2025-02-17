package main
import "fmt"
func main() {
    arr := []int{23, 45, 67, 12, 89, 34}
    max, min := arr[0], arr[0]
    for _, num := range arr {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    fmt.Println("Maximum:", max)
    fmt.Println("Minimum:", min)
}
