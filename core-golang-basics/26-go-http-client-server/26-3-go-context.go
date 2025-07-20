/***
The context package is used to manage deadlines, timeouts, cancellation signals, and request-scoped values - especially across API calls, goroutines, and servers.
It's essential for writing concurrent, cancelable, and resource-safe applications.

#Why context?
| Use Case                     | Why It Matters                                 |
| ---------------------------- | ---------------------------------------------- |
| **Timeouts / Deadlines**     | Automatically cancel long-running operations   |
| **Cancellation propagation** | Gracefully stop work when parent cancels       |
| **Scoped values**            | Pass request IDs, auth tokens, etc. down chain |
| **Prevent goroutine leaks**  | Avoid runaway background workers               |
*/



package main
import (
  "context"
  "fmt"
  "log"
  "net/http"
  "time"
  )

//hanlder handles incoming requests with context..
func handler(w http.ResponseWriter, r *http.Request){
  ctx := r.Context()
  log.Println("Request started")
  
  //create a new context with a timeout of 3 seconds.
  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
  defer cancel()
  
  //simulate long-running work in a goroutine.
  done := make(chan string, 1)
  go func(){
    time.Sleep(5*time.Second)  //pretent this is slow DB/API
    done <- "work finished"
  }()

  select {
    case result :=<-done:
      fmt.Println(w, result)
    case <-ctx.Done():
      log.Println("Request canceled:", ctx.Err())
      http.Error(w, "⏳ Request timeout or canceled", http.StatusRequestTimeout)
  }
}


func main(){
  //create a new http server.
  server := &http.Server{
    Addr:                      ":8080"
    ReadTimeout:               10*time.Second
    WriteTimeout:              10*time.Second
    Handler:                   http.HandlerFunc(handler),    //all requests go here.
  }

	fmt.Println("🚀 Server running at http://localhost:8080")
  if err!= server.ListenAndServe(); err!=nil{
    fmt.Println(err)
  }
}
