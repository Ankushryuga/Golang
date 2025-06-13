# A pointer :
    =>
    is a powerful feature in the Golang programming language that allows you to work with the memory addresses.

In Golang, Dereferencing a pointer means accessing the value stored at the memory address pointer is pointing to. This concept is important in Golang as it enables you to optimize performance, share data between functions without copying, and manipulate data.


To declare a pointer, we use the * operator, and to get the memory address of a variable, we can use &(AND) operator.


# Example:
    =>
      package main
      import "fmt"
      
      func main() {
         var x int = 64   
         var p *int        
         p = &x          
      
         fmt.Println("Value of x=", x)
         fmt.Println("Address of x=", &x)
         fmt.Println("Value of p=", p)
         fmt.Println("Dereferenced value of p=", *p)
         *p = 21 
         fmt.Println("New value of x=", x)
      }


package main
import "fmt"

type Student struct {
   name string
   age  int
}

func main() {
   p := &Student{"Indu", 18}
   fmt.Println("Name=", p.name)
   fmt.Println("Age=", (*p).age)
}
