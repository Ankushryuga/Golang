/***
The select statement in Go is a powerful control structure used to wait on multiple channel operations- It allows a goroutine to handle multiple communication events
without blocking on just one.



| Feature        | Note                                                               |
| -------------- | ------------------------------------------------------------------ |
| Blocking       | If no case is ready, select blocks unless `default` exists         |
| Random case    | If multiple cases are ready, one is chosen at random               |
| Loop + select  | Use `for + select` to listen forever or until a signal comes       |
| Select timeout | Use `time.After` to implement timeouts                             |
| Done signaling | Use `context.Context.Done()` in a select for graceful cancellation |


# Basic Syntax: select.
select {
  case val := <-ch1:
    fmt.Println("Received from ch1: ", val)
  case ch2 <- 42:
    fmt.Println("Sent to ch2")
  default:
    fmt.Println("No Channel is ready")
}

🧠 How It Works
  - select waits until one of the cases is ready (a channel is ready to send or receive).
  - If multiple cases are ready, one is chosen at random.
  - If no case is ready, and there is no default, it blocks.
  - If default exists and no case is ready, default is executed immediately (non-blocking).

########### Timeout pattern with select: This pattern is common for canceling slow operations. 
func timeoutPattern(){
  select{
    case res := <-ch:
      fmt.Println("Received:", res)
    case <-time.After(2*time.Second):
      fmt.Println("Timeout")
  }
}

########### Default Case for Non-Blocking: This prevents blocking but can lead to busy waiting if misused.
select{
  case msg := <-ch:
    fmt.Println("Received: ", msg)
  default:
    fmt.Println("No data available")
}

########## Looping with select: Used in goroutines to listen for multiple channel signals, including cancellation.
for {
  select{
    case msg := <-ch1:
      fmt.Println("Received:", msg)
    case <-quit:
      fmt.Println("Shutting down")
      return
  }
}



*/

package main
import (
  "fmt"
  "time"
  )

func main(){
  ch1 := make(chan string)
  ch2 := make(chan string)

  go func(){
    time.Sleep(1*time.Second)
    ch1 <- "from ch1"
  }()

  go func(){
    time.Sleep(2*time.Second)
    ch2 <- "from ch2"
  }()

  select{
    case msg1:=<-ch1:
      fmt.Println(msg1)
    case msg2:=<-ch2:
      fmt.Println(msg2)
  }
  //O/P: from ch1 (likely, because it's faster).
}
