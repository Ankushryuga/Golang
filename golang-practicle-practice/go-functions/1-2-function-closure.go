/***
Go function closure: Its a function value that references variable from outside its body.
The closure bind these variables and make them accessible even after closing the outer function.
*/

package main

import "fmt"

func main(){
  updateCounter := func(initial int) func() int{
    count := initial  //initialize count with the passed value..
    return func() int{
      count++
      return count
    }
  }

  //create a closure with an initial value..
  inc_x := updateCounter(100)

  //use the closure:
  fmt.Println(inc_x())
  fmt.Println(inc_x())

  //create another closure with an initial value..
  int_y := updateCounter(200)
  //use the closure::
  fmt.Println(int_y())
  fmt.Println(int_y())
}
