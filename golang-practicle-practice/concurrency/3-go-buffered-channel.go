/***
go Buffered channel:  It is asyznchronous channel, sender will not block untill the buffer is full, receiver can read as long as the channel is not empty.

Buffered channels accepts a limited number of values without a corresponding receiver for those values.
## syntax:
  channel := make(chan int, 3)    // A buffered channel with capacity 3.


NOTE:::   range only works properly if the channel is closed, because it will keep waiting for new value if it's stil open..

NOTE::: recover() only catches panics in the same goroutine.

NOTE::: without close(), range will block forever after the last value in the buffer is read, waiting for more.
closing the channel tells go that no more values will be sent, so range stops after reading all buffered values.



comma ok idiom, used to handle operations that may succeed or fail..
*/

//example:

package main
import ("fmt"
"time"
)



func main(){
  ch := make(chan int, 4)
  ch <- 1
  ch <- 11
  ch <- 111
  ch <- 1111

  /**
  ch <- 11111  //it will block or ( panic if no receiver) because buffer is full.
  */

  close(ch)
  go func(){
    defer func(){
      if r:=recover(); r!=nil{
        fmt.Println("Recovered from gorouting panic: ", r)
      }
    }()
    ch <- 11111
  }()

  //Allow time..
  time.Sleep(100 * time.Millisecond)

  for val := range ch{
    fmt.Println(val)
  }

  //using range with ok-idiom

  for{
    val, ok := <- ch{
      if !ok{
        fmt.Println("channel is closed")
        break
      }
      fmt.Println("Recived : ", val)
    }
  }
  
}

