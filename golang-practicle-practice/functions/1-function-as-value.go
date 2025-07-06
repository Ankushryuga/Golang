/**
go have flexibility to create functions on fly and use them as value...
*/

package main

import "fmt"

func main(){
  getSquareRoot := func(x float64) float64{
    return math.sqrt(x)
  }
  //use the functions::
  fmt.Println(getSquareRoot(9))
}
