/***
Problem: Long-running goroutines can leak or keep running when no longer needed.

## solution:
Use "context.withCancel" to tell goroutines to exit early..
*/

/**
  ## When to use:
  1. For cancellation of goroutines after timeout, error or shutdown.
  2. Web servers, pipelines and background workers.
**/

package main

import (
  "context"
  "fmt"
  "time"
  )

func worker(ctx context.Context){
  for{
    select{
      case <-ctx.Done():
        fmt.Println("Worker cancelled")
        return
      default:
        fmt.Println("working...")
        time.Sleep(time.Second)
    }
  }
}


func main(){
  ctx, cancel := context.WithCancel(context.Background())
  go worker(ctx)

  time.Sleep(3 * time.Second)
  cancel()    //signal the worker to stop.

  time.Sleep(1 * time.Second)  //wait to see o/p
  
}
