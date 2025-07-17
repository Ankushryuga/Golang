/***
Struct Embedding:  
  Embedding means you put a struct directly inside another struct without giving it a field name, so the embedded struct's fields and methods become available on the outer struct
  as if ther were declared there.

Benefits of Embedding:
  - Code reuse without inheritance.
  - Simpler and more explicit composition.
  - Method promotion.
  - Useful for things like mixins, extending structs, or implementing interfaces.

| Feature          | Explanation                                         |
| ---------------- | --------------------------------------------------- |
| Embedding        | Embed a struct by declaring it without a field name |
| Field Promotion  | Embedded struct’s fields accessible directly        |
| Method Promotion | Embedded struct’s methods accessible directly       |
| Override         | Outer struct fields/methods override embedded ones  |
| Composition      | Preferred over inheritance in Go                    |

*/

package main
import "fmt"

type Person struct{
  Name string
  Age int
}

type Employee struct{
  Person        //embedded struct..
  ID int
}



/****
######################## Embedding Methods:   If the embedded struct has methods, those methods are promoted to the embedding struct.
*/

func (p Person) Greet(){
  fmt.Printf("Hello, my name is %s\n", p.Name)
}

/**************
####################### Overriding Fields/Methods:  If the outer struct declares a field or method with the same name as the embedded one, it overrides the embedded version.
*/
type Animal struct{
  Name string
}
type Dog struct{
    Animal
    Name string  /overrides embeded Animal.Name.
}



func main(){
  e:= Employee{
    Person := Person{Name:"Alice", Age:40}
    ID: 100
  }
  fmt.Println(e.Name)
  fmt.Println(e.Age)
  fmt.Println(e.ID)

  e1 := Employee{Person:Person{Name:"Bob"}, ID:1002}
  e1.Greet()  //works! o/p: Hello, my name is Bob

  //overriding...
  a1:=Dog{
      Animal: Animal{Name:"Dog"},
      Name: "NICKWA",
  }
  fmt.Println(a1.Name)  //NICKWA
  fmt.Println(a1.Animal.Name)  //Dog.
}

/***
How It Works:
Employee embeds Person.
You can access Person fields (Name, Age) directly on Employee — no need for e.Person.Name.
If you want, you can still access explicitly: e.Person.Name.
*/


/***
## Embedding Multiple Struct....
*/
type ContactInfo struct{
    Email string
    Phone string
}

type EmployeeTWO struct{
    Person
    ContactInfo
    ID int
}

e := EmployeeTWO{
    Person: Person{Name:"Evenline"},
    ContactInfo: ContactInfo{Email: "eveline@gmail.com", phone:"123-456"},
    ID: 100,
}

//fmt.Println(e.Email)  // eve@example.com
//fmt.Println(e.Phone)  // 123-456
