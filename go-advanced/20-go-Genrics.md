# Genrics:
      =>
      Go can usually infer type parameters, so you don't have to declare them. The Generics are implemented at compile time, so they don't have runtime overhead like reflection. But overuse of generics can cause code bloat because of monomorphization (each type instantiation creates different code).

# Type Parameters:
      =>
      Type parameters in Go allow you to define functions, structs, or interfaces that can operate on multiple types, specified using square brackets []

      func Print[T any](value T) { ... }

# Example:
      =>
      package main
      import "fmt"
      
      type Stringer interface {
         String() string
      }
      func PrintString[T Stringer](value T) {
         fmt.Println(value.String())
      }
      type MyType struct {
         data string
      }
      func (m MyType) String() string {
         return m.data
      }
      func main() {
         val := MyType{data: "Hello, Generics!"}
         PrintString(val) 
      }



# Constraints
      => The Constraints is used to specify what types are allowed for a type parameter. Here, we use predefined constraints like any (any type) or comparable (types that support == and !=) and also we can define custom constraints using interfaces.

# Example:
      => package main
        import (
           "fmt"
           "golang.org/x/exp/constraints"
        )
        // Add numbers of any numeric type
        func Add[T constraints.Integer | constraints.Float](a, b T) T {
           return a + b
        }
        func main() {
           fmt.Println(Add(1, 2))
           fmt.Println(Add(1.5, 2.5))
        }


# Generic Functions:
      => The Generic Functions are the functions that can operate on multiple types.
      =>
      package main
      import "fmt"
      
      // Print any type
      func Print[T any](value T) {
         fmt.Println(value)
      }
      func main() {
         Print(42)       
         Print("Hello")   
         Print(3.14)      
      }

# Generic Types:
    =>The Structs, slices, or other types that can work with type parameters is called as generic types.

# Example:
    =>
    package main
    import "fmt"
    
    type Stack[T any] struct {
       items []T
    }
    func (s *Stack[T]) Push(item T) {
       s.items = append(s.items, item)
    }
    func (s *Stack[T]) Pop() T {
       if len(s.items) == 0 {
          panic("stack is empty")
       }
       item := s.items[len(s.items)-1]
       s.items = s.items[:len(s.items)-1]
       return item
    }
    func main() {
       intStack := Stack[int]{}
       intStack.Push(1)
       intStack.Push(2)
       fmt.Println(intStack.Pop()) 
    
       stringStack := Stack[string]{}
       stringStack.Push("Go")
       stringStack.Push("Generics")
       fmt.Println(stringStack.Pop())
    }
