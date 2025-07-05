/***
go Buffered channel:  It is asyznchronous channel, sender will not block untill the buffer is full, receiver can read as long as the channel is not empty.

Buffered channels accepts a limited number of values without a corresponding receiver for those values.
## syntax:
  channel := make(chan int, 3)    // A buffered channel with capacity 3.


NOTE:::   range only works properly if the channel is closed, because it will keep waiting for new value if it's stil open..
*/

//example:

package main
import "fmt"



func main(){
  ch := make(chan int, 4)
  ch <- 1
  ch <- 11
  ch <- 111
  ch <- 1111

  /**
  ch <- 11111  //it will block or ( panic if no receiver) because buffer is full.
  */

  for {
    if ok
  }
}
