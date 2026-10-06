package main

import (
	"errors"
	"fmt"
	"sync"
)

type Payment struct {
	ID     int
	From   int
	To     int
	Amount float64
}

type Account struct {
	ID int
	Owner string
	Balance float64
}

type BankProcessor struct {
	Accounts map[int]Account
}

type Processor interface {
	Process(Payment) error 
}

func(b *BankProcessor) Process(p Payment) error {
	accountFrom, ok := b.Accounts[p.From]
	if !ok {
		return errors.New("карта не найдена")
	}
	
	accountTo, ok := b.Accounts[p.To]
	if !ok {
		return errors.New("карта не найдена")
	}

	if p.Amount <= 0 {
    	return errors.New("сумма должна быть больше нуля")
	}

	if accountFrom.Balance < p.Amount {
		return errors.New("недостаточно средств")
	}

	if p.From == p.To {
		return errors.New("id отправителя и получателя совпадают")
	}

	accountFrom.Balance -= p.Amount
	b.Accounts[p.From] = accountFrom
	
	accountTo.Balance += p.Amount
	b.Accounts[p.To] = accountTo

	return nil
}

func worker(processor Processor, payment Payment, wg *sync.WaitGroup) {
	err := processor.Process(payment)
	defer wg.Done()
	
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Успешный перевод")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	processor := &BankProcessor{
        Accounts: map[int]Account{
            1: {ID: 1, Owner: "Egor", Balance: 10000},
            2: {ID: 2, Owner: "Alex", Balance: 5000},
			3: {ID: 3, Owner: "Bob", Balance: 6000},
        },
    }

	payment1 := Payment{
		ID: 1,
		To: 2,
		From: 1,
		Amount: 100,
	}
	payment2 := Payment{
		ID: 2,
		To: 1,
		From: 3,
		Amount: 200,
	}
	payment3 := Payment{
		ID: 1,
		To: 3,
		From: 2,
		Amount: 300,
	}

	go worker(processor, payment1, &wg)
	go worker(processor, payment2, &wg)
	go worker(processor, payment3, &wg)

	wg.Wait()
}