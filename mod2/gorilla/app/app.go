package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customer", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCostumer).Methods(http.MethodPost)

	router.HandleFunc("/customer/{customer_id:[0-9]*}", getCostumer).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
