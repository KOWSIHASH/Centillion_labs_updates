package main
import "fmt"
func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    if num > 0 {
        fmt.Println("Positive Number")
    } else if num < 0 {
        fmt.Println("Negative Number")
    } else {
        fmt.Println("Zero")
    }
}
