package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var numbers []int

func getNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
}

func addNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	num, _ := strconv.Atoi(params["number"])
	numbers = append(numbers, num)
	json.NewEncoder(w).Encode(numbers)
}

func getSum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	json.NewEncoder(w).Encode(sum)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/numbers", getNumbers).Methods("GET")
	r.HandleFunc("/numbers/{number}", addNumber).Methods("POST")
	r.HandleFunc("/sum", getSum).Methods("GET")

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
