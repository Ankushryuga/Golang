package main

import (
    "fmt"
    "time"
)

// Worker function that accepts a receive-only channel
func worker(jobs <-chan string) {
    for job := range jobs {
        fmt.Println("Processing job:", job)
        time.Sleep(time.Second) // simulate work
    }
    fmt.Println("Worker finished")
}

func main() {
    // Create a channel of string type
    jobChannel := make(chan string)

    // Start worker in a goroutine
    go worker(jobChannel)

    // Send jobs to the channel
    jobChannel <- "Job1"
    jobChannel <- "Job2"
    jobChannel <- "Job3"

    // Close the channel to indicate no more jobs
    close(jobChannel)

    // Wait to ensure worker finishes before main exits
    // time.Sleep(4 * time.Second)
}
