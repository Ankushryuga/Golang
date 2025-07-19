/***
- sync.WaitGroup for coordinating goroutines.
- Error handling to collect any issues.
- context.Context for timeouts or cancellation.


## Goal:
- Launch multiple workers to fetch URLs.
- Stop all work after 3 seconds.
- Collect erros if any.
*/
package main
import (
  "context"
  "fmt"
  "net/http"
  "sync"
  "time"
  )
//worker functions..
func fetchURL(ctx context.Context, url string wg *sync.WaitGroup, errors chan<- error){
  defer wg.Done()

  //Respect context timeout or cancel.
  req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
  if err != nil{
      errors <-fmt.Errorf("Request creation failed: %v", err)
      return
  }
  res, err := http.DefaultClient.Do(req)
  if err != nil{
      errors <- fmt.Errorf("Failed to fetch %s: %v", url, err)
      return
  }
  defer res.Body.Close()
  fmt.Println("Fetched: ", url, "Status:", res.Status)
}

func main(){
  //1 setup context with timeout.
  ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
  defer cancel()
  
  //2 Channels and waitgroup.
  var wg sync.WaitGroup
  errors := make(chan error, 10)  //buffered error channnel.

  //3 List of URLs 
  urls := []string{
    "https://example.com"
    "https://httpstat.us/200?sleep=1000",  //1s delay
    "https://httpstat.us/200?sleep=4000", //4s delay(will timeout)
  }

  //4 launch goroutines.
  for _, url := range urls{
    wg.Add(1)
    go fetchURL(ctx, url, &wg, errors)
  }

  //5 Wait for all goroutines in a separate goroutine.
  go func(){
    wg.Wait()
    close(errors)
  }()

  //6 Read from the error channel.
  for err := range errors{
    fmt.Println("Errors:", err)
  }

  fmt.Println("Finished all fetches")
}

/***

✅ Output Explanation
context.WithTimeout cancels all requests after 3 seconds.

Errors are sent to a channel and printed out.

WaitGroup ensures we wait for all goroutines to complete before closing the error channel.

Each fetchURL gracefully respects the context deadline.


| Feature                 | Use Case                                      |
| ----------------------- | --------------------------------------------- |
| `WaitGroup`             | Synchronize completion of multiple goroutines |
| `context.Context`       | Timeout or cancel long-running goroutines     |
| `error channel`         | Collect errors from multiple goroutines       |
| `defer wg.Done()`       | Always clean up properly                      |
| `NewRequestWithContext` | Respect cancellation and timeout              |


*/
