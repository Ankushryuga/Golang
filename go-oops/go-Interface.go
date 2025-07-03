/**
Interface in go, is a way to define a set of methods that a type must implement. This allows you to achieve polymorphism,

which is another aspect of inheritance.
*/
package main

import "fmt"

type Speaker interface{
  Speak()
}

type Animal struct{
  Name string
}

func (a *Animal) Speak(){
  fmt.Println(a.Name, "makes a sound")
}

type Dog struct{
  Animal
  Breed string
}

func main(){
  var s Speaker
  d := Dog{
    Animal: Animal{Name:"Rex"},
    Breed: "German shepherd",
  }

  s=&d
  s.Speak()
}
