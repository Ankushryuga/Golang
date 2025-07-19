/***
Channel: Its like a pipe that goroutines communicate with each other and synchronizes.

⚠️ ########## Gotchas & Best Practices
  - Never send on closed channels — panics.
  - Range loops: only exit after close; if no close channel may block forever.
  - Select: prefer non-default, blocking selects unless explicit non-block behavior is needed.
  - Buffered channels: understand memory vs. back-pressure tradeoff.
  - Closing semantics: only the sender should close. Receivers should only read until closed.


| Concept                                      | Purpose                                                              |
| -------------------------------------------- | -------------------------------------------------------------------- |
| `make(chan T[, N])`                          | Create channel (buffer size N if provided)                           |
| `ch <-` / `<- ch`                            | Send / receive                                                       |
| `close(ch)`                                  | Stop sending                                                         |
| `range ch`                                   | Receive until channel closes                                         |
| `select { … }`                               | Multi-way channel operation                                          |
| Directional channels                         | Enforce safe send-only/receive-only use                              |
| Patterns: pipelines, fan‑in/out, worker pool | Structuring concurrency flow                                         |
| Best practices                               | Only sender closes; avoid races; manage blocking behavior explicitly |




## Sending & Receiving:
ch := make(chan int)
go func(){
  ch <- 42    //sending..
}()

val := <-ch  //receiving..
fmt.Println(val)

## If no goroutine is ready to receive, the send blocks.
## If no goroutine is ready to send, the receive blocks.
## This ensures safe communication without explicit looks.

############# Buffered vs Unbuffered Channels
**********- Unbuffered channels (default)
            - Block both sender and receiver untill both are ready.
            - Good for synchronization
              # Example:
                ch := make(chan string)
**********- Buffered channels
            - Sender only blocks if the buffer is full.
            - Receiver only block if buffer is empty.
            # Example:
              ch := make(chan string, 2)  //Buffer size = 4.
              ch <- "one"
              ch <- "two"
              // ch <- "three" //Blocks if buffer is full



############### Directional Channels
For type safety and clarity..
func send(ch chan<- int){ ch <- 100 }  //send only..
func receive(ch <- chan int)  {fmt.}   //receive only


################# Range Loop: Signal end of data cleanly with a range loop
close(ch)
for v := range ch{
  fmt.Println(v)
}
//This exits automatically when channel is closed and drained..



############## Select Statement:
Used to wait on multiple channel operations.
select {
  case msg := <-ch1
    fmt.Println("Received:", msg)
  case ch2 <- "hi":
    fmt.Println("Send")
  default:
    fmt.Println("No communications")
}

--- select blocks until one case is ready.
--- default executes immediately if none are ready.
-- you can use time.After for timeouts.



##### Timeouts and Cancellation:
Typical pattern using select with time.After:

select {
  case v:= <-ch:
    fmt.Println("got", v)
  case <-time.After(2*time.Second):
    fmt.Println("timeout")
}
// Or use context.Context with Done() channel for cancellation.



### Common Patterns:
1. Pipeline  - Chaining stages via channels:
  func stage1(in <-chan int) <-chan int{...}
  func stage2(in <-chan int) <-chan int{...}
  out := stage2(stage1(input))
  for v := range out{...}

2. Fan-out/Fan-in  - Spawn multiple workers consuming from one "in" channel, and merge outputs  into an "out" channel.
3. Worker pool - fixed worker goroutines consume tasks concurrently but limit concurrency.
4. Rate Limiting - use a filled buffer of time, Tick or a bounded channel channel to throttle work.

*/



package main
import ("fmt"
        "time"
        )

func worker(id int, job <-chan int, results chan<-int){
  for j := range job{
      fmt.Printf("worker %d started job %d\n", id, j)
      time.Sleep(time.Second)
      fmt.Printf("Worker %d finished job %d\n", id, j)
      result <- j*2
  }
}

func main(){
  job := make(chan int, 5)
  result := make(chan int, 5)
  for w :=1; w<=3; i++{
    go worker(w, job, result)
  }

  for j:= 1; j<=5; j++{
      job<-j
  }
  close(job)

  for a:=1; a<=5; a++{
    fmt.Println("Result:", <- results)
  }
}
