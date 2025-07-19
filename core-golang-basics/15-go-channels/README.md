# Bounded Concurrency 
It Limit the number of goroutines(concurrent workers) processing at the same time. 
- Prevents overwhelming system resources (CPU, memory, DB connections).
- Keeps concurrency manageable and controlled.
**Key Idea:**
    - A fixed-size pool of workers - like a worker pool in Go

# Example:
      =>
      const maxConcurrent=5
      sem := make(chan struct{}, maxConcurrent)
      for i:=0;i<100;i++{
        sem <- struct{}{}  //acquire a slot.
        go func(i int){
          defer func(){<-sem}()  //release slot
          doWork(i)
        }(i)
      }



# Throttling
Limit the rate of operations over time(e.g., only 10 requests per seconds).
- Stay under API rate limit.
- Avoid overloading downstream systems.
- Enforce fairness or delay.
**Key Point**
  Control how often something is allowed to happen.
# Example:  
      ticker := time.NewTicker(100 * time.Millisecond) // 10 per second
      defer ticker.Stop()
      
      for _, job := range jobs {
          <-ticker.C // wait for tick
          go process(job)
      }

