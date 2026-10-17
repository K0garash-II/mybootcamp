package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	ID      int     `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

func main() {
	account := Account{
		ID: 7,
		Owner: "Alex",
		Balance: 4500,
	}

	acc, err := json.Marshal(account)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(acc))

	data := []byte(`{"id":2,"owner":"Bob","balance":2500}`)

	var x Account

	err1 := json.Unmarshal(data, &x)
	if err1 != nil {
		fmt.Println("Error", err1)
		return
	}

	fmt.Println(x.Owner, x.Balance)
}