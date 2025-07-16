/**
- Go arrays are fixed-size sequences of elements of the same type.
- Their size is part of the type (e.g., [5]int and [4]int are different types).
*/

package main

import "fmt"



//Go Array Example::
func goArrayExample() {
  // 1st way: Declare an array of 5 integers (default values are 0)
  var a [5]int
  fmt.Println("array: ", a)

  // Assign a value to an index
  a[4] = 100
  fmt.Println("after set: ", a)
  fmt.Println("get element at index 4:", a[4])
  fmt.Println("length of array a:", len(a))

  // 2nd way: Declare and initialize an array at the same time
  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("declared with values: ", b)

  // 3rd way: Use [...] to let Go infer the length from the number of elements
  b = [...]int{1, 2, 3, 4, 5}
  fmt.Println("auto-sized array: ", b)

  // Array with specific index values: index 3 gets 400, rest get default or specified
  b = [...]int{100, 3: 400, 500} // index 0=100, index 3=400, index 4=500
  fmt.Println("index-specified array: ", b)

  // Declare a 2D array with 2 rows and 3 columns
  var twoDArray [2][3]int


  //in go range is used for iterables (like arrays or slices). 2 is an int, so you can't range over it.
  // Fill the 2D array with values based on i+j
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoDArray[i][j] = i + j
    }
  }
  fmt.Println("2D array with i+j values:", twoDArray)

  // Initialize a 2D array directly
  twoD := [2][3]int{
    {1, 2, 3}, // Row 0
    {1, 2, 3}, // Row 1
  }
  fmt.Println("2D array with fixed values:", twoD)
}




//Go Slice Example:::
func goSliceExample(){
  var slice1 []int
  
}


func main(){
  goArrayExample()
}
