/***
## some tricky thing about go channels:
  1. Channel Basics  : chan T - send(ch <- val), receive (val := <-ch), and close(ch).
  2. Bufferd (make(chan int, capacity)) vs unbuffered (make(chan in)): => make(chan int, 5) vs make(chan in)
  3. Closing channels: Not required unless you need to signal to receivers that no more data will come.
  4. Range over channels: Used to iterate untill the channel is closed.

# q1: what happens if you send to a closed channel?
  => Panic.

# q2: what happens if you receive from a closed channel?
  => Returns a zero value immediately

# q3: difference b/w buffered and unbuffered channels in go?
  =>
  Buffered channels allow sending without a receiver if space is available, unbuffered channels block until a receiver is ready.

# q4: How do you avoid deadlocks when using channels:
  => ensure gorouting are reeading and writing properly, use select, and avoid circular dependencies.

# q5: Why use select in Go?
  => to handle multiple channel operations and avoid blocking indefinitely.
*/



/***
## Some more Advanced::::: 
1. sync.WaitGroup vs channels for synchronization.
2. Use of context.Context with channels to cancel goroutines.
3. Avoiding goroutine leaks with channel closure.
4. Pipeline patterns using channels.


*/
package main
import ("fmt"
  "time"
        )
//1 basic
func basicGoRoutineChannel(){
  ch := make(chan string)
  go func(){
    ch <- "hello"    //sending
  }()
  msg := <-ch    //receving
fmt.Println(msg)
  //practice: modify this to use buffered, unbuffered or sending multiple message, and use goroutine inside for loop.
}

//2 worker pool using channels: worker function (id, send channel, receive channel)...
func worker(id int, jobs<-chan int, results chan <-int){
  for j:= range jobs{
    fmt.Println("Worker", id, "processing job", j)
    time.Sleep(time.Second)
    results <- j*2
  }

  //practice create 3 workers, 5 jobs, collect and print all resutls).
}

//3 Implement a timeout using: select

func selectExample(){
  ch := make(chan int)
  select {
    case res := <-ch:
      fmt.Println("Received: ", res)
    case <-time.After(2 * time.Second):
      fmt.Println("TImeout")
  }

  //practice: Simulate a  long-running operation and add timeout handling..
}


//4 fan-in pattern:   its a go concurrency pattern where multiple input channels are merged into a single output channel).
//producer: it's a goroutine that continuously send messages into a channel, - like a "data-producing machine".
func producer(msg string) <-chan string{    //it returns a receive-only channel : <-chan string., program can receive data from this channel and treat it as an independent stream of messages.
  
  ch := make(chan string)  
  go func(){
    for i:=0; ;i++{
      ch <- fmt.Sprintf("%s, %d", msg, i)
      time.Sleep(time.Millisecond * 500)
    }
  }()
  return ch
}

func fanIn(input1, input2<-chan string) <-chan string{
  out := make(chan string)
  go func(){
    for{
      select{
        case msg := <-input1:
          out <- msg
        case msg := <-input2:
          out <- msg
      }  
    }
  }()
  return out
}




func main(){

  ch1 := producer("Producer-A")
  ch2 := producer("Producer-B")

  merge := fanIn(ch1, ch2)
   for i:=0;i<10;i++{
     fmt.Println(<-merge)
   }


  //5: Channel Closing and Range::
  ch := make(chan int)
  go func(){
    for i:=0;i<5;i++{
      ch <- i
    }
    close(ch)
  }()

  for val := range ch{
    fmt.Println(val)
  }
  //practice: try sending after close, handle that gracefully...
}



