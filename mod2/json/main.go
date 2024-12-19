package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customer", getAllCustomers)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func greet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello, world")
}

func getAllCustomers(w http.ResponseWriter, _ *http.Request) {
	customers := []Customer{
		{"MVL", "Paris", "1234567"},
		{"Firouzja", "Chartre", "1000000"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
