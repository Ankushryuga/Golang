package main

import ("fmt"
        "time"
        )
func printA(){
  for i:=0;i<5;i++{
    fmt.Println("A",i)
    time.Sleep(time.Millisecond * 10)
  }
}

func printB(){
  for i:=0;i<5;i++{
    fmt.Println("B", i)
    time.Sleep(time.Millisecond * 10)
  }
}

func main(){
    go printA()
    go printB()

    time.Sleep(time.Second)
  }
