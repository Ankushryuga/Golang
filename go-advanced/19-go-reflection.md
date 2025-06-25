# Reflection:  
      =>
      Reflection in Go is an important feature that enables you to look inside and tweak the types and values of variables during runtime. It is supported by the reflect package, which belongs to the Go standard library. It is useful in dynamic scenarios such as serialization, deserialization, and generic programming.

# Serialization/Deserialization:
      => Reflection allows libraries such as encoding/json and encoding/xml to dynamic encoding and decoding data.
      => Generic Programming: Prior to Go 1.18, reflection was an important instrument for generic-like behavior implementation.

      https://www.tutorialspoint.com/go/go_reflection.htm

# Example:
package main
import (
   "fmt"
   "reflect"
)
func main() {
   var x float64 = 3.14
   t := reflect.TypeOf(x)
   fmt.Println("Type:", t) 
   v := reflect.ValueOf(x)
   fmt.Println("Value:", v)
   fmt.Println("Kind:", v.Kind())
   if v.Kind() == reflect.Float64 {
      fmt.Println("x is a float64")
   }
   p := reflect.ValueOf(&x)
   e := p.Elem() 
   if e.CanSet() {
      e.SetFloat(2.71) 
      fmt.Println("Modified value of x:", x) 
   }
}


# 
package main
import (
   "fmt"
   "reflect"
)

func main() {
   s := []int{1, 2, 3}
   v := reflect.ValueOf(s)
   for i := 0; i < v.Len(); i++ {
      fmt.Println(v.Index(i).Interface())
   }
}