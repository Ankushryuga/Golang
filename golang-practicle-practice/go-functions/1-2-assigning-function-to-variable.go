/**
Go allows you to assign a function directly to a variable. The function should have a return type.
*/

package main
import "fmt"

func addTwoNumbers(a, b int) int{
  return a+b
}

func main(){
  sum := addTwoNumbers
  result := sum(19,30)
  fmt.Println("result: ", result) 
}
