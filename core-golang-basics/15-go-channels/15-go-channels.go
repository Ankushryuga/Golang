/***
Channel: Its like a pipe that goroutines communicate with each other and synchronizes.



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


*/
