/***
When working with multiple goroutines to finish, you may prefer to use a WaitGroup.
*/

package main
import (
    "fmt"
    "time"
  )

func worker(done chan bool){
  fmt.Println("working....")
  time.Sleep(time.Second)
  fmt.Println("done")
  done <- true
}
func main(){
  done := make(chan bool, 1)    
  go worker(done)  //start a worker goroutine, giving it the channel to notify on. Block untill we receiver a notification from the worker on the channel.
  <-done    //if we remove <-done ths program will exit before the worker even started.
}
