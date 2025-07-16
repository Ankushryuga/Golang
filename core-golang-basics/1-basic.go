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


//Loops: golang uses single looping construct the for loop..
func goLoopsExamples(){
  //1. Traditional for loop.
  for i:=0;i<5;i++{
    fmt.Println(i)
  }
  //2. while loop equivalent:
  sum := 1
  for sum<100{
    fmt.Println(sum)
    sum+=sum
  }

  //3. Infinite loop: Omitting all 3 components create an infinite loop, which can be terminated using break statement.
  for {
    fmt.Println("Infinite Loop Example")
    break
  }

  //4. for ...range loop: 
  nums := []int{102,103, 104, 105}
  for index, value := range nums{
    fmt.Println("Index: ", index, "value:", value)
  }
  //To iterate only value, you can use the blank identifier (_) for the index.
  for _, value := range nums{
    fmt.Println("value: ", value)
  }
  //To iterate only index..
  for index, _ := range nums{
    fmt.Println("Index: ", index)
  }
  
}
