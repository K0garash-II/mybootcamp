package functions

import (
	"10.09task/methods"
	"fmt"
)

func ShowHistory(accounts map[int]methods.Account, id int) {
	account := accounts[id]
	fmt.Println("История операций по счету:")
	for i, v := range account.History {
		fmt.Printf("  %d. %s\n", i+1, v)
	}
}
