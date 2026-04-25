package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) PrintDetails() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}

func main()
