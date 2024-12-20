package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello, world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"MVL", "Paris", "1234567"},
		{"Firouzja", "Chartre", "1000000"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCostumer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Print(w, vars["customer_id"])
}

func createCostumer(w http.ResponseWriter, _ *http.Request) {
	fmt.Print(w, "Post request received")
}
