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
	num, err := strconv.Atoi(params["number"])
	if err != nil {
		http.Error(w, "Invalid number format", http.StatusBadRequest)
		return
	}
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

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/numbers", getNumbers).Methods("GET")
	r.HandleFunc("/numbers/{number}", addNumber).Methods("POST")
	r.HandleFunc("/sum", getSum).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Println("Server is running on port 2007...")
	log.Fatal(http.ListenAndServe(":2007", r))
}
