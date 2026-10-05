package main

import (
	"10.09task/functions"
	"10.09task/methods"
	"fmt"
)

func main() {
	accounts := make(map[int]methods.Account)
	for {
		var userAns int

		fmt.Println("| ========== БАНК ========== |")
		fmt.Println("| 1. Войти в личный кабинет  |")
		fmt.Println("| 2. Открыть счет            |")
		fmt.Println("| 3. Закрыть счет            |")
		fmt.Println("| Выход - Ctrl+C             |")
		fmt.Scan(&userAns)
		for {
			if userAns == 1 {
				var userID int
				var userPass int
				fmt.Println("Введите id аккаунта")
				fmt.Scan(&userID)
				fmt.Println("Введите пароль от аккаунта")
				fmt.Scan(&userPass)
				_, err := functions.LogInAccount(accounts, userID, userPass)
				if err != nil {
					fmt.Println("Ошибка: ", err)
					break
				} else {
					for {
						var ans int
						fmt.Println(" ==== личный кабинет ==== ")
						functions.ShowBalance(accounts, userID)
						fmt.Println(" 1. Пополнить баланс      ")
						fmt.Println(" 2. Перевод между счетами ")
						fmt.Println(" 3. Показать историю      ")
						fmt.Println(" 0. Выход                 ")
						fmt.Scan(&ans)
						if ans == 1 {
							var balanceToPush float64
							fmt.Println("Введите сумму пополнения: ")
							fmt.Scan(&balanceToPush)
							functions.PushToAccount(accounts, userID, balanceToPush)
							fmt.Println("Баланс успешно пополнен!")
						}
						for {
							if ans == 2 {
								var idToFind int
								var trans float64
								fmt.Println("Введите id счета, куда хотите перевести: ")
								fmt.Scan(&idToFind)
								fmt.Println("Введите сумму, которую хотите перевести: ")
								fmt.Scan(&trans)
								err := functions.Transfer(accounts, userID, idToFind, trans)
								if err != nil {
									fmt.Println("Ошибка: ", err)
									break
								} else {
									fmt.Println("Успешный перевод")
								}
							}
							break
						}
						if ans == 3 {
							functions.ShowHistory(accounts, userID)
						}
						if ans == 0 {
							break
						}
					}
				}
			}
			break
		}
		
		if userAns == 2 {
			var owner string
			var password int
			var password2 int
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
			account := methods.Account{ID: id, Owner: owner, Password: password, Balance: balance}
			functions.AddAccount(accounts, account)
			fmt.Printf("Счет успешно зарегестрирован под ID %v!\n", id)
			fmt.Printf("Ваш баланс %g\n", balance)
		}
		if userAns == 3 {
			var iddel int
			var passwordToFind int
			fmt.Println("Введите id счета, который хотите закрыть: ")
			fmt.Scan(&iddel)
			fmt.Println("Введите пароль от счета, который хотите закрыть")
			fmt.Scan(&passwordToFind)
			err := functions.DeleteAccount(accounts, iddel, passwordToFind)
			if err != nil {
				fmt.Println("Ошибка: ", err)
			} else {
				fmt.Printf("Счет с id %v успешно закрыт\n", iddel)
			}
		}
	}
}
