package main

import "fmt"

type Account struct {
    Owner   string
    Balance float64
}

func (a Account) showBalance() {
	fmt.Println(a.Balance)
}

func (a *Account) deposit(amount float64) {
	a.Balance += amount
}

func main() {
	account := Account{
    	Owner:   "Ivan",
    	Balance: 1000,
	}
	account.showBalance()
	account.deposit(500)
	account.showBalance()
}