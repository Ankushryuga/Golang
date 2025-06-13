# Inheritance:
    =>
    Golang does not support direct inheritance like some other languages. You can achieve it by using he "embedding" that means you can embed one struct inside another struct. 
    This allows the outer struct to use the fields and methods of the inner struct.

# Composition (Embedding)
    => The Composition (Embedding) is used to combine structs by embedding one struct inside another(which is similar to inheritance).

# Example:
    =>
    package main
    import "fmt"
    type Animal struct {
       Name string
    }
    func (a *Animal) Speak() {
       fmt.Println(a.Name, "makes a sound")
    }
    type Dog struct {
       Animal 
       Breed  string
    }
    func main() {
       d := Dog{
          Animal: Animal{Name: "Rex"},
          Breed:  "German Shepherd",
       }
       d.Speak()
    }


  # Interface:
      =>Interface in Golang is a way to define a set of methods that a type must implement. This allows you to achieve polymorphism, which is another aspect of inheritance.

# Example:
      =>
      package main
      import "fmt"
      type Speaker interface {
         Speak()
      }
      type Animal struct {
         Name string
      }
      func (a *Animal) Speak() {
         fmt.Println(a.Name, "makes a sound")
      }
      type Dog struct {
         Animal 
         Breed  string
      }
      
      func main() {
         var s Speaker
         d := Dog{
            Animal: Animal{Name: "Rex"},
            Breed:  "German Shepherd",
         }
         s = &d
         s.Speak()
      }

# Method Overriding:
      => Go does not support method overriding directly, but you can achieve similar behavior by defining methods with the same name on the derived struct.

# Example:
      => 
      package main
      import "fmt"
      type Animal struct {
         Name string
      }
      func (a *Animal) Speak() {
         fmt.Println(a.Name, "makes a sound")
      }
      type Dog struct {
         Animal
         Breed  string
      }
      func (d *Dog) Speak() {
         fmt.Println(d.Name, "barks")
      }
      func main() {
         d := Dog{
            Animal: Animal{Name: "Rex"},
            Breed:  "German Shepherd",
         }
         d.Speak()
      }


# Type Assertion and Type Switching:
      =>If you need to access specific methods or fields to a derived type, you can use type assertion or type switching.

# Example
      =>
      package main
      import "fmt"
      type Speaker interface {
         Speak()
      }
      type Animal struct {
         Name string
      }
      func (a *Animal) Speak() {
         fmt.Println(a.Name, "makes a sound")
      }
      type Dog struct {
         Animal
         Breed string
      }
      func (d *Dog) Speak() {
         fmt.Println(d.Name, "barks")
      }
      func main() {
         var s Speaker
         d := Dog{
            Animal: Animal{Name: "Rex"},
            Breed:  "German Shepherd",
         }
         s = &d
         // Type assertion
         if dog, ok := s.(*Dog); ok {
            fmt.Println("Breed:", dog.Breed)
         }
         // Type switch
         switch v := s.(type) {
         case *Dog:
            fmt.Println("This is a dog of breed:", v.Breed)
         case *Animal:
            fmt.Println("This is an animal named:", v.Name)
         }
      }

  
      
