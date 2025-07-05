/**
sync.WaitGroup: used to wait for a group of goroutines to finish, 
you can use channels to signal completion, though its more manual..

## When to use:
  1. sync.WaitGroup for simple wait-until-done logic
  2. channels for more complex synchronization (like status, updates, results, etc)
*/

package main

import (
  "fmt"
  "sync"
  )

func worker(id int, wg *sync.WaitGroup){
  defer wg.Done()  ///Notify Waitgroup that this goroutine is done.
  fmt.Println("worker ", id , " is working")
}

//using channel for synchronization::
func worker1(ch chan bool){
  fmt.Println("working")
  ch <- true  //send signal when done..
}

func main(){
  var wg sync.WaitGroup
  for i:=1; i<=5;i++{
    wg.Add(1)    //Add one task to wait for.
    go worker(i, &wg)  
  }
//using channel for sync
  ch1 := make(chan bool)
  go worker1(ch1)

  <-ch1  //wait for signal.
  fmt.Println("Done")
//end

  wg.Wait()  //Block untill all workers are done.
  fmt.Println("All workers finished")
}

