package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var numbers []int

func getNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(numbers)
	if err != nil {
		return
	}
}

func addNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	num, err := strconv.Atoi(params["number"])
	if err != nil {
		http.Error(w, "Invalid number format", http.StatusBadRequest)
		return
	}
	numbers = append(numbers, num)
	err = json.NewEncoder(w).Encode(numbers)
	if err != nil {
		return
	}
}

func getSum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	err := json.NewEncoder(w).Encode(sum)
	if err != nil {
		return
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/numbers", getNumbers).Methods("GET")
	r.HandleFunc("/numbers/{number}", addNumber).Methods("POST")
	r.HandleFunc("/sum", getSum).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe("127.0.0.1:2007", r))
}
