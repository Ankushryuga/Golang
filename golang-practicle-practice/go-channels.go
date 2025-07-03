package main
import "fmt"


/**
A channel in GO is like pipe; you can send a value into one end and receive it on the other->

1. Send or write:
ch <- 44    //send 44 into the channel.

2. Receive or read:
value := <-ch
*/

func typesOfChannel(){
  //1 unbuffered channel: its synchronous, send blocks untill a receiver is ready. useful for signal and tight coordination b/w goroutinges.
  ch1 := make(chan int)  
  go func(){
    ch1 <- 32 //waits untill receiver is ready..  writing to channel
  }()

  val := <-ch1    //reading from channel::

  //2 buffered channel: its asynchronous: send does not block untill the buffer is full., receivers can read as long as the channel is not empty.

  ch2 := make(chan int, 3)
  ch2 <- 1
  ch2 <- 2
  ch2 <- 3

  //3 Bidirectional channel(default): can be both sent to and received from., most common channel.
  ch3 := make(chan string)
  go func(){
    ch3 <- "Hello from bidirectional channel"
  }()

  fmt.Println(<-ch3)
  
  //4 send only channel:  can only send., used for abstraction or safe API design.
  var ch4SendOnly chan<- int = make(chan int)
  go func(ch chan<- int){
    ch <- 90
  }(ch4SendOnly)
  
  //5 Receive only channel://can only recieve values.
  var ch5RecvOnly <-chan int = make(chan int)
  go func(){
      temp := make(chan int)
    go func(){temp<-88}()
    ch5RecvOnly=temp
    val := <-ch5RecvOnly
    fmt.Println("Received from receive-only channel", val)
  }()
  
  //6. nil channel:
  var ch chan int

  //7. channel of custom types: You can create channels of any type: struct, slice, map, interface, etc.
  type Task struct {
    id int
  }
  taskChan := make(chan Task)
  go func(){
    taskChan <- Task{id:100}
  }()

  task := <-taskChan

  fmt.Println("Received task with Id: ", task.id)
}

func main(){
  //
   typesOfChannel() 
}

/***
| Channel Type  | Declaration           | Usage                   |
| ------------- | --------------------- | ----------------------- |
| Unbuffered    | `make(chan int)`      | Sync communication      |
| Buffered      | `make(chan int, N)`   | Async, limited buffer   |
| Bidirectional | `chan int`            | Send + Receive          |
| Send-only     | `chan<- int`          | Only send               |
| Receive-only  | `<-chan int`          | Only receive            |
| Nil channel   | `var ch chan int`     | Blocks on send/receive  |
| Custom type   | `make(chan MyStruct)` | Strongly typed channels |


*/
