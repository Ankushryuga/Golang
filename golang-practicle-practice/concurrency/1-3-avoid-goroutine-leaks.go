/***
Problem: If a goroutine is waiting on a channel that will never send, it's leaked (it lives forever):

## solution:
1. Always close channels if you own them and know no more data will be sent.
2. Use select with context.Context to exit if needed.
3. Use for range ch to gracefully exit when channel is closed.
*/


package main
import "fmt"

func leakyWorker(ch <-chan string){
  for msg := range ch{
    fmt.Println(msg)
  }
}

func main(){
  ch := make(chan int)
  go leakyWorker(ch)
  //program exists or channel is never closed-goroutine is stuck..
}
