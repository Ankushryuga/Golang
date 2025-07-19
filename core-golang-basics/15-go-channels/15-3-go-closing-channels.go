/****
Closing a channel Indicates that no more values will be sent to it, this can be useful to communicate completion to the channels receivers.

Reading from a closed channel succeeds immediately, returning the zero value of underlying type, this optional second return value is true, 
if the value reeived was delived by a successful send operation to the channel, or false if it was zero value generated because the channel is closed and empty.

*/

package main
import "fmt"

func main(){
  jobs := make(chan int, 5)
  done := make(chan bool)

  go func(){
    for {
      j, more := <-jobs
      if more{
          fmt.Println("received job", j)
      }else{
          fmt.Println("received all job")
          done <- true
          return
      }
    }
  }()

  for j:=1;j<=5;j++{
    jobs <- j
    fmt.Println("sent job", j)
  }
  close(jobs)
  fmt.Println("sent all jobs")
  <-done

  _, ok := <-jobs
  fmt.Println("received more jobs:", ok)
}
