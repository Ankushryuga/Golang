# Custom server with timeout:
    server := &http.Server{
      Addr:                 ":8080",
      ReadTimeout:          5*time.Second,      //max duration to read full request...
      ReadHeaderTimeout     2*time.Second      //just the header
      WriteTimeout:         10*time.Second,    //max time to write a response
      IdleTimeout:          120*time.Second,  //max time to keep idle connections
      Handler:              nil    //default multiplexer
    }

    server.ListenAndServe()

# Why timeout on server?
  Timeouts in a HTTP server are critical for building robust, secure, and high-performance web services.
# why timeout matters?
1. Prevent Slowloris Attack: A slowloris attacker keeps a connection open by sending data extremely slowy, without a ReadTimeout, your server may hang indefinitely waiting for more input.
   ✅ Fix: Set ReadTimeout to limit how long the server waits for a request.

2. Protect Against Hanging Clients: If a client starts uploading data and never finishes, your server might waste memory or threads.
  ✅ Fix: ReadHeaderTimeout (specific to reading headers) and ReadTimeout prevent clients from hogging resources.

3. Fail Fast When Backend Is Slow:If your server processes requests by calling a database or another API, those can stall. You don’t want a slow backend to tie up your HTTP connections forever.
  ✅ Fix: Use WriteTimeout to ensure the response gets sent in a reasonable time.

4. Free Up Idle Connections: Idle clients can keep connections open using HTTP keep-alive. Without an IdleTimeout, you may run out of file descriptors or memory.
  ✅ Fix: IdleTimeout limits how long to keep idle connections open.

5. Avoid Hanging Goroutines: Each connection often involves a goroutine. If timeouts are missing, slow clients or bad data can cause goroutines to accumulate and eventually crash your app.
  ✅ Fix: Timeouts help goroutines terminate cleanly and free resources.





| Feature        | Code Example            |
| -------------- | ----------------------- |
| Basic server   | `http.ListenAndServe()` |
| Custom routes  | `http.HandleFunc()`     |
| Static files   | `http.FileServer()`     |
| JSON responses | `json.NewEncoder()`     |
| Timeouts       | `http.Server{}`         |
| Method routing | `r.Method switch`       |


| Timeout             | Purpose                                      | Protection                |
| ------------------- | -------------------------------------------- | ------------------------- |
| `ReadTimeout`       | Limits full request reading (headers + body) | Slowloris, resource abuse |
| `ReadHeaderTimeout` | Only the headers                             | Hanging clients           |
| `WriteTimeout`      | Limits how long to send response             | Slow backends             |
| `IdleTimeout`       | Keeps idle connections from staying forever  | Resource exhaustion       |


