//use case 1: Handling interface{} Parameters (e.g., JSON, gRPC, Reflection)

package main

import "fmt"

func handleInterfaceParameter(v interface{}){
  switch val := v.(type){
    case string:
      fmt.Println("It's a string: ", val)
    case int:
      fmt.Println("It's a int: ", val)
    case bool:
      fmt.Println("It's a bool: ", val)
    default:
      fmt.Println("unknownn type: ")
  }
}
