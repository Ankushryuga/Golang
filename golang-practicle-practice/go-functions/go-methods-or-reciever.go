/***
Methods are like functions but with a key difference: they have a receiver argument, which allows access to the receiver's properties.
The receiver can be a struct or non struct type, but both must be in the same package. 
Methods cannot be created for types defined in other pacakges, including built-in types like int or string;
otherwise compiler will raise an error.
*/


package main
import "fmt"

type person struct{
  name string 
  age int
}

//defining a method with struct receiver.
func (p person) display(){
  fmt.Println("Name: ", p.name)
  fmt.Println("Age: ", p.age)
}

func main(){
  a := person{name:"ankush", age:26}

  a.display()
}
