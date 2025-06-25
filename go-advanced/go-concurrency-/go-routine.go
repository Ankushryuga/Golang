/**
GoRoutine:
its a thread of execution that's managed by go's runtime, it is essential part of Go's concurrency model and enables you to run functions or methods simulatenously with other functions or methods.

# it is create using "go" keyword:
*/

/*
*
Go Channels:
channels are like pipe, that allow you to communiate b/w goroutine safely and synchronize their execution.

They provide a way to send and receive values b/w goroutines, making it easy to build concurrent programs.
it is created using make function:

#Syntax:
ch := make(chan int)
ch := make(chan int, 100)

use the <- operator to send and receive data.

ch <- data (send)
data := <- ch (receive)
*/

/*
*
Go Select Statement:

	=> its a control structure that allows a goroutine to wait on multiple communication operators (channel operations)
	it is similar to a switch statement but designed for channels, the select statement blocks until one of its cases can proceed, then it executes that case.
	select {
		case msg := <- ch1:
		//handle msg from channel 1
		case msg := <- ch2:
		//handle msg from channel 2
	}
*/

/*
*
Go WaitGroups:

	=> a waitgroup is a synchronization primitive provided by the sync package that allows you to wait for a collection of goroutines to finish executing.
	it is useful when you need to ensure that multiple goroutines have completed their tasks before continuing with the main program:

	#syntax:
*/
package goconcurrency

import (
	"fmt"
	"time"
)

func gorouteCheck(ch chan int) {
	time.Sleep(time.Second)
	ch <- 42 //send value into channel
}

func main() {
	// ch := make(chan int)
	// //go routine
	// go gorouteCheck(ch)
	// value := <-ch //receive:
	// fmt.Println("received: ", value)

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
