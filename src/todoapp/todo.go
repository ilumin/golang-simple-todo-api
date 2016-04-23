package main

import "time"

// Todo a structure for todo object
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

// Todos a collection of Todo
type Todos []Todo
