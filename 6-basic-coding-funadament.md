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
