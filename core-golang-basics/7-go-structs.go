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

4. Using new:
  p:= new(Person)
  p.Name="John"
*/

/**
####################### Structs with pointers
- When you want to modify the struct.
- When you want to avoid copying large structs.

func updateAge(p *Person, newAge int){
p.Age=newAge
}
##### Pointer vs Valu Receiver::
1. Value Receiver: gets a copy(doesn't modify original)
2. Pointer Receiver: operates on the original instance.

func (p *Person) Birthday(){
p.Age++
}
}
*/

/***
Embedded structs (Composition):
Go has no inheritance, but you can embed one struct into another to reuse fields/methods.
type Address struct{
  City string
}
type Employee struct{
  Name string
  Address
}
# You can access embedded fields directly:
e := Employee{Name: "Joe", Address: Address{City: "NYC"}}
fmt.Println(e.City) // works due to promotion


########## Anonymous structs:
# Useful for temporary structures.
emp := struct{
  ID int
  Name string
}{Id: 101, Name:"Sam"}
fmt.Println(emp.Name)

######### Tags Used for JSON, DB etc...
type User struct {
    ID    int    `json:"id"`
    Email string `json:"email_address"`
}
These tags are used with reflection in encoding, database, and validation libraries.


############# Structs with Interfaces: Structs can implement interfaces implicitly:
type Speaker interface {
    Speak()
}
func (p Person) Speak() {
    fmt.Println("Hi, I'm", p.Name)
}
You don’t need to declare that a struct implements an interface — Go checks compatibility at compile-time.
############## Structs in slices and map::
people := []Person{
    {Name: "Alice", Age: 25},
    {Name: "Bob", Age: 30},
}

m := map[string]Person{
    "a": {"Anna", 22},
}
| Tip                                                                                         | Why                                  |
| ------------------------------------------------------------------------------------------- | ------------------------------------ |
| Use pointer receivers when the method modifies the struct or to avoid copying large structs | Improves performance                 |
| Use field tags when interacting with JSON or databases                                      | Enables encoding/decoding            |
| Prefer named initialization for clarity                                                     | Avoids bugs from positional mistakes |
| Use embedding for code reuse instead of inheritance                                         | Idiomatic Go style                   |


### Common Mistakes
| Mistake                                                          | Fix                                       |
| ---------------------------------------------------------------- | ----------------------------------------- |
| Using value receiver when you need to modify struct              | Use pointer receiver `(*Type)`            |
| Forgetting to initialize nested structs                          | Always initialize before use              |
| Assuming inheritance with struct embedding                       | Embedding is composition, not inheritance |
| Comparing structs with uncomparable fields (like slices or maps) | Avoid direct comparison                   |


### SUmmary
| Feature         | Syntax/Example                       |
| --------------- | ------------------------------------ |
| Define struct   | `type Person struct { Name string }` |
| Initialize      | `p := Person{"Tom", 20}`             |
| Field access    | `p.Name`                             |
| Pointer access  | `(*p).Name` or `p.Name`              |
| Method          | `func (p Person) Greet() {}`         |
| Embedded struct | `type Emp struct { Person }`         |
| Tags            | `json:"name"`                        |
| Anonymous       | `struct { Name string }{...}`        |


*/
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p *Person) Birthday() {
    p.Age++
}

func (p Person) Greet() {
    fmt.Println("Hi, I'm", p.Name)
}

func main() {
    p := Person{Name: "Alice", Age: 29}
    p.Greet()
    p.Birthday()
    fmt.Println("New age:", p.Age)
}

