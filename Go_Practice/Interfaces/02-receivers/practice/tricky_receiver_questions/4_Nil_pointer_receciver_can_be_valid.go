package main

import "fmt"

type User_Nil_Pointer_Receiver struct {
	Name string
}

func (u *User_Nil_Pointer_Receiver) PrintName() {
	if u == nil {
		fmt.Println("user is nil")
		return
	}
	fmt.Println(u.Name)
}

/**
A method with pointer receiver can be called on a nil pointer, as long as the method handles nil safely
But this would avoid panic:
func (u * User_Nil_Pointer_Receiver) PrintName(){
	fmt.Println(u.Name)
}
*/
