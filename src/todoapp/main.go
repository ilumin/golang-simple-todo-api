package main

import (
	"log"
	"net/http"
)

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoDetail", "GET", "/todos/{todoID}", TodoDetail},
}

func main() {
	router := NewRouter(routes)
	log.Fatal(http.ListenAndServe(":8080", router))
}
