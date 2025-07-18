/***
Goroutines are lightweight threads  managed by the Go runtime. They allows you to run functions concurrently..
meaning they can run independently of the main program flow.

## What is Goroutine?
A goroutine is a function or method running asynchronously in the same address space.

# Basic Syntax:
go functioname(arg...) //this starts a new goroutine

## Key concepts::
1. concurrency vs parallelism:
  - Concurrency: multiple task making progress at the same time (handled by the go schedulers).
  - Parallelism: Actually running on multiple CPU cores (depends on GOMAXPROCS)

2. GOMAXPROCS:
  runtime.GOMAXPROCS(n):  // sets the number of OS threads that can run Go code simultaneously,
  By default, it's set to the number of logical CPUs.


## Lightweight:
  - They start with a very small stack and can grow/shrink dynamically.
  - You can spawn thousands or even millions of goroutines.
## Managed by Go Runtime:
  - The Go scheduler(M:N scheduler) maps multiple goroutines onto fewer OS threads.
  - It handles pausing/resuming, garbage collection, communication and more.
## Non-blocking execution:
  - Goroutines run concurrently, not necessarily in parallel.
  - Actual parallelism depends on available CPU cores and GOMAXPROCS.


##### How Goroutines work(GMP Model)?
G (Goroutine) – represents the execution context (stack, instruction pointer).

M (Machine) – a thread that executes Go code.

P (Processor) – schedules goroutines and runs them on machines.

###### Flow:
  - You spawn a goroutine.
  - It's put into a queue managed by a P.
  - An available M runs it.
  - Goroutines are preempted, paused, and resumed as needed.
  - Go maps many goroutines to fewer OS threads: M:N scheduling.
*/

package main

//Spawing many goroutines.
func ManyGoRoutines(){
  for i:=0;i<5;i++{
    go func(i int){
      fmt.Println("Goroutine", i)
    }(i)
  }
  time.Sleep(time.Second)    //always pass loop variables as parameters., or you'll get unexpected result...
}


/***
## Synchronization:

1.Channels : Safe communication.
ch := make(chan int)

go func(){
  ch <- 42  //sending.
}()
val := <-ch  //receive

fmt.Println(val)


2. WaitGroup : Wait for multiple goroutines to finish
var wg sync.WaitGroup

for i:=0;i<5;i++{
  wg.Add(1)
  go func(i int){
    defer wg.Done()
    fmt.Println("worker", i)
  }(i)
}
wg.Wait()    //blocks until all done..


3. Mutexes : Protect shared memory.
var mu sync.Mutex
counter := 0
go func(){
  mu.Lock()
  counter++
  mu.Unlock()
}()       //Use sync.Mutex or sync.RWMutex to avoid race conditions 

############################### Common Pitfals.

1. Race Conditions:
//Not safe:::
counter :=0
for i:=0;i<100;i++{
  go func(){
    counter++  
  }()
}
//Fix use Mutex or Channels.

Also run with go run -race to detect data races.


2. Goroutine Leaks: A goroutine that never finishes because it's waiting for something that will never happen -e.g., blocked on a channel forever.
Example:
func goroutineLeak(){
  ch := make(chan int)
  go func(){
    val := <-ch  //Block forever.
    fmt.Println(val)
  }()
}

//If no one ever sends to ch, the goroutine **leaks**.
Detection:
  - Use profiling (pprof)
  - Watch for high memory/CPU
  - Ensure channels are closed or always consumed


### TO gracefully cancel goroutines.
# Use context.Context

ctx, cancel := context.WithCancel(context.Background())

go func(ctx context.Context){
  for {
    select {
      case <-ctx.Done():
        fmt.Println("Cancelled")
        return
      default:
        //default work
    }
  }
}(ctx)

time.Sleep(time.Second)
cancel()  //signal cancellation.
*/












