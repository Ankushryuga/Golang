/***
you can access command-line arguments using os and flag package.
*/

package main
import (
  "fmt"
  "os"
  )

func main(){
  args := os.Args //args[0] is the program name.
  fmt.Println("Program name", args[0])

  if len(args)>1{
    fmt.Println("Arguments", args[1:])
  }else{
    fmt.Println("No argument provided")
  }
}
//go run main.go hello world
