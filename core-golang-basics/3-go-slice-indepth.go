/**
- In Go slices are designed to hold elements of a single, specific type. This means you cannot directly create a slice that contains elements of different types.
# However, there are workarounds to achieve similar functionality.

- 1. Using interface{}
- 2. Using struct{}
- 3. Map
*/
/***
- Using interface{} sacrifices type safety, so you need to use type assertions or type switch to work with the values.
- structs and map provide more structure but may require additional effor to manage.
*/

package main
import "fmt"

//1-Using interfaces
func interfaceSlice(){
  mixedSlice := []interface{}{42, "hello", 3.14, true}
  for _, v:= range mixedSlice{
    fmt.Printf("Value: %v, Type: %T\n", v, v)
  }
}

//2 Using struct
type Mixed struct{
  IntVal      int
  StringVal   string
  FloatVal    float64
}

func structSlice(){
  mixedSlice := []Mixed{
    {IntVal: 42, StringVal: "Hello", FloatVal: 3.14}
    {IntVal: 43, StringVal: "World", FloatVal: 3.14223}
  }

  for _, v:= range mixedSlice{
    fmt.Println(v)
  }
}


//3 Using Map:
func mapSlice(){
  mixedSlice := []map[string]interface{}{
    {"id":1, "name":"Ankush", "Score":11.3},
    {"id":2, "name":"Banti", "Score":121.3},
    {"id":3, "name":"AN", "Score":11.3}
  }
  for _, v:=range mixedSlice{
    fmt.Println(v)
  }
}



