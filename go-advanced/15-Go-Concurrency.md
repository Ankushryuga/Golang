# Goroutines:
    => A Goroutine is a thread of execution that's managed by Go's runtime. It is an essential part of Go's concurrency
       model and enables you to run functions or methods simultaneously with other functions or methods.

        It is created using the go keyword. Example: go myFunction(). They have very little overhead and are quite efficient.


# Syntax:
    => go functionName(parameters)

# Example: 
    =>
    package main
    import (
       "fmt"
       "time"
    )
    func sayHello() {
       fmt.Println("Hello goroutine!")
    }
    func main() {
       go sayHello() 
       time.Sleep(time.Second) 
       fmt.Println("Hello main function!")
    }


# Go Channels:
    => Channels is a powerful features that allows goroutines to communicate with each other and synchronize their execution.
       It provide a way to send and receive values b/w goroutines, making it easy to build concurrent programs

NOTE: It is created using the **make function**


# Syntax:
    =>
    ch := make(chan int)
    ch := make(chan int, 100)
# Use the <- operator to send and receive data.
    ch <- data (send) 
    data := <-ch (receive).


# Example:
    =>
    package main
    import (
       "fmt"
       "time"
    )
    func worker(ch chan int) {
       time.Sleep(time.Second)
       ch <- 42 // Send value into channel
    }
    func main() {
       ch := make(chan int) // Create an unbuffered channel
       go worker(ch) // Start a goroutine
       value := <-ch // Receive value from channel
       fmt.Println("Received:", value)
    }

# Go Select Statement:
    => its a control structure that allows a goroutine to wait on multiple communication operations (Channel operations).
    It is similar to a switch statement but designed for channels. The select statement blocks until one of its cases can proceed, then it executes that case.


# Syntax:
   => 
   select {
    case msg := <-ch1:
        // handle msg from channel 1
    case msg := <-ch2:
        // handle msg from channel 2
    }


# Example:
    =>
    package main
    import (
       "fmt"
       "time"
    )
    func main() {
       ch1 := make(chan string)
       ch2 := make(chan string)
       go func() {
          time.Sleep(time.Second)
          ch1 <- "Message from ch1"
       }()
       go func() {
          time.Sleep(time.Second * 2)
          ch2 <- "Message from ch2"
       }()
       for i := 0; i < 2; i++ {
          select {
          case msg1 := <-ch1:
             fmt.Println(msg1)
          case msg2 := <-ch2:
             fmt.Println(msg2)
          }
       }
    }


# Go WaitGroups:
    => A WaitGroup is a synchronization primitive provided by the sync package that allows you to wait for a collection of goroutines to finish executing. It is useful when you need to ensure that multiple goroutines have completed their tasks before continuing with the main program.


# Syntax:
    => var wg sync.WaitGroup
        wg.Add(1)
        go func() {
            defer wg.Done()
            // goroutine work
        }()
        wg.Wait()
# Example:
    =>
    package main
    import (
       "fmt"
       "sync"
       "time"
    )
    func worker(id int, wg *sync.WaitGroup) {
       defer wg.Done() 
       fmt.Printf("Worker %d starting\n", id)
       time.Sleep(time.Second)
       fmt.Printf("Worker %d done\n", id)
    }
    func main() {
       var wg sync.WaitGroup 
       for i := 1; i <= 3; i++ {
          wg.Add(1) 
          go worker(i, &wg) 
       }
       wg.Wait()
       fmt.Println("All workers done")
    }

#  Go Mutexes:    
    =>In Golang, the mutexesis provided by the sync package, are synchronization primitives used to protect shared resources and prevent race conditions when multiple goroutines access and modify those resources concurrently. A Mutex ensures that only one goroutine can access the critical section of code at a time, thus maintaining data integrity.

# Syntax:
    => 
    var mu sync.Mutex
    mu.Lock()
    // critical section
    mu.Unlock()

# Example:
    => 
    package main
    import (
       "fmt"
       "sync"
    )
    var (
       counter int
       mu      sync.Mutex
    )
    func increment(wg *sync.WaitGroup) {
       defer wg.Done()
       mu.Lock()
       counter++
       mu.Unlock()
    }
    func main() {
       var wg sync.WaitGroup
       for i := 0; i < 5; i++ {
          wg.Add(1)
          go increment(&wg)
       }
       wg.Wait()
       fmt.Println("Counter:", counter)
    }
     
