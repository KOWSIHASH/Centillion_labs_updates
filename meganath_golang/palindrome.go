package main
import "fmt"
func main() {
    var num, temp, rev int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    temp = num
    for temp > 0 {
        rev = rev*10 + temp%10
        temp /= 10
    }
    if num == rev {
        fmt.Println(num, "is a Palindrome")
    } else {
        fmt.Println(num, "is Not a Palindrome")
    }
}
