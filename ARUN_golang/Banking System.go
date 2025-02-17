
package main
import "fmt"
type Account struct {
	owner  string
	balance float64
}
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
		fmt.Println("Deposited:", amount)
	} else {
		fmt.Println("Invalid deposit amount!")
	}
}
func (a *Account) Withdraw(amount float64) {
	if amount > 0 && amount <= a.balance {
		a.balance -= amount
		fmt.Println("Withdrawn:", amount)
	} else {
		fmt.Println("Insufficient balance or invalid amount!")
	}
}
func (a *Account) CheckBalance() {
	fmt.Println("Balance:", a.balance)
}
func main() {
	acc := Account{owner: "John Doe", balance: 5000}
	fmt.Println("Welcome", acc.owner)
	acc.CheckBalance()
	acc.Deposit(2000)
	acc.Withdraw(1500)
	acc.CheckBalance()
}
