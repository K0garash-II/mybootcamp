package main

import (
	"errors"
	"fmt"
)

type Payer interface {
	Push(id int, amount float64) error
}

type Adder interface {
	Add(id int, owner string, password int)
}

type DebitCard struct {
	Owner    string
	Password int
	ID       int
	Balance  float64
}
type DebitcardStore struct {
	cards map[int]DebitCard
}

type CreditCard struct {
	Owner    string
	Password int
	ID       int
	Balance  float64
}
type CreditCardStore struct {
	cards map[int]CreditCard
}

type CryptoWallet struct {
	Owner    string
	Password int
	ID       int
	Balance  float64
}
type CryptoWalletStore struct {
	wallets map[int]CryptoWallet
}

func (d *DebitcardStore) Push(id int, amount float64) error {
	card, ok := d.cards[id]
	if !ok {
		return errors.New("Карта не найдена")
	}

	card.Balance += amount
	d.cards[id] = card

	fmt.Printf("Дебетовая карта успешно пополнена!\n")

	return nil
}
func (c *CreditCardStore) Push(id int, amount float64) error {
	card, ok := c.cards[id]
	if !ok {
		return errors.New("Карта не найдена")
	}

	commission := 0.0
	if amount >= 1000 {
		commission = amount * 0.01
	}

	card.Balance += amount - commission
	c.cards[id] = card

	fmt.Printf("Кредитная карта успешно пополнена!\n")
	fmt.Printf("Комиссия за пополнение составила %0.2fр\n", commission)

	return nil
}
func (w *CryptoWalletStore) Push(id int, amount float64) error {
	wallet, ok := w.wallets[id]
	if !ok {
		return errors.New("Кошелек не найден")
	}

	commission := 0.0

	var ans int
	fmt.Println("Выберете сеть для перевода")
	fmt.Println("1. ETH")
	fmt.Println("2. SOL")
	fmt.Println("3. TRC")

	fmt.Scan(&ans)
	if ans == 1 {
		commission = amount * 0.05
	}
	if ans == 2 {
		commission = amount * 0.01
	}
	if ans == 3 {
		commission = amount * 0.02
	}

	wallet.Balance += amount - commission
	w.wallets[id] = wallet

	fmt.Printf("Кошелек успешно пополнен!\n")
	fmt.Printf("Комиссия за пополнение составила %0.2fр\n", commission)

	return nil
}
func Pusher(p Payer, id int, amount float64) error {
	return p.Push(id, amount)
}

func (d *DebitcardStore) Add(id int, owner string, password int) {
	d.cards[id] = DebitCard{
		ID:       id,
		Owner:    owner,
		Password: password,
		Balance:  0,
	}
}
func (c *CreditCardStore) Add(id int, owner string, password int) {
	c.cards[id] = CreditCard{
		ID:       id,
		Owner:    owner,
		Password: password,
		Balance:  0,
	}
}
func (w *CryptoWalletStore) Add(id int, owner string, password int) {
	w.wallets[id] = CryptoWallet{
		ID:       id,
		Owner:    owner,
		Password: password,
		Balance:  0,
	}
}
func addAny(a Adder, id int, owner string, password int) {
	a.Add(id, owner, password)
}

func openAny() (string, int, error) {
	var owner string
	var password int
	var password2 int
	fmt.Println("Введите имя держателя")
	fmt.Scan(&owner)
	fmt.Println("Придумайте пароль")
	fmt.Scan(&password)
	fmt.Println("Повторите пароль")
	fmt.Scan(&password2)
	if password != password2 {
		return " ", 0, errors.New("пароли не совпадают")
	}
	return owner, password, nil
}

func main() {

	debitStore := DebitcardStore{
		cards: make(map[int]DebitCard),
	}
	creditStore := CreditCardStore{
		cards: make(map[int]CreditCard),
	}
	walletStore := CryptoWalletStore{
		wallets: make(map[int]CryptoWallet),
	}

	var ans int
	for {
		fmt.Println("1. Открыть дебетовую карту")
		fmt.Println("2. Открыть кредитную карту")
		fmt.Println("3. Открыть криптокошелек")
		fmt.Println("4. Пополнить дебетовую карту")
		fmt.Println("5. Пополнить кредитную карту")
		fmt.Println("6. Пополнить криптокошелек")
		fmt.Scan(&ans)

		if ans == 1 {
			id := 1
			for _, p := range debitStore.cards {
				if p.ID >= id {
					id = p.ID + 1
				}
			}
			owner, password, err := openAny()
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}
			addAny(&debitStore, id, owner, password)
			fmt.Printf("Дебетовая карта успешно зарегестрирована под id %v\n", id)

		} else if ans == 2 {
			id := 1
			for _, p := range creditStore.cards {
				if p.ID >= id {
					id = p.ID + 1
				}
			}
			owner, password, err := openAny()
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}
			addAny(&creditStore, id, owner, password)
			fmt.Printf("Кредитная карта успешно зарегестрирована под id %v\n", id)

		} else if ans == 3 {
			id := 1
			for _, p := range walletStore.wallets {
				if p.ID >= id {
					id = p.ID + 1
				}
			}
			owner, password, err := openAny()
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}
			addAny(&walletStore, id, owner, password)
			fmt.Printf("Криптокошелек успешно зарегестрирован под id %v\n", id)

		} else if ans == 4 {
			var id int
			var amount float64
			fmt.Println("Введите id дебетовой карты, которую хотите пополнить")
			fmt.Scan(&id)
			fmt.Println("Введите сумму пополнения")
			fmt.Scan(&amount)
			err := Pusher(&debitStore, id, amount)
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}

		} else if ans == 5 {
			var id int
			var amount float64
			fmt.Println("Введите id кредитной карты, которую хотите пополнить")
			fmt.Scan(&id)
			fmt.Println("Введите сумму пополнения")
			fmt.Scan(&amount)
			err := Pusher(&creditStore, id, amount)
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}

		} else if ans == 6 {
			var id int
			var amount float64
			fmt.Println("Введите id кошелька, который хотите пополнить")
			fmt.Scan(&id)
			fmt.Println("Введите сумму пополнения")
			fmt.Scan(&amount)
			err := Pusher(&walletStore, id, amount)
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}
		} else if ans == 0 {
			break
		}
	}
}
