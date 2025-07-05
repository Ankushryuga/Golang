/**
type assetion: its a way to retrieve the dynamic value of an interface It is used to extract the underlying value of an interface and convert it to a specific type. 
This is useful when you have an interface with a value of an unknown type and you need to operate on that value.
*/

/***
Internals: What happens under the hood?
Every interface{} in Go contains.
1. A type pointer
2. A data pointer

when you assert iface.(T), Go checks whether iface.type==T, If yes, it gives back the data as type T 
*/


package main

import "fmt"

func main(){
    var i interface{}="hello"
    s := i.(string)


    //safe type assertion: (comma-ok idiom)
    var in interface{} = "Hi"
    s, ok := in.(string)
    fmt.Println(s, ok)        //output: Hi, true

    n, ok := in.(int)
    fmt.Println(n, ok)    //output: 0, false
    
    fmt.Println(s)
  }
