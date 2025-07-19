/***
Timers: are for hen you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.

Tickers use a similar mechanism to timers: a channel that is sent values. 

Tickers can be stopped like timers, once a ticker is stopped it wont'r receive any more values on its channel.
*/

package main
import (
    "fmt"
    "time"
  )

func main(){
  ticker := time.NewTicker(500 * time.Millisecond)
  done := make(chan bool)

  go func(){
    for{
      select{
        case <-done:
          return
        case t:=<-ticker.C:
          fmt.Println("Tick at", t)
      }
    }
  }()

  time.Sleep(1600 * time.Millisecond)
  ticker.Stop()
  done <- true
  fmt.Println("Ticker stopped")
}
