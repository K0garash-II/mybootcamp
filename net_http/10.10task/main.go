package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go!")
}

func accountsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
    	http.Error(w, "Missing account id", http.StatusBadRequest)
    	return
	}
	fmt.Fprintf(w, "Requested account: %s", id)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/accounts", accountsHandler)

	fmt.Println("Server started at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
