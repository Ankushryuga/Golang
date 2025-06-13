A package is like a folder that puts your program into order by classifying related functions, types, and files that makes your program neat, reusable, and modular. For instance, the fmt package provides facilities to display messages on the screen and avoids name conflicts.


# Here, we have two significant kinds of packages to learn −

1. Standard Library Packages:These are pre-built for operations such as dealing with I/O, dealing with strings, dealing with concurrency, and more. For instance, fmt package deals with formatted I/O and the net/http package deals with creating HTTP servers.
2. User-defined Packages: Users can define their own packages to encapsulate and reuse their code in different projects.


# Blank Import:
import _ "strings"



# Access Modifiers:
    => 
      The visibility of identifiers (functions, variables, constants, etc.) is determined by their naming convention:
      
      An identifier with a name that begins with an uppercase letter is exported and can be accessed outside the package, 
      and an identifier with a lowercase name is unexported and can only be accessed within the package.

# Example:
    =>
    package main
    import "fmt"
    // Exported function
    func HelloWorld() {
        fmt.Println("Hello, World!")
    }
    // Unexported function
    func helloWorld() {
        fmt.Println("hello, world!")
    }
    func main() {
       HelloWorld()  // Calling the exported function
       helloWorld()  // Calling the unexported function
    }


# Alising Imports:
    =>When importing packages, you can give them an alias to simplify usage in your code

# Example:
      import io "fmt"
      func main() {
          io.Println("Using alias for fmt package")
      } 

# Initializing Packages:
    => You can implement an init() function inside a package, which gets executed when the package is imported. This helps in initializing configurations or variables.

    =>
      package example
      import "fmt"
      func init() {
          fmt.Println("Package initialized")
      }

# Third-Party Packages:
    => You can use third-party packages from external sources by integrating them with the Go Modules system.

# First, initialize your project:
    => go mod init your_project_name
# Then, add a dependency:
    => go get github.com/example/package
#  Finally, import and use the package in your code:
    => import "github.com/example/package"


# Package Documentation:
    => The go doc tool can help you generate and view documentation for your package.
      go doc fmt


# Example:
    => 
    /myproject
      main.go
      greetings
         hello.go
        //greetings/hello.go
         package greetings
        import "fmt"
        func Hello() {
            fmt.Println("Hello from greetings package!")
        }
        //main.go
        package main
        import (
            "myproject/greetings"
        )
        func main() {
            greetings.Hello()
        }
