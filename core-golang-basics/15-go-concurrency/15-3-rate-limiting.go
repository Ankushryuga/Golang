package main
import ("fmt"
        "time"
        )

func main(){
  requests := make(chan int, 5)
  for i:=1;i<=5;i++{
    requsts <- i
  }
  close(requests)
  limiter := time.Tick(200*time.Millisecond)

  for req := range requests{
    <-limiter
    fmt.Println("request", req, time.Now())
  }

  burstyLimiter := make(chan time.Time, 3)

  for range 3{
    burstyLimiter <- time.Now()
  }

  go func(){
    for t:= range time.Tick(200*time.Millisecond){
      burstyLimiter <- t
    }
  }()

  burstyRequests := make(chan int, 5)
  for i:=1;i<=5;i++{
    burstyLimiter <- i
  }
  close(burstyLimiter)

  for i:=1;i<=5;i++{
    burstyRequests <- i
  }
  close(burstyRequests)
  for req:= range burstyRequests{
    <- burstyLimiter
    fmt.Println("Request", req, time.Now())
  }
}





