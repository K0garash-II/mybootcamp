
package main

import "testing"

func TestProcessDuplicatePayment(t *testing.T) {
	bank := &Banc{
		Accounts: map[int]*Account{
			1: {ID: 1, Owner: "Egor", Balance: 1000},
			2: {ID: 2, Owner: "Alex", Balance: 500},
		},
		Processed: make(map[int]Result),
	}

	payment := Payment{
		ID:     1,
		From:   1,
		To:     2,
		Amount: 200,
	}

	firstResult, firstErr := bank.Process(&payment)

	if firstErr != nil {
		t.Fatalf("первый перевод завершился ошибкой: %v", firstErr)
	}

	if !firstResult.Success {
		t.Fatal("первый перевод должен быть успешным")
	}

	if bank.Accounts[1].Balance != 800 {
		t.Fatalf("баланс отправителя: ожидали 800, получили %v",
			bank.Accounts[1].Balance)
	}

	if bank.Accounts[2].Balance != 700 {
		t.Fatalf("баланс получателя: ожидали 700, получили %v",
			bank.Accounts[2].Balance)
	}

	secondResult, secondErr := bank.Process(&payment)

	if secondErr != nil {
		t.Fatalf("повторный вызов завершился ошибкой: %v", secondErr)
	}

	if !secondResult.Success {
		t.Fatal("повторный вызов должен вернуть успешный результат")
	}

	if bank.Accounts[1].Balance != 800 {
		t.Fatalf("повторное списание: баланс отправителя стал %v",
			bank.Accounts[1].Balance)
	}

	if bank.Accounts[2].Balance != 700 {
		t.Fatalf("повторное зачисление: баланс получателя стал %v",
			bank.Accounts[2].Balance)
	}

	if firstResult.PaymentID != secondResult.PaymentID ||
		firstResult.Success != secondResult.Success {
		t.Fatal("результат повторного вызова отличается от первого")
	}
}