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

func main() {
	accounts := []Account{
		{ID: 1, Owner: "Egor", Balance: 1000},
		{ID: 2, Owner: "Bob", Balance: 1500},
		{ID: 3, Owner: "Tom", Balance: 2000},
	}
	http.HandleFunc("/accounts", accountsHandler(accounts))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}

}
