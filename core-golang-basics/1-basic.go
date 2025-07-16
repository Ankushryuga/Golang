/**
This program demonstrates:
- Variable declarations
- Constants
- Different types of for loops
- If/else condition
- Switch statement
*/

package main

import "fmt"

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
