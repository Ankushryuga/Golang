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
#
