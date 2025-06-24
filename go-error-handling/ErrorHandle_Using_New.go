package main

import (
	"errors"
	"fmt"
)

func checkName(name string) error {
	//create a new error:
	newError := errors.New("Invalid name")
	if name != "NewName" {
		return newError
	}
	return nil
}
func main() {
	name := "Hello"
	err := checkName(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("valid name")
	}
}
