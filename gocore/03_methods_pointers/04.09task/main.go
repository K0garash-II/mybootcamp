package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID      int
	Owner   string
	Balance float64
	History []string
}

func (a *Account) deposit(amount float64) {
	a.Balance += amount
	a.History = append(a.History, fmt.Sprintf("Пополнение: +%.2f\n", amount))
}

func (a *Account) withdraw(amount float64) error {
	if a.Balance < amount {
		return errors.New("недостаточно средств")
	}
	if amount <= 0 {
		return errors.New("невалидный ввод")
	}
	a.Balance -= amount
	a.History = append(a.History, fmt.Sprintf("Снятие: -%.2f\n", amount))
	return nil
}

func (a Account) showBalance() {
	fmt.Printf("Владелец: %s\n Баланс: %.2f\n", a.Owner, a.Balance)
}

func (a Account) showHistory() {
	fmt.Printf("История операций по счету: %s\n", a.History)
}

func (a *Account) transfer(to *Account, amount float64) error {
	if a.Balance < amount {
		return errors.New("недостаточно средств")
	}
	if amount <= 0 {
		return errors.New("невалидный ввод")
	}
	a.Balance -= amount
	a.History = append(a.History, fmt.Sprintf("Перевод: -%.2f\n", amount))
	to.Balance += amount
	to.History = append(to.History, fmt.Sprintf("Пополнение: +%.2f\n", amount))
	return nil
}

func main() {
	accounts := []Account{
		{ID: 1, Owner: "Иван", Balance: 500},
		{ID: 2, Owner: "Пётр", Balance: 1000},
	}
	for {
		var ans int

		if ans < 0 || ans > 5 {
			fmt.Println("Некорректный ввод")
			break
		}

		fmt.Println("| ======== БАНК ======== |")
		fmt.Println("| 1. Показать баланс     |")
		fmt.Println("| 2. Пополнить баланс    |")
		fmt.Println("| 3. Снять деньги        |")
		fmt.Println("| 4. Перевести деньги    |")
		fmt.Println("| 5. Показать историю    |")
		fmt.Println("| 0. Выход               |")

		fmt.Scan(&ans)

		if ans == 1 {
			var ans int
			fmt.Println("Введите ID счета")
			fmt.Scan(&ans)
			for i := range accounts {
				if accounts[i].ID == ans {
					accounts[i].showBalance()
				}
			}
		}

		if ans == 2 {
			var ans int
			fmt.Println("Введите ID счета")
			fmt.Scan(&ans)
			for i := range accounts {
				if accounts[i].ID == ans {
					var toDep float64
					fmt.Println("Введите сумму пополнения:")
					fmt.Scan(&toDep)
					accounts[i].deposit(toDep)
					fmt.Printf("Аккаунт с ID %d пополнен на сумму %.2f\n", ans, toDep)
					fmt.Printf("Текущий баланс счета: %.2f\n", accounts[i].Balance)
				} 
			}
		}

		if ans == 3 {
			var ans int
			fmt.Println("Введите ID счета")
			fmt.Scan(&ans)
			for i := range accounts {
				if accounts[i].ID == ans {
					var toDraw float64
					fmt.Println("Введите сумму для снятия:")
					fmt.Scan(&toDraw)
					accounts[i].withdraw(toDraw)
					fmt.Printf("Успешное снятие со счета %d на сумму %.2f\n", ans, toDraw)
					fmt.Printf("Текущий баланс счета: %.2f\n", accounts[i].Balance)
				} 
			}
		}

		if ans == 4 {
			var ans1 int
			fmt.Println("Введите ID счета, с которого хотите перевести деньги")
			fmt.Scan(&ans1)
			for i := range accounts {
				if accounts[i].ID == ans1 {
					var ans2 int
					fmt.Println("Введите ID счета, куда хотите перевести деньги")
					fmt.Scan(&ans2)
					for q := range accounts {
						if accounts[q].ID == ans2 && ans2 != ans1 {
							var toTrans float64
							fmt.Println("Введите сумму перевода: ")
							fmt.Scan(&toTrans)
							err := accounts[i].transfer(&accounts[q], toTrans)
							if err != nil {
								fmt.Println("Ошибка: ", err)
							}
						}
					}
				} 
			}
		}

		if ans == 5 {
			var ans int
			fmt.Println("Введите ID счета")
			fmt.Scan(&ans)
			for i := range accounts {
				if accounts[i].ID == ans {
					accounts[i].showHistory()
				}
			}
		}

		if ans == 0 {
			break
		}
	}
}
