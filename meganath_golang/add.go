package main
import "fmt"
func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)
    sum := a + b
    fmt.Println("Sum:", sum)
}
