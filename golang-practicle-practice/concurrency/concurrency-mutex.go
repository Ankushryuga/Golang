package main

import (
	"fmt"
	"time"
	"sync"
)
// var counter int
var (
    counter int
    mu sync.Mutex
    )
func increment() {
    mu.Lock()
	counter++ // NOT thread-safe
	mu.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Final counter:", counter) // inconsistent output
}
