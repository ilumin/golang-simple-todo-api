package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoDetail)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index request handler for "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// TodoIndex request handler for "/todos"
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Learn GO"},
		Todo{Name: "Learn to write API by GO!"},
	}

	json.NewEncoder(w).Encode(todos)
}

// TodoDetail request handler for "/todos/{todoID}"
func TodoDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID := params["todoID"]
	fmt.Fprintln(w, "Todo detail:", todoID)
}
