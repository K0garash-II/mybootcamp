package main

import (
	"errors"
	"fmt"
)

type Payer interface {
	Pay(amount float64) error
}

type Card struct {
	Number  string
	Balance float64
}

type BankAccount struct {
	Number  string
	Balance float64
}

type CryptoWallet struct {
	Number  string
	Balance float64
}

type GiftCard struct {
	Number  string
	Balance float64
}

func (c *Card) Pay(amount float64) error {
	if c.Balance < amount {
		return errors.New("Недостаточно средств")
	}
	c.Balance -= amount
	fmt.Println("Успешная оплата с карты")
	return nil
}

func (b *BankAccount) Pay(amount float64) error {
	if b.Balance < amount {
		return errors.New("Недостаточно средств")
	}
	b.Balance -= amount
	fmt.Println("Успешная оплата с банковского счета")
	return nil
}

func (c *CryptoWallet) Pay(amount float64) error {
	var commission float64
	commission = (amount / 100) * 5
	if c.Balance < amount+commission {
		return errors.New("Недостаточно средств")
	}
	c.Balance -= amount + commission
	fmt.Println("Успешная оплата криптокошельком")
	fmt.Printf("Комиссия за перевод составила %0.2fр\n", commission)
	return nil
}

func (g *GiftCard) Pay(amount float64) error {
	if amount < 100 {
		return errors.New("Минимальная сумма для оплаты не достигнута")
	}
	if g.Balance < amount {
		return errors.New("Недостаточно средств")
	}
	g.Balance -= amount
	fmt.Println("Успешная оптлата с подарочной карты")
	return nil
}

func processPayment(p Payer, amount float64) error {
	return p.Pay(amount)
}

func main() {
	card := Card{
		Balance: 10000,
	}
	bank := BankAccount{
		Balance: 5000,
	}
	crypto := CryptoWallet{
		Balance: 1000,
	}

	gift := GiftCard{
		Balance: 1000,
	}

	err := processPayment(&card, 5000)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}

	err = processPayment(&bank, 6000)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}

	err = processPayment(&crypto, 1000)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}

	err = processPayment(&crypto, 800)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}

	err = processPayment(&gift, 80)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}

	err = processPayment(&gift, 1100)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}

	err = processPayment(&gift, 500)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
}
