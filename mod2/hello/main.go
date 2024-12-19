package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}
