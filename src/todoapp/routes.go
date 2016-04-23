package main

import "net/http"

// Route a struct for Route object
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes a collection for Route object
type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoDetail", "GET", "/todos/{todoID}", TodoDetail},
}
