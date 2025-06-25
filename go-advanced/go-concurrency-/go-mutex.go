/**
Go Mutex:
it provided by sync package, are synchronization primitives used to protect shared resources and prevent race conditiions.
when multiple goroutines access and modify those resources concurrently. A Mutex ensures that only one goroutine can
access the critical section of code at a time, thus maintaining data integrity.


# syntax:
var mu sync.Mutex
mu.Lock()
//critical section

mu.Unlock()
*/

package main

import (
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
