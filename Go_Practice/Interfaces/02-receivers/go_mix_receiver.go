package main

import "fmt"

// Example of mix receiver in Go
type Student struct {
	Name  string
	Marks int
}

// value receiver
func (s Student) PrintResult() {
	fmt.Println(s.Name, "got", s.Marks, "marks")
}

// pointer receiver
func (s *Student) AddGraceMarks(marks int) {
	s.Marks += marks
}
