package main

import "fmt"

type Todo struct {
	Title     string
	Completed bool
}

func (todo *Todo) MarkCompleted() {
	todo.Completed = true
}

func (todo *Todo) MarkIncomplete() {
	todo.Completed = false
}

func (todo Todo) Print() {
	status := "Incomplete"
	if todo.Completed {
		status = "Completed"
	}
	fmt.Printf("%s - %s\n", todo.Title, status)
}
