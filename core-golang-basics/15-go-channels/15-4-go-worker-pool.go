/**
A worker pool in Go is a concurrency pattern used to efficiently manage a set of goroutines(workers) that process jobs from a shared queue(channel).
It helps control resource usage and avoid spawning an unbounded number of goroutines.
###### Why use a worker pool?
  - To limit the number of goroutines doing work concurrently.
  - To reuse worker goroutines for processing multiple jobs.
  - To handle large numbers of tasks efficiently and safely.
  - To implement rate limiting, throttling, or bounded concurrency.

####### Basic structure of a worker pool:
  - Job channel: A channel to send work to the workers.
  - Worker goroutines: Fixed number of goroutines that process jobs from the channel.
  - Result channel(optional): A channel to collect the output of each job.


🔍 Breakdown of the Example
jobs is a buffered channel carrying integers (jobs).

results is where each worker sends back its result.

worker() runs as a goroutine and loops over the jobs channel.

We spawn 3 worker goroutines.

After sending all jobs, we close(jobs) so workers know when to stop.

The main goroutine collects the results after all jobs are processed.

| Concept            | Explanation                                            |
| ------------------ | ------------------------------------------------------ |
| Goroutines         | Lightweight threads to run workers                     |
| Channel            | Used to distribute work and gather results             |
| Close jobs channel | Prevents deadlocks; tells workers there’s no more work |
| Blocking behavior  | Workers block on `range jobs`, waiting for work        |
| Efficiency         | Limits concurrency to a fixed number (`numWorkers`)    |

🏗️ Variations
Context for cancellation: Use context.Context to cancel workers on timeout or signal.

Rate limiting: Add delays or use time.Ticker inside workers.

Error handling: Add an error channel or wrap results in structs like type JobResult struct.


*/

package main
import (
  "fmt"
  "time"
  )

func worker(id int, jobs <-chan int, results chan<- int){
 for j:= range jobs{
   fmt.Printf("Worker %d started job %d\n", id, j)
   time.Sleep(time.Second)  
   fmt.Printf("Worker %d finished job %d\n", id, j)
   results <- j*2
 } 
}

func main(){
  const numJobs=5
  const numWorkers=3
  jobs := make(chan int, numJobs)
  result := make(chan int, numJobs)

  //start worker goroutines.
  for w:=1;w<=numWorkers;w++{
    go worker(w, jobs, result)
  }
  //send job to the job channel
  for j:=1;j<=numJobs;j++{
    jobs <- j
  }
  close(jobs)  //important close the job channel to signal no more jobs.
  //collect result.
  for a:=1;a<=numJobs;a++{
    result := <-results
    fmt.Println("Results", result)
  }
}

