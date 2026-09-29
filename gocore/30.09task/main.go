package main

import (
	"errors"
	"fmt"
)

type Payer interface {
	Push(id int, amount float64) error
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

	fmt.Printf("Ваша дебетовая карта успешно пополнена!\n")

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

	fmt.Printf("Ваша кредитная карта успешно пополнена!\n")
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

	fmt.Printf("Ваша кошелек успешно пополнен!\n")
	fmt.Printf("Комиссия за пополнение составила %0.2fр\n", commission)

	return nil
}

func Pusher(p Payer, id int, amount float64) error {
	return p.Push(id, amount)
}

func main() {

	debitStore := DebitcardStore{
		cards: make(map[int]DebitCard),
	}
	debitStore.cards[1] = DebitCard{
		ID:       1,
		Owner:    "Egor",
		Balance:  0,
		Password: 123,
	}

	creditStore := CreditCardStore{
		cards: make(map[int]CreditCard),
	}
	creditStore.cards[2] = CreditCard{
		ID:       2,
		Owner:    "Egor",
		Balance:  0,
		Password: 234,
	}

	walletStore := CryptoWalletStore{
		wallets: make(map[int]CryptoWallet),
	}
	walletStore.wallets[3] = CryptoWallet{
		ID:       3,
		Owner:    "Egor",
		Balance:  0,
		Password: 345,
	}

	var ans int

	fmt.Println("1. Пополнить дебетовую карту")
	fmt.Println("2. Пополнить кредитную карту")
	fmt.Println("3. Пополнить криптокошелек")
	fmt.Scan(&ans)

	if ans == 1 {
		var amount float64
		var id int
		fmt.Println("Введите сумму пополнения")
		fmt.Scan(&amount)
		fmt.Println("Введите id карты")
		fmt.Scan(&id)
		err := Pusher(&debitStore, id, amount)
		if err != nil {
			fmt.Println("Ошибка: ", err)
		}
	}

	if ans == 2 {
		var amount float64
		var id int
		fmt.Println("Введите сумму пополнения")
		fmt.Scan(&amount)
		fmt.Println("Введите id карты")
		fmt.Scan(&id)
		err := Pusher(&creditStore, id, amount)
		if err != nil {
			fmt.Println("Ошибка: ", err)
		}
	}

	if ans == 3 {
		var amount float64
		var id int
		fmt.Println("Введите сумму пополнения")
		fmt.Scan(&amount)
		fmt.Println("Введите id криптокошелька")
		fmt.Scan(&id)
		err := Pusher(&walletStore, id, amount)
		if err != nil {
			fmt.Println("Ошибка: ", err)
		}
	}

}
