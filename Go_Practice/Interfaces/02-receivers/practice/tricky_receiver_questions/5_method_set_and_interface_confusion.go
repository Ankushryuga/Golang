package main

import "fmt"

type Printer interface {
	Print()
}

type User_method_set_interface_example struct {
	Name string
}

func (u *User_method_set_interface_example) Print() {
	fmt.Println(u.Name)
}
