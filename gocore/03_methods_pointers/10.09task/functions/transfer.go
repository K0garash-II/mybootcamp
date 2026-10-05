package functions

import (
	"10.09task/methods"
	"errors"
	"fmt"
)

func Transfer(accounts map[int]methods.Account, id1, id2 int, balance float64) error {
	account1 := accounts[id1]
	account2, ok := accounts[id2]
	if !ok {
		return errors.New("счет не найден")
	}
	if balance > account1.Balance {
		return errors.New("недостаточно средств")
	}
	if balance <= 0 {
		return errors.New("невалидный ввод")
	}
	if account1.ID == account2.ID {
		return errors.New("некорректный ID")
	}
	account1.Balance -= balance
	account1.History = append(account1.History, fmt.Sprintf("Перевод: -%.2f\n", balance))
	accounts[id1] = account1

	account2.Balance += balance
	account2.History = append(account2.History, fmt.Sprintf("Пополнение: +%.2f\n", balance))
	accounts[id2] = account2
	
	return nil
}
