package functions

import (
	"10.09task/methods"
	"fmt"
)

func PushToAccount(accounts map[int]methods.Account, id int, balance float64) {
	account := accounts[id]
	account.Balance += balance
	account.History = append(account.History, fmt.Sprintf("Пополнение: +%.2f\n", balance))
	accounts[id] = account
}
