package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Account struct {
	ID       int
	Owner    string
	Password string
	Balance  float64
}

func handlerDeposit(accounts map[int]Account) http.HandlerFunc {
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

		account, ok := accounts[idInt]
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
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

		account.Balance += amountFloat
		fmt.Fprintf(w, "Account %v replenished. Balance: %0.2f\n", id, account.Balance)

		accounts[account.ID] = account
	}
}

func handlerRePass(accounts map[int]Account) http.HandlerFunc {
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

		account, ok := accounts[idInt]
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		pass := r.URL.Query().Get("password")
		if pass != account.Password {
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
			return
		}

		passNew := r.URL.Query().Get("new_password")
		if strings.TrimSpace(passNew) == "" {
			http.Error(w, "Invalid new password", http.StatusBadRequest)
			return
		}
		account.Password = passNew
		accounts[idInt] = account
		fmt.Fprintf(w, "%s, your password changed", account.Owner)
	}
}

func handlerGetAccountInfo(accounts map[int]Account) http.HandlerFunc {
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

		account, ok := accounts[idInt]
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		pass := r.URL.Query().Get("password")
		if pass != account.Password {
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "ID: %v\n Name: %s\n Password: %s\n Balance: %0.2f", account.ID, account.Owner, account.Password, account.Balance)
	}
}

func handlerDeleteAccount(accounts map[int]Account) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid account id", http.StatusBadRequest)
			return
		}

		account, ok := accounts[idInt]
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		pass := r.URL.Query().Get("password")
		if pass != account.Password {
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "%s, your account with id %v deleted", account.Owner, account.ID)
		delete(accounts, idInt)
	}
}

func main() {
	accounts := map[int]Account{
		1: {ID: 1, Owner: "Tom", Password: "1a", Balance: 1000},
		2: {ID: 2, Owner: "Bob", Password: "2b", Balance: 2000},
		3: {ID: 3, Owner: "Rey", Password: "3c", Balance: 3000},
	}

	http.HandleFunc("/accounts/deposit", handlerDeposit(accounts))
	http.HandleFunc("/accounts/repassword", handlerRePass(accounts))
	http.HandleFunc("/accounts", handlerGetAccountInfo(accounts))
	http.HandleFunc("/accounts/delete", handlerDeleteAccount(accounts))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}
