package main
import "fmt"
type Person struct {
    Name string
    Age  int
}
func main() {
    p := Person{Name: "Meganath", Age: 23}
    fmt.Println("Name:", p.Name, "Age:", p.Age)
}
