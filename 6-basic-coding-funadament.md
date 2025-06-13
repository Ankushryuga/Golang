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
