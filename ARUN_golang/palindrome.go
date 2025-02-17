package main
import "fmt"
func isPalindrome(n int) bool {
	original, reversed := n, 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}
func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Println(num, "is a palindrome number.")
	} else {
		fmt.Println(num, "is not a palindrome number.")
	}
}
