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


##### JSON tags in structs
| Tag                     | Behavior                                     |
| ----------------------- | -------------------------------------------- |
| `json:"name"`           | Maps JSON field `"name"` to struct field     |
| `json:"name,omitempty"` | Omits the field if it's empty (`zero` value) |
| `json:"-"`              | Ignores the field during (un)marshalling     |


type Book struct {
    Title     string  `json:"title"`
    Author    string  `json:"author,omitempty"`
    Internal  string  `json:"-"` // never included in JSON
}
 Handling Optional Fields with *type
If a field is optional (nullable), use pointers:

type UpdateUserRequest struct {
    Name  *string `json:"name,omitempty"`
    Email *string `json:"email,omitempty"`
}
This way you can distinguish between "zero value" and "not provided."



🚫 Handling Unknown Fields
By default, Go ignores unknown fields in JSON. To prevent this (e.g., for strict APIs), decode into map[string]interface{} or use custom unmarshalling logic.
🛠️ Full Example: API Handler with JSON

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Validate login, generate token, etc.
    res := LoginResponse{
        Token: "secret-token-123",
        User: User{
            ID:       1,
            Username: "john_doe",
            Email:    req.Email,
        },
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}
| Feature                       | Benefit / Usage                     |
| ----------------------------- | ----------------------------------- |
| Structs model real-world data | Useful for DB, APIs, services, etc. |
| JSON tags                     | Control encoding/decoding           |
| Embedded structs              | Composition and reuse               |
| Pointers in struct            | Represent optional/nullable fields  |
| Marshal / Unmarshal           | Convert between JSON and Go types   |


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



//Part 2: How Structs Work with JSON APIs in Go’s encoding/json package provides powerful functionality to encode (marshal) and decode (unmarshal) JSON using structs.

/**
package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    // JSON to struct (Unmarshal)
    jsonStr := `{"id":101, "name":"Laptop", "price":999.99}`
    var p Product
    err := json.Unmarshal([]byte(jsonStr), &p)
    if err != nil {
        panic(err)
    }
    fmt.Println("Struct:", p)

    // Struct to JSON (Marshal)
    out, _ := json.MarshalIndent(p, "", "  ")
    fmt.Println("JSON:\n", string(out))
}

*/

