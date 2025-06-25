package main

import (
	"github.com/ankush/error/deferExample"
)

//	func checkName(name string) error {
//		//create a new error:
//		newError := errors.New("Invalid name")
//		if name != "NewName" {
//			return newError
//		}
//		return nil
//	}
func main() {
	name := "Hello"
	defer deferExample.DeferExample(name)
	// err := checkName(name)
}
