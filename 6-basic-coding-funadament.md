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
