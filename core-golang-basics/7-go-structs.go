/**
A struct is a composite data type in Go that groups together zero or more fields under a single name.
## //struct declarations
type StructName struct{
  FieldName FieldType
  Another AnotherType
}
//Example:
type Person struct{
  Name string
  Age int
}

##  Declaring and Initializing structs.
1. Named Initialization (recommended for clarity)
    => p := Person{Name: "Alice", Age: 30}

2. Positional Initialization(order matters):
    => p := Person{"Bob", 30}

3. Zero Value: 
    => var p Person
    //p.Name = "", p.Age=0.
*/
package main
import "fmt"
