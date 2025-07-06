/***
Anonymous function also known as lambda function or function literals, are defined without a name..
Go Anonymous functions allow you to define temporary, one-time use functions exactly where you want them, without giving them a name.

## Ideal for:
1. Executing code on the spot.
2. Satisfying variable usage from surrounding code.
3. Scheduling fast jobs in goroutine

Variable capturing:
This means that even after the surrounding function has returned, the closure retains access to those variables—not just their values at the time of capture, but the variables themselves.
*/


package main

import "fmt"

/**
Anonymous function can access variables from their enclosing scope, creating closure...
*/
func counter() func() int{
  count := 0
  return func() int{
    count++
    return count
  }
}

func multiplier(factor int) func(int) int{
  return func(val int) int{
    return val * factor
  }
}


//variable capcture::
func variableCapture(){
  var focus []func()
  for i:=0;i<4;i++{
    focus = append(focus, func(){
      fmt.Println(i)
    })
  }

  for _, f := range focus{
    f()
  }
}

func main(){
  greet := func(){
    fmt.Println("Tutorials point")
  }
  greet()  //call the anonymous function.
  
  add := func(a, b int) int{
    return a+b
  }

  //Invoke function expression::
  //without parameter:
  func (){
    fmt.Println("Executing immediately")
  }()

  //with parameters:
  result := func(a, b int) int{
    return a * b
  }(50, 10)

  fmt.Println(result)
  
  sum := add(34, 55)
  fmt.Println("sum: ", sum)


  c := counter()
  fmt.Println(c())    //1
  fmt.Println(c())  //2
  fmt.Println(c())  //3

  double := multiplier(2)
  tripler := multiplier(3)

  fmt.Println(double(6))  //2*6=12
  fmt.Println(tripler(4))  //3*4=12

  variableCapture()
}
