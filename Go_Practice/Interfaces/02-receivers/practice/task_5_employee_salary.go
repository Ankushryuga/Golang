package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (em *Employee) IncreaseSalary(percent float64) {
	em.Salary = em.Salary * (1 + percent/100)
}

func (em Employee) PrintSalary() {
	fmt.Println("Salary", em.Salary)
}
