/**
To wait for multiple goroutines to finish, we can use a wait group.
sync.WaitGroup: It's a go concurrency primitve used to wait for a collection of goroutines to finish.

# WaitGroup:
  -> It waits for a set of goroutines to finish executing. It's ideal when you're spawning goroutines and need to wait
  for all of them to complete before continuing.

# CORE API...
| Method   | Description                                                                |
| -------- | -------------------------------------------------------------------------- |
| `Add(n)` | Increases the count by `n` (typically the number of goroutines to wait on) |
| `Done()` | Decreases the count by 1 (called when a goroutine finishes)                |
| `Wait()` | Blocks until the counter becomes 0                                         |

# WaitGroup vs Channels
| Feature        | `WaitGroup`              | Channels                             |
| -------------- | ------------------------ | ------------------------------------ |
| Coordination   | Great for "wait for all" | Good for pipeline/data communication |
| Error handling | Manual                   | Can send structured results          |
| Control flow   | Only count, no data      | Can carry values (data, status)      


########## How it works?
- Add before launching a goroutine.
- Each goroutine calls Done() when done.
- Wait() blocks until all Done() calls match Add() calls.


#Note:
✅ Tip: Always call Done() with defer to avoid forgetting it.

## Common Mistakes:
| Mistake                            | Fix                                                   |
| ---------------------------------- | ----------------------------------------------------- |
| Calling `Add` inside a goroutine   | Always call `Add` **before** starting the goroutine   |
| Forgetting `Done()`                | Use `defer wg.Done()` inside the goroutine            |
| Reusing `WaitGroup` after `Wait()` | Never reuse a `WaitGroup`—use a new one for new tasks |


*/

package main
import (
  "fmt"
  "sync"
  )

func worker(id int, wg *sync.WaitGroup){
  defer wg.Done()  //signal completion.
  fmt.Printf("Worker %d starting \n", id)
  //simulate work.
  fmt.Printf("Worker %d done\n ", id)
}
/***
///######## Concurrent Fetching::
func fetchURL(url string, wg *sync.WaitGroup){
  defer wg.Done()
  res, err := http.Get(url)
  if err != nil{
    fmt.Println("Error fetching: ", url)
    return
  }
  fmt.Println(url, res.Status)
}

func main(){
  var wg sync.WaitGroup
  urls := []string{"https://google.com", "https://golang.org"}
  for _, url:=range urls{
    wg.Add(1)
    go fetchURL(url, &wg)
  }
  wg.Wait()
  fmt.Println("All fetchs done")
}
*/
func main(){
  var wg sync.WaitGroup
  for i:=1;i<=4;i++{
    wg.Add(1)    //increment counter..
    go worker(i, &wg)
  }
  wg.Wait()  //wait for all goroutines to complete.
  fmt.Println("All workers done")
}











