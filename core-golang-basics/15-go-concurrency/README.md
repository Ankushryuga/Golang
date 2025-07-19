##  <- channel name:
it means you are performing a receive operation from a channel but discarding the value, you're only interested in waiting for a signal, but not the actual data.


## What it means (<- channelname)?
- you're saying **"Block until the channel has something**, but I don't care what the value is"
- It's used for **synchronization or signaling**, not for tranferring data.

## Example:
    =>
      done := make(chan struct{})
      go func(){
      //do some work
      time.Sleep(1*time.Second)
      done<-struct{}{}//send signal.
      }()
      <-done //wait for signal.
      fmt.Println("WOrk is done")

      The main goroutine waits at **<-done** until it receives the signal from the worker.


⚠️ Why Use It?
  - You want to block/wait, but you don’t need the value sent.
  
  - The channel is being used for event notification, not data transmission.
  
  - Common with chan struct{} — a zero-sized type that takes no memory.

| Code                | Meaning                                         |
| ------------------- | ----------------------------------------------- |
| `<-ch`              | Receive a value from `ch` (and block if needed) |
| `val := <-ch`       | Receive and store the value                     |
| `<-done` (no value) | Just wait/block for a signal                    |


## context.Context:
      select {
        case <-ctx.Done():
          fmt.Println("Context cancelled or timedout")
      }
      - This is saying: "Wait until context is canceled or times out", without caring about the actual error.


