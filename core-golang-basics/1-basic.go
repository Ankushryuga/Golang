/***
This page is contains variables, constants, for, if/else, switch.
*/
package main

import "fmt"
//Variables::
//var i int=10
func variableMethod(){
  var a = "initial"
  b := 5
//  c, d := 1,2
  var e, f int = 1,2
  fmt.Println(e, f)

  var g=true
  var h int
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(g)
  fmt.Println(h)
}

//Constants:: const declares a constant value.
const s string = "Hello"
func constantExample(){
  fmt.Println(s)
  const n=500000000
  const d=3e20/n
  fmt.Println(d)
  fmt.Println(int64(d))
}
