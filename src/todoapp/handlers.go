package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index request handler for "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex request handler for "/todos"
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Learn GO"},
		Todo{Name: "Learn to write API by GO!"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoDetail request handler for "/todos/{todoID}"
func TodoDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["todoID"]
	fmt.Fprintln(w, "Todo detail:", todoID)
}
