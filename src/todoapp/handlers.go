package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index request handler for "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex request handler for "/todos"
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

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

// TodoCreate request handler for create todo on POST "/todos"
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if encError := json.NewEncoder(w).Encode(err); encError != nil {
			panic(encError)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
