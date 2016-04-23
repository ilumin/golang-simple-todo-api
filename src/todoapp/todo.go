package main

import "time"

// Todo a structure for todo object
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Todos a collection of Todo
type Todos []Todo
