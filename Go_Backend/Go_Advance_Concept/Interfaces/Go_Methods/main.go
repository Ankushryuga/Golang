// Go Methods examples
/**
Go methods are functions that is associated to a type.
**/

package main

import "fmt"

type Employee struct{
  ID    int
  Name  string
  Desg  string
}

// Get employee information and error info
func (e *Employee) GetEmployee(ID  int) Employee, error {
}
