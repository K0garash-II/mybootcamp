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
	ID      int
	Owner   string
	Balance float64
}
type Banc struct {
	Accounts  map[int]*Account
	mu        sync.Mutex
	Processed map[int]Result
}
type Result struct {
	PaymentID int
	Success   bool
	Error     error
}

func produce(payments chan<- Payment) {
	var id int = 1
	var from int
	var to int
	var amount float64
	var counter int
	var max int = 5
	for counter < max {
		fmt.Println("Откуда хотите отправить перевод?")
		fmt.Scan(&from)
		fmt.Println("Куда хотите отправить перевод?")
		fmt.Scan(&to)
		fmt.Println("Введите сумму перевода")
		fmt.Scan(&amount)

		payment := Payment{
			ID:     id,
			From:   from,
			To:     to,
			Amount: amount,
		}

		payments <- payment
		id++
		counter++
	}
	close(payments)
}

func (b *Banc) Process(p *Payment) (Result, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	res, ok := b.Processed[p.ID]
	if ok {
		return res, nil
	}

	accFrom, ok := b.Accounts[p.From]
	if !ok {
		return Result{}, errors.New("аккаунт не найден")
	}
	accTo, ok := b.Accounts[p.To]
	if !ok {
		return Result{}, errors.New("аккаунт не найден")
	}

	if p.Amount <= 0 {
		return Result{}, errors.New("сумма перевода должна бьть больше нуля")
	}

	if accFrom.Balance < p.Amount {
		return Result{}, errors.New("недостаточно средств")
	}

	if p.From == p.To {
		return Result{}, errors.New("id отправителя и получателя совпадают")
	}

	accFrom.Balance -= p.Amount
	b.Accounts[p.From] = accFrom

	accTo.Balance += p.Amount
	b.Accounts[p.To] = accTo
	
	result := Result{
		PaymentID: p.ID,
		Success: true,
		Error: nil,
	}

	b.Processed[p.ID] = result

	return result, nil

}


func worker(payments <-chan Payment, results chan<- Result, b *Banc, wg *sync.WaitGroup) {
	defer wg.Done()
	for payment := range payments {
		res, err := b.Process(&payment)
		if err != nil {
			results <- Result{
				PaymentID: payment.ID,
				Success:   false,
				Error:     err,
			}
			continue
		}
		results <- res
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	account := &Banc{
		Processed: make(map[int]Result),
		Accounts: map[int]*Account{
			1: {ID: 1, Owner: "Egor", Balance: 10000},
			2: {ID: 2, Owner: "Alex", Balance: 5000},
			3: {ID: 3, Owner: "Bob", Balance: 6000},
		},
	}
	payments := make(chan Payment, 10)
	results := make(chan Result, 10)


	go produce(payments)

	go worker(payments, results, account, &wg)
	go worker(payments, results, account, &wg)
	go worker(payments, results, account, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	var readerWG sync.WaitGroup
	readerWG.Add(1)

	go func() {
		defer readerWG.Done()
		for res := range results {
			fmt.Println(res)
		}
	}()

	readerWG.Wait()
}
