# Keywords
# loops  
    =>
       var b int = 15
   var a int
   numbers := [6]int{1, 2, 3, 5} 

   /* for loop execution */
   for a := 0; a < 10; a++ {
      fmt.Printf("value of a: %d\n", a)
   }
   for a < b {
      a++
      fmt.Printf("value of a: %d\n", a)
   }
   for i,x:= range numbers {
      fmt.Printf("value of x = %d at %d\n", x,i)
   } 

    str := "TutorialsPoint"

    // Loop through the string by index
    fmt.Println("Characters of the string with their index:")
    for i := 0; i < len(str); i++ {
        fmt.Printf("[%2d]: %c\n", i, str[i])
    }



    # Infinite loop 
      for {
      fmt.Println()
      }
# case
# switch case 

# Statements:
        => Break, continue, goto:
        Goto: it provides an unconditional jump from the goto to a labeled statement in the same         function.
## Example:
        package main

        import "fmt"
        
        func main() {
           /* local variable definition */
           var a int = 10
        
           /* do loop execution */
           LOOP: for a < 20 {
              if a == 15 {
                 /* skip the iteration */
                 a = a + 1
                 goto LOOP
              }
              fmt.Printf("value of a: %d\n", a)
              a++     
           }  
        }

         
etc.



## Functions:
    => 
            A function declaration tells the compiler about a function name, return type, and parameters.
            A function definition provides the actual body of the function.

# Functions are also known as method, sub-routine, or procedure.


# Defining a Function in Go Language
        => A function is defined in Go language by using the func keyword followed by the function_name, parameter list, and the return type.

# Syntax:
        =>
        func function_name( [parameter list] ) [return_types]
        {
           body of the function
        }

# Components of a Function
1. Func
2. Function Name
3. Parameters
4. Return type
5. Function Body



# Example:        
        =>
            package main
            import "fmt"
            func swap(x, y string) (string, string) {
               return y, x
            }
            func main() {
               a, b := swap("Mahesh", "Kumar")
               fmt.Println(a, b)
            }


## Function Usage:
        => A function can be used in the following ways:
            1. Function as value: can be created on the fly and can be used as values.
            2 Function Closures: are anonymous functions and can be used in dynamic programming
            3. Method: are special function with a receiver.


# Function as value:
        =>
        create a functions on the fly and use them as values, 

        Example:
        func main(){
            /* declare a function variable.*/
            getSquareRoot := func(x float64) float64 {
                return math.Sqrt(x)
            }
            //use the function.
            fmt.Println(getSquareRoot(9))
        }


# Assigning a function to a variables:
        =>
        func addTwoNumbers(a int, b int) int {
            return a+b
        }
        func main(){
            //assign the function to a variable.
            sum := addTwoNumbers
            result := sum(100, 200)
            fmt.Println("sum: ", result);
        }

# Passing a function as an argument:
        =>
        func calculation (x int, y int, op func(int, int) int) int {
            return op(x, y)
        }

        func multiplyNumbers(x int, y int) int {
            return x*y
        }

        func main(){
            result := calculation(2, 4, multiplyNumbers)
            fmt.Println("Result", result)
        }


## Returning a funciton as a value:
        =>
        A function can also return a function as a value in Go, it is useful when you want to return an expression calculated by creating any function.


# example:
        => 
        func calculate(facotr int) func (int) int {
            return func (value int) int {
                return factor * value
            } 
        } 
        func main(){
            multiply := calculate(2)
            result := multiply(20)
            fmt.Println("result: ", result)
        }
        


## Function Closure:
    =>
    Function closure: it's a function value that references variable from outside its body, the closures bind these variable and make them accessible even after closing the outer function.

# Creating closure:
    A closure is created when an inner function captures and retains access to the variables of its enclosing function. You can create a closure by following:
        => 
        counter := func() func() int {
            count := 0
            return func() int {
                count++
                return count
            }
        }

        //use:
        increment := counter()

        => counter function defines a local variable count, and the returned inner function, which will be known as closure, has access to count and modifies it on each call.

        =>
            package main
            import "fmt"
            
            func main() {
                // First, define an outer function
                updateCounter := func() func() int {
                    // define a local variable inside the function
                    count := 100
                    return func() int {
                        count++
                        return count
                    }
                }
            
                // Now, creating a closure
                increment := updateCounter()
            
                // Using (calling) the closure
                fmt.Println(increment())
                fmt.Println(increment())
            }

# Passing values into closures:
    =>
        package main
        import "fmt"
        
        func main() {
            // Define an outer function that accepts a value
            updateCounter := func(initial int) func() int {
                count := initial // Initialize count with the passed value
                return func() int {
                    count++
                    return count
                }
            }
        
            // Create a closure with an initial value
            inc_x := updateCounter(100)
        
            // Use the closure
            fmt.Println(inc_x())
            fmt.Println(inc_x())
        
            // Create another closure with an initial value
            inc_y := updateCounter(200)
        
            // Use the closure
            fmt.Println(inc_y())
            fmt.Println(inc_y())    
        }



## Method:
        =>
        Go programming language supports special types of functions called methods. In method declaration syntax, a "receiver" is present to represent the container of the function. This receiver can be used to call a function using "." operator. For example

# Syntax:
    =>
    func (variable_name variable_data_type) function_name() [return_type]{
       /* function body*/
    }

# Example:
    =>
    package main
    import (
       "fmt" 
       "math" 
    )
    
    /* define a circle */
    type Circle struct {
       x,y,radius float64
    }
    
    /* define a method for circle */
    func(circle Circle) area() float64 {
       return math.Pi * circle.radius * circle.radius
    }
    
    func main(){
       circle := Circle{x:0, y:0, radius:5}
       fmt.Printf("Circle area: %f", circle.area())
    }

# Methods with Struct Type Receiver
        => You can create a method with a struct type receiver that operates on a copy of the struct, so whatever changes will be done in the method, they do not affect the original struct.

# Example:
        =>
        package main
        import "fmt"
        
        // Struct
        type Rectangle struct {
            width, height float64
        }
        
        // Define a method  with struct type receiver
        func (rect Rectangle) Area() float64 {
            return rect.width * rect.height
        }
        
        func main() {
            rectObj := Rectangle{width: 2.4, height: 4.5}
            fmt.Println("Area of Rectangle:", rectObj.Area())
        }

# Methods with Non-Struct Type Receiver
        => You can also create a method withnon-struct type receivers of such type whose definition is present in the same package.

        => package main
            import "fmt"
            
            // Creating a custom type based on int
            type value int
            
            // Defining a method with a non-struct receiver
            func (v value) cube() value {
                return v * v * v
            }
            
            func main() {
                x := value(3)
                y := x.cube()
            
                fmt.Println("Cube of", x, "is", y) 
            }





## Methods with Pointer Receiver:
        =>You can create a method that can havepointer receivers. This approach reflects the changes done in the method in the caller.

# Example:
        =>
        package main
        import "fmt"
        
        type student struct {
            grade string
        }
        
        // Method with pointer receiver to modify data
        func (s *student) updateGrade(newGrade string) {
            s.grade = newGrade
        }
        
        func main() {
            s := student{grade: "B"}
            
            fmt.Println("Before:", s.grade)
            
            // Calling the method to update the grade
            s.updateGrade("A")
            
            fmt.Println("After:", s.grade)
        }



## Anonymous or Lambda Function or Function Literals: are function without name:
        =>
        1. Executing code on the spot
        2. Satisfying variable usage from surrounding code
        3. Scheduling fast jobs in goroutines
        4. Sweeping up with defer statements
        
        func(parameters) return_type {
            // function body
        }


# Example:
            =>
            package main
            import "fmt"
            func main() {
               greet := func() {
                  fmt.Println("Tutorialspoint")
               }
               greet() // Call the anonymous function
            
               // With parameters and return value
               add := func(a, b int) int {
                  return a + b
               }
               sum := add(9, 5)
               fmt.Println("Sum =", sum)
            }


# Invoked Function expression:
        =>
        package main
        import "fmt"
        func main() {
           // Without parameters
           func() {
              fmt.Println("Executed immediately!")
           }()
           
           // With parameters and return value
           result := func(a, b int) int {
              return a * b
           }(6, 4)
           fmt.Println("Result:", result)
           
           // With multiple statements
           func(name string) {
              fmt.Printf("Hello, %s!\n", name)
              fmt.Println("Welcome to Go!")
           }("Revathi")
        }

# Closures are variable capturing:
        => Anonymous functions can access variables from their enclosing scope, creating closures:

# Example:
            => package main
                import "fmt"
                func main() {
                   c := counter()
                   fmt.Println(c()) // 1
                   fmt.Println(c()) // 2
                   fmt.Println(c()) // 3
                   
                   double := multiplier(2)
                   triple := multiplier(3)
                   fmt.Println(double(6)) 
                   fmt.Println(triple(6)) 
                }  
                func counter() func() int {
                   count := 0
                   return func() int {
                      count++
                      return count
                   }
                }
                func multiplier(factor int) func(int) int {
                   return func(x int) int {
                      return x * factor
                   }
                }


# Passing Anonymous Functions as Arguments
        =>
        package main
        import "fmt"
        func main() {
           squares := transform([]int{2, 4, 6}, func(x int) int {
              return x * x
           })
           fmt.Println(squares) 
        }
        
        func transform(numbers []int, op func(int) int) []int {
           result := make([]int, len(numbers))
           for i, v := range numbers {
              result[i] = op(v)
           }
           return result
        }


# As Callbacks:
        =>
        package main
        import "fmt"
        func main() {
           squares := transform([]int{2, 4, 6}, func(x int) int {
              return x * x
           })
           fmt.Println(squares) 
        }
        
        func transform(numbers []int, op func(int) int) []int {
           result := make([]int, len(numbers))
           for i, v := range numbers {
              result[i] = op(v)
           }
           return result
        }

# Recursive Anonymous functions:
        => package main
            import "fmt"
            func main() {
               var factorial func(int) int
               factorial = func(n int) int {
                  if n <= 1 {
                     return 1
                  }
                  return n * factorial(n-1)
               }
               fmt.Println(factorial(5))
            }


# Struct Fields with Function Types:
        =>
        Structs in Go can hold functions as fields, allowing each instance to have different behavior.
        the processor struct has two function fields - transform for converting strings and validate for validations. You can pass various anonymous functions to these fields while creating instances so that each processor will behave differently but have the same interface.

# Example:
        =>
        package main
        import (
           "fmt"
           "strings"
        )
        func main() {
           p := Processor{
              transform: func(s string) string { return strings.ToUpper(s) },
              validate:  func(s string) bool { return len(s) > 0 },
           }
           fmt.Println(p.transform("hello")) 
           fmt.Println(p.validate(""))      
        }
        type Processor struct {
           transform func(string) string
           validate  func(string) bool
        }



## Strings:

    package main

import "fmt"

func main() {
   var greeting =  "Hello world!"
   
   fmt.Printf("normal string: ")
   fmt.Printf("%s", greeting)
   fmt.Printf("\n")
   fmt.Printf("hex bytes: ")
   
   for i := 0; i < len(greeting); i++ {
       fmt.Printf("%x ", greeting[i])
   }
   
   fmt.Printf("\n")
   const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
   
   /*q flag escapes unprintable characters, with + flag it escapses non-ascii 
   characters as well to make output unambigous */
   fmt.Printf("quoted string: ")
   fmt.Printf("%+q", sampleText)
   fmt.Printf("\n")  
}

        
 
