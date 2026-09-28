package main

import (
	"errors"
	"fmt"
)

type Debit struct {
	ID           int
	Owner        string
	Password     int
	DebitBalance float64
	History      []string
}
type Credit struct {
	ID            int
	Owner         string
	Password      int
	CreditBalance float64
	History       []string
}

type deposit interface {
	deposit()
}
type withdraw interface {
	withdraw()
}
type add interface {
	add()
}

func (d *Debit) add(debits map[int]Debit, debit Debit) {
	debits[debit.ID] = debit
}
func (c *Credit) add(credits map[int]Credit, credit Credit) {
	credits[credit.ID] = credit
	c.CreditBalance = 100000
}

func (d *Debit) deposit(dep float64) {
	d.DebitBalance += dep
	d.History = append(d.History, fmt.Sprintf("Пополнение: +%.2f\n", dep))
}
func (c *Credit) deposit(dep float64, to *Credit) {
	c.CreditBalance += dep
	c.History = append(c.History, fmt.Sprintf("Пополнение: +%.2f\n", dep))

	var commission float64
	commission = dep / 100

	c.CreditBalance -= commission
	c.History = append(c.History, fmt.Sprintf("Комиссия за перевод: -%.2f\n", commission))

	to.CreditBalance += commission
}

func (d *Debit) transfer(to *Debit, amount float64) error {
	if d.DebitBalance < amount {
		return errors.New("недостаточно средств")
	}
	if amount <= 0 {
		return errors.New("невалидный ввод")
	}
	d.DebitBalance -= amount
	d.History = append(d.History, fmt.Sprintf("Перевод: -%.2f\n", amount))

	to.DebitBalance += amount
	to.History = append(to.History, fmt.Sprintf("Пополнение: +%.2f\n", amount))
	return nil
}

func main() {
	//debits := make(map[int]Debit)
	//credits := make(map[int]Credit)
	//credits = map[int](Credit){
	//	999: {ID: 999, Owner: "Egor", CreditBalance: 0, History: []string{}},
	//}

}
