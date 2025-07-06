/***
Go supports special types of functions called methods.
In method declaration syntax, a "receiver" is present to represent the container of the function.
The receiver can be used to call a function using "." operator.


In sort:
Methods are functions with a receiver, allowing you to associate behaviour with a type.

//syntax:
func (variable_name variable_data_type) function_name() [return_type]{}

## NOTE: 
methods aer scoped to package, not instance.

GO don't allow method name overloading...
func (b Box) Print() {}
func (b Box) Print(v int) {} // ❌ Compile error
*/

package main

import "fmt"

type User struct{
  Name string
}
//value receiver::
func (u User) setNamea(name string){
  u.Name=name      //this won't change the original..
}

func (u *User) setName(name string){
  u.Name=name  //changes the original
}


//Method with Nil Receiver:::
type Logger struct{
  Prefix string
}

func (l *Logger) Log(msg string){
  if l==nil{
      fmt.Println("nil logger: ", msg)
      return
  }
  fmt.Println(l.Prefix, msg)
}

//end
func main(){
  u := User{Name:"Ankush"}
  u.setName("Raj")
  fmt.Println(u.Name)


  var l *Logger=nil
  l.Log("test")
}
