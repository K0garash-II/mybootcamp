package functions

import (
	"10.09task/methods"
	"errors"
)

func DeleteAccount(accounts map[int]methods.Account, id int, password int) error {
	account, ok := accounts[id]
	if !ok {
		return errors.New("счет не найден")
	}
	if password != account.Password {
		return errors.New("неверный пароль")
	}
	delete(accounts, id)
	return nil
}