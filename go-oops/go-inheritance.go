/**
Go supports inheritance through composition:
it is used to combine structs by embedding one struct inside another
*/
package main

import "fmt"

type Animal struct{
  Name string
}

func (a *Animal) Speak(){
  fmt.Println(a.Name, "makes a sound")
}

type Dog struct{
  Animal,
  Breed string
} 

func main(){
  d := Dog{
    Animal: Animal{Name:"Rex"},
    Breed: "German shepherd",
    }
  d.Speak()
}
