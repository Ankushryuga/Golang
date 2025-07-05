/**
What is pipeline:
  => A pipeline connects stages of computation with channels, Each stage does work and passes the result to the next.
*/


package main

import "fmt"


//stage 1: Generate numbers:
func generator(num ...int)<-chan int{
  out := make(chan int)
  go func(){
    for _, n:= range num{
      out <- n
    }
    close(out)
  }()
  return out
}

//stage 2: square number:
func square(num <-chan int) <-chan int{
  out := make(chan int)
  go func(){
      for n := range num{
        out <- n*n
      }
    close(out)
  }()
  return out
}


//stage 3: print..

func main(){

  gen := generator(1,2,3,4)
  sqr := square (gen)

  for result := range sqr{
    fmt.Println(result)
  }
}
