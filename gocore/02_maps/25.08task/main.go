package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID       int
	Owner    string
	Password int
	Balance  float64
	History  []string
}

func logInAccount(accounts map[int]Account, id int) (Account, error) {
	account, ok := accounts[id]
	if !ok {
		return Account{}, errors.New("cчет не найден")
	}
	return account, nil
}

func addAccount(accounts map[int]Account, account Account) {
	accounts[account.ID] = account
}

func deleteAccount(accounts map[int]Account, id int) error {
	_, ok := accounts[id]
	if !ok {
		return errors.New("cчет не найден")
	}
	delete(accounts, id)
	return nil
}

func replenishAccount(accounts map[int]Account, id int, balance float64) error {
	account, ok := accounts[id]
	account.Balance += balance
	accounts[id] = account
	if !ok {
		return errors.New("cчет не найден")
	}
	return nil
}

func debitFromAccounts(accounts map[int]Account, id int, balance float64) error {
	account, ok := accounts[id]
	account.Balance -= balance
	accounts[id] = account
	if !ok {
		return errors.New("cчет не найден")
	}
	return nil
}

func main() {
	accounts := make(map[int]Account)
	for {
		fmt.Println("| ========== БАНК ========== |")
		fmt.Println("| 1. Войти в личный кабинет  |")
		fmt.Println("| 2. Открыть счет            |")
		fmt.Println("| 3. Закрыть счет            |")
		fmt.Println("| 0. Выход                   |")

		var owner string
		var password int
		var password2 int

		var ans int
		fmt.Scan(&ans)

		for {
			if ans == 1 {
				var idToFind int
				var passwordToFind int
				var err error
				var balanceToPush float64
				fmt.Println("Введите ID счета: ")
				fmt.Scan(&idToFind)
				account, err := logInAccount(accounts, idToFind)
				if err != nil {
					fmt.Println("Ошибка: ", err)
					break
				} else {
					fmt.Println("Введите пароль")
					fmt.Scan(&passwordToFind)
					if account.Password != passwordToFind {
						fmt.Println("Ошибка, неверный пароль")
						break
					} else {
						for {
							var ans int

							fmt.Println(" ===== личный кабинет ===== ")
							fmt.Println(" 1. Пополнить баланс        ")
							fmt.Println(" 2. Перевод между счетами   ")
							fmt.Println(" 3. Показать баланс счета")
							fmt.Println(" 0. Выход                   ")
							fmt.Scan(&ans)

							if ans == 1 {
								fmt.Println("Введите сумму, которую хотите внести: ")
								fmt.Scan(&balanceToPush)
								err = replenishAccount(accounts, idToFind, balanceToPush)
								fmt.Println("Баланс успешно пополнен")
							}
							for {
								if ans == 2 {
									var id2 int
									var balanceToDebit float64
									var err2 error
									fmt.Println("Введите сумму, которую хотите перевести: ")
									fmt.Scan(&balanceToDebit)
									account, err = logInAccount(accounts, idToFind)
									if account.Balance >= balanceToDebit {
										err = debitFromAccounts(accounts, idToFind, balanceToDebit)
									} else {
										fmt.Println("Недостаточно средств")
										break
									}
									fmt.Println("Введите ID счета, куда хотите перевести")
									fmt.Scan(&id2)
									err2 = replenishAccount(accounts, id2, balanceToDebit)
									if err2 != nil {
										fmt.Printf("Счет с ID %v не найден\n", id2)
										break
									} else {
										fmt.Println("Успешный перевод")
									}
								}
								break
							}

							if ans == 3 {
								account, err = logInAccount(accounts, idToFind)
								if err == nil {
									fmt.Printf(" Текущий баланс счета - %v\n", account.Balance)
								}
							}

							if ans == 0 {
								break
							}
						}
					}
				}
			}
			break
		}

		for {
			if ans == 2 {
				id := 1
				for _, p := range accounts {
					if p.ID >= id {
						id = p.ID + 1
					}
				}
				fmt.Println("Введите имя держателя: ")
				fmt.Scan(&owner)
				fmt.Println("Придумайте пароль: ")
				fmt.Scan(&password)
				fmt.Println("Введите пароль еще раз: ")
				fmt.Scan(&password2)
				if password != password2 {
					fmt.Println("Ошибка, пароли не совпадают")
					break
				}
				balance := 0.0
				account := Account{ID: id, Owner: owner, Password: password, Balance: balance}
				addAccount(accounts, account)
				fmt.Printf("Счет успешно зарегестрирован под ID %v!\n", id)
				fmt.Printf("Ваш баланс %g\n", balance)
			}
			break
		}

		for {
			if ans == 3 {
				var iddel int
				var passwordToFind int
				fmt.Println("Введите id счета, который хотите закрыть: ")
				fmt.Scan(&iddel)
				account, err := logInAccount(accounts, iddel)
				if err != nil {
					fmt.Println("Ошибка: ", err)
					break
				} else {
					fmt.Println("Введите пароль от счета, который хотите закрыть")
					fmt.Scan(&passwordToFind)
					if account.Password != passwordToFind {
						fmt.Println("Ошибка, неверный пароль")
					} else {
						err = deleteAccount(accounts, iddel)
						if err != nil {
							fmt.Println("Ошибка: ", err)
							break
						} else {
							fmt.Printf("Счет с id %v успешно закрыт\n", iddel)
						}
					}
				}
			}
			break
		}

		if ans == 0 {
			break
		}
	}
}
