package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Account struct {
	ID      int
	Owner   string
	Balance float64
}

func accountsHandler(accounts []Account) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		id := r.URL.Query().Get("id")
		if id == "" {
			for i := range accounts {
				fmt.Fprintln(w, accounts[i])
			}
			return 
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return 
		}
		if id != "" {
			var found bool = false
			for i := range accounts {
				if accounts[i].ID == idInt {
					found = true
					fmt.Fprintf(w, "Requested account: %v", accounts[i])
					return 
				} 
			}
			if !found {
				http.Error(w, "Not Found", http.StatusNotFound)
			}
			return 
		}
	}
}

func depositHandler(accounts []Account) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not alllowed", http.StatusMethodNotAllowed)
			return 
		}

		id := r.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return 
		}

		amount := r.URL.Query().Get("amount")
		amountFloat, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			http.Error(w, "Invalid account amount", http.StatusBadRequest)
			return 
		}

		if amountFloat <= 0 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return 
		}

		var found bool = false
		for i := range accounts {
			if accounts[i].ID == idInt {
				found = true
				accounts[i].Balance += amountFloat
				fmt.Fprintf(w, "Account %v replenished. Balance: %0.2f\n", id, accounts[i].Balance)
				return 
			} 
		}
		if !found {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
	}
}

func main() {
	accounts := []Account{
		{ID: 1, Owner: "Egor", Balance: 1000},
		{ID: 2, Owner: "Bob", Balance: 1500},
		{ID: 3, Owner: "Tom", Balance: 2000},
	}
	http.HandleFunc("/accounts", accountsHandler(accounts))
	http.HandleFunc("/accounts/deposit", depositHandler(accounts))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}

}
