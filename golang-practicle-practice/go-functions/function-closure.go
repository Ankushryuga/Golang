/**
Go function closure is a function value that references variables from outside its body. The closures bind these variables 
and make them accessible even after closing the outer function.

## Creating a closure::
counter := func() func() int{
  count := 0
  return func(){
    count++
    return count
  }
}

## use:
increment := counter()
*/

package main
import "fmt"

func main(){

  // creating a closure::
  updateCounter := func() func() int{
    count := 100
    return func() int{
      count++
      return count
    }
  }

  //using
  increment := updateCounter()

  fmt.Println(increment())
  fmt.Println(increment())
}
