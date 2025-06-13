# Select:
    => Go select statement is useful to handle multiple channel operations and selects the ready channel against multiple channels.

# Select Statement:
    => When you are working with multiple channels in Go language, the select statement waits on multiple channel operations and executes 
    the code block for the channel that is ready.

# Rule for using a select statement in Go:
    =>
    1. You can have any number of case statements within a select, each case is followed by the value to be compared to and a colon.
    2. the type for a case must be the a communication channel operation.
    3. When the channel operation occured the statements following that case will execute, No break is need in case statement.
    4. A select statement can have an optional default case, which must appear at the end of the select, the default can be used for perfoming
    a task when none of the case is true, No break is needed in the default case.


# Example:
      =>
      package main
      import "fmt"
      func main(){
      var c1, c2, c3 chan int
      var i1, i2 int
      select {
        case i1 = <-c1:
          fmt.Printf("received ", i1, " from c1\n")
        case c2 <- i2:
          fmt.Printf("Send ", i2, " to c2\n")
        case i3, ok := (<-c3):  //same as: i3, ok := <-c3
          if ok {
            fmt.Printf("received ", i3, " from c3\n")
          } else {
            fmt.Printf("c3 is closed \n")
          }
          default: 
            fmt.Printf("no communication\n")
      }
      }


## Nested Select Statement:
    =>A nested select statement allows you to use one or more select statements inside a select statement.

## Example:
    =>
      package main
      import (
      	"fmt"
      	"time"
      )
      
      func main() {
      	// Defineing channels
      	ch1 := make(chan string)
      	ch2 := make(chan int)
      
      	// Goroutine to send data to channel1
      	go func() {
      		time.Sleep(1 * time.Second)
      		ch1 <- "I am in channel 1"
      	}()
      
      	// Goroutine to send data to channel2
      	go func() {
      		time.Sleep(3 * time.Second)
      		ch2 <- 10
      	}()
      
      	// Nested select statement
      	select {
      	case ch_data1 := <-ch1:
      		fmt.Println("Received from channel1:", ch_data1)
      		select {
      		case ch_data2 := <-ch2:
      			fmt.Println("Received from channel2:", ch_data2)
      		default:
      			fmt.Println("No data in channel2 yet")
      		}
      	case ch_data2 := <-ch2:
      		fmt.Println("Received from channel2:", ch_data2)
      	default:
      		fmt.Println("No messages received yet from any channel")
      	}
      }
