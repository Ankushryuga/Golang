/*
*
WaitGroup is a synchronization primitive provided by the
sync package that allows you to wait for a collection of goroutines
*/
package goconcurrency

import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n", id)
}
func Helper() {
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All Workers done")
}
