/**
type assetion: its a way to retrieve the dynamic value of an interface It is used to extract the underlying value of an interface and convert it to a specific type. 
This is useful when you have an interface with a value of an unknown type and you need to operate on that value.
*/


package main

import "fmt"

func main(){
    var i interface{}="hello"

    s := i.(string)
    fmt.Println(s)
  }
