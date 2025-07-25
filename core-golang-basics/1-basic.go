/**
This program demonstrates:
- Variable declarations
- Constants
- Different types of for loops
- If/else condition
- Switch statement
*/

package main

import (
  "fmt"
  "time"
  )

// Demonstrates various ways to declare and use variables
func variableMethod() {
  // Declare a string variable with 'var'
  var a = "initial"

  // Short variable declaration using :=
  b := 5

  // Declare and initialize multiple variables of same type
  var e, f int = 1, 2

  // Boolean variable
  var g = true

  // Default value (zero) of uninitialized int variable
  var h int

  // Print variables
  fmt.Println("a =", a)
  fmt.Println("b =", b)
  fmt.Println("e, f =", e, f)
  fmt.Println("g =", g)
  fmt.Println("h =", h) // Will print 0 (default int value)
}

// Constants in Go (fixed values that cannot be changed)
const s string = "Hello"

func constantExample() {
  fmt.Println("Constant string s =", s)

  // Declare numerical constants
  const n = 500000000
  const d = 3e20 / n // 3e20 is a float64 scientific notation

  // Print result
  fmt.Println("d =", d)
  fmt.Println("int64(d) =", int64(d)) // Convert float64 to int64
}

// Different loop types in Go
func goLoopsExamples() {
  fmt.Println("1. Traditional for loop:")
  for i := 0; i < 5; i++ {
    fmt.Println("i =", i)
  }

  fmt.Println("\n2. While-style loop (for with condition):")
  sum := 1
  for sum < 100 {
    fmt.Println("sum =", sum)
    sum += sum
  }

  fmt.Println("\n3. Infinite loop with break:")
  for {
    fmt.Println("Inside infinite loop (breaking immediately)")
    break
  }

  fmt.Println("\n4. for...range loop (over slice):")
  nums := []int{102, 103, 104, 105}

  // Iterate over index and value
  for index, value := range nums {
    fmt.Println("Index:", index, "Value:", value)
  }

  // Iterate over values only
  for _, value := range nums {
    fmt.Println("Value:", value)
  }

  // Iterate over indices only
  for index := range nums {
    fmt.Println("Index:", index)
  }
}

// Demonstrates if/else and switch statements
func controlStructures() {
  x := 10

  fmt.Println("\n5. If/Else Example:")
  if x%2 == 0 {
    fmt.Println("x is even")
  } else {
    fmt.Println("x is odd")
  }

  // declaring and initializing a new variable num, just for this if block.
  //This variable num is scoped to if-else chain only - it doesn't exist outside.
  //go allows this short declaration inside if and switch statements.
  //then the code checks conditions on num:
  if nums := 9; nums<0{
    fmt.Println(num, "is negative")
  }else if num<10{
    fmt.Println(num, "has 1 digit")
  }else{
    fmt.Println(num, "has multiple digits")
  }
  
  fmt.Println("\n6. Switch Example:")
  
  switch x {
  case 5:
    fmt.Println("x is five")
  case 10:
    fmt.Println("x is ten")
  case 20:
    fmt.Println("x is twenty")
  default:
    fmt.Println("x has an unknown value")
  }


  //Another switch example::
  switch time.No().Weekday(){
    case time.Saturday, time.Sunday:
      fmt.Println("It's the weekend")
    default:
      fmt.Println("It's weekday")
  }

  ///Another:
  t := time.Now()
  switch{
    case t.Hour()<12:
      fmt.Println("It's before noon")
    default:
      fmt.Println("It's after noon")
  }
//Another::
  whatAmI := func(i interface{}){
    switch t:=i.(type){
      case bool:
        fmt.Println("I am boolean")
      case int:
        fmt.Println("I am Integer")
      default:
        fmt.Printf("Don't know type %T\n", t) 
    }
  }

  whatAmI(true)
  whatAmI(12)
  whatAmI("HEHEHEHE")
  
}

// Entry point of the program
func main() {
  fmt.Println("=== Variable Method ===")
  variableMethod()

  fmt.Println("\n=== Constant Example ===")
  constantExample()

  fmt.Println("\n=== Loop Examples ===")
  goLoopsExamples()

  fmt.Println("\n=== Control Structures ===")
  controlStructures()
}
