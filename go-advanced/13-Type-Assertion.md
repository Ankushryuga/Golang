# type assertion :
    => its a way to retrieve the dynamic value of an interface. It is used to extract the underlying value of an interface and convert it to a specific type.
     This is useful when you have an interface with a value of an unknown type and you need to operate on that value.

## Syntax:
    =>
    t := Value.(TypeName)

# Basic Type Assertion
    => The basic type assertion is a way to inform the compiler about the specific type of a variable when it can't automatically determine it that helps to avoid runtime errors.
# Example:
    =>
      package main
      import (
          "fmt"
      )
      
      func main() {
          var i interface{} = "hello"
          
          // Type assertion to retrieve the string value
          s := i.(string)
          fmt.Println(s)
      }

# Type Assertion with Ok-idiom
    => The Ok-idiom is a common pattern for handling errors, and it's often used with type assertions to ensure that operations succeed.

# Example:
    =>
    package main
    import (
        "fmt"
    )
    
    func main() {
        var i interface{} = "hello"
        
        // Type assertion with Ok-idiom
        if s, ok := i.(string); ok {
            fmt.Println("String value:", s)
        } else {
            fmt.Println("Not a string")
        }
    }

# Handling Failed Type Assertion
      =>
      package main
      import (
          "fmt"
      )
      
      func main() {
          var i interface{} = 42
      
          // Attempting to assert to a wrong type
          if s, ok := i.(string); ok {
              fmt.Println(s)
          } else {
              fmt.Println("Type assertion failed")
          }
      }
