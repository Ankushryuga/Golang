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


/**
## to see context switches:
        => go run -race filename.go

## or for profilling:
        => go tool trace

## you can trace goroutine scheduling and context switches visually.
**/
