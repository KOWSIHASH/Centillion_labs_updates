package main
import "fmt"
func main() {
    numbers := [5]int{10, 20, 30, 40, 50}
    for i, num := range numbers {
        fmt.Printf("Element at index %d is %d\n", i, num)
    }
}
