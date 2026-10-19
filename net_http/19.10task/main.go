package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Account struct {
	ID      int     `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type Bank struct {
	mu       sync.Mutex
	accounts map[int]Account
}

func handlerAdd(b *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var account Account

		err := json.NewDecoder(r.Body).Decode(&account)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if account.ID <= 0 || strings.TrimSpace(account.Owner) == "" || account.Balance < 0 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		b.mu.Lock()

		_, ok := b.accounts[account.ID]
		if ok {
			http.Error(w, "Account with this ID already exists", http.StatusConflict)
			b.mu.Unlock()
			return
		}

		b.accounts[account.ID] = account

		b.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(account)
		if err != nil {
			fmt.Println("Failed to encode response:", err)
			return
		}
	}
}

func main() {
	var mut sync.Mutex
	accounts := make(map[int]Account)
	bank := Bank{
		mu:       mut,
		accounts: accounts,
	}
	http.HandleFunc("/accounts", handlerAdd(&bank))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
