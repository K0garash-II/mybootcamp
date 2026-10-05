package functions

import (
	"10.09task/methods"
	"errors"
)

func LogInAccount(accounts map[int]methods.Account, id int, password int) (methods.Account, error) {
	account, ok := accounts[id] 
	if !ok {
		return methods.Account{}, errors.New("счет не найден")
	}
	if password != account.Password {
		return methods.Account{}, errors.New("неверный пароль")
	}
	return account, nil
}