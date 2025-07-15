/***
In Golang, command line arguments are a fundamental way to provide input to a program at runtime. There are 2 primary ways to access and parse these arguments.

1. Using the os.Args slice:
  - The os package provides the Args variables, which is a slice of strings([]string) containing all the command-line arguments.
  - os.Args[0] always represents the name of the executable program itself.
  - Subsequent elements, os.Args[1:], contain the actual arguments passed by the user.
  - You can iterate over this slice to access individual arguments or use indexing to retrieve specific arguments by their position.


2. Using the flag package:
  - For more complex application requiring named flags(e.g., -file, --verbose), the built-in flag package is recommended.
  
*/



package main

import (
  "fmt"
  "os"
  )

func main(){
  fmt.Println("Program name: ", os.Args[0])
  if len(os.Args)>1{
    fmt.Println("Arguments")
    for i, args := range os.Args[1:]{
      fmt.Println("Arg", i+1,args)
    }
  }else{
    fmt.Println("No additional argument passed")
  }
}
