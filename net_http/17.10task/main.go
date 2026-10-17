package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Account struct {
	ID      int     `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		fmt.Println("Failed to encode response:", err)
    	return
	}
}

func main() {
	http.HandleFunc("/accounts", handlerPost)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}