package main

import "fmt"

// Problem 6: Value Receiver Method is Available to Both Value and Pointer

type Value_Receiver struct {
	Name string
}

func (u Value_Receiver) Print() {
	fmt.Println(u.Name)
}

/**
NOTE: If a method has a value receiver, both "Value_Receiver" and "*Value_Receiver" can call it.
*/
