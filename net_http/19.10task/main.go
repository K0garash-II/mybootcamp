package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func handlerPush(b *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return
		}
		if idInt <= 0 {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
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

		b.mu.Lock()

		account, ok := b.accounts[idInt]
		if !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			b.mu.Unlock()
			return 
		}

		account.Balance += amountFloat
		b.accounts[idInt] = account

		b.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(account)
		if err != nil {
			fmt.Println("Failed to encode response:", err)
			return
		}
	}
}

func handlerGetAcc(b *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return
		}
		if idInt <= 0 {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return 
		}

		b.mu.Lock()

		account, ok := b.accounts[idInt]
		if !ok {
			http.Error(w, "Not found", http.StatusNotFound)
			b.mu.Unlock()
			return 
		}

		b.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

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
	http.HandleFunc("/accounts/deposit", handlerPush(&bank))
	http.HandleFunc("/accounts/info", handlerGetAcc(&bank))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
