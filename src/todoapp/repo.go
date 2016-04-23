package main

import "fmt"

var currentID int

var todos Todos

func init() {
	// RepoCreateTodo(Todo{Name: "Learn GO!"})
	// RepoCreateTodo(Todo{Name: "Learn write API with GO!"})
}

// RepoCreateTodo create todo item
func RepoCreateTodo(todo Todo) Todo {
	currentID++
	todo.ID = currentID
	todos = append(todos, todo)
	return todo
}

// RepoFindTodo find todo item
func RepoFindTodo(id int) Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}

	return Todo{}
}

// RepoDestroyTodo remove todo item
func RepoDestroyTodo(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Cound not find Todo with id of %d to delete", id)
}
