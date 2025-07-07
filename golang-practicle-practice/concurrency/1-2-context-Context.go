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
/**
## Context:::
context: it manages concurrent ooperations, particularly across API boundaries and goroutines.
context is used to carry deadlines, cancellation signals, and request-scoped values across different parts of a program.

## Context Switching:
Its a OS system concept that descipes the process of saving the state (or "context") of a currently executing process 
or thread and loading the state of another process or thread, allowing the CPU to switch its attention b/w them..

In Go, context switches occur at the goroutine level. The Go runtime's scheduler manages the execution of goroutines on OS threads.



*/
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
