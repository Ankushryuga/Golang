/**
Interface: Its a type that specifies method signatures, a type satisfies an interface if it implements all the methods declared in that interface.

## syntax:
type Animal interface{
  Speak() string
}
Any type that has a method with the exact same signature will implement Animal implicitly - there's no implement keyword in golang.

type Dog struct{}
func (d Dog) Speak()string{
return "woof"
}
var a Animal = Dog{}// Dog satisfies Animal..

## Interface Internals(How they work)
# Under the hood, an interface value is a pair.
    - A type (called itab - interface table)
    - A value (concrete value implementing that type).

## so this:
var a Animal = Dog{}
//is internally like:
type: *Dog
value : Dog{}
*/
package main
import "fmt"

//Type assertion and type switch..: use it when you want to extract the concrete type from an interface:
var a Animal = Dog{}
d := a.(Dog)  //asserts that a is a Dog.
//Use the comma-ok idiom to avoid panics.
d, ok := a.(Dog)
if ok{
  fmt.Println("its a dog")
}


///Type switch:
switch v := a.(type) {
  case Dog: 
    fmt.Println("Dog says:", v.Speak())
  case Cat:
    fmt.Println("Cat says:", v.Speak())
  default:
    fmt.Println("Unknown animal")
}

//Empty interface::interface{}
//The empty interface can represent any value b/c every type satisfies it(has zero method requirements)
var x interface{}=42
var y interface{}="hello"
var z interface{}=Dog{}


//Interface Nil vs Nil Interface value..
var d *Dog = nil
var a Animal = d
fmt.Println(a==nil)  //false

//E
