/**
Go allows you to assign a function directly to a variable. The function should have a return type.
*/
/**
go have flexibility to create functions on fly and use them as value...
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


   getSquareRoot := func(x float64) float64{
    return math.sqrt(x)
  }
  //use the functions::
  fmt.Println(getSquareRoot(9))
}
