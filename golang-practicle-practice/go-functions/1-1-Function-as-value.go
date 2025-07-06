/**
Go allows you to assign a function directly to a variable. The function should have a return type.
*/
/**
go have flexibility to create functions on fly and use them as value...
*/

package main
import ("fmt"
"math"
)

func addTwoNumbers(a, b int) int{
  return a+b
}



//passing a function as an argument::
/**
you can pass a function to a function as an argument, you just need to use the function declaration as an argument..
*/
func calculation(x int, y int, op func(int, int) int) int {
  return op(x, y)
}
func multiplyNumbers(x, y int) int{
  return x*y
} 

//Returning a function as a value
/***
A function can also return a function as a value in the Go language, It is useful when you want to 
return an expression calculated by creating any function.
*/
// returning a function as a value..
func returnCalculation(factor int) func(int) int{
  return func(val int) int{
      return  factor * val
  }
}
                 
func main(){
  sum := addTwoNumbers
  result := sum(19,30)
  fmt.Println("result: ", result) 


   getSquareRoot := func(x float64) float64{
    return math.Sqrt(x)
  }
  //use the functions::
  fmt.Println(getSquareRoot(9))

  returnCalResult := returnCalculation(2)
  result1 := returnCalResult(20)
  fmt.Println("result1: ", result1)
}
