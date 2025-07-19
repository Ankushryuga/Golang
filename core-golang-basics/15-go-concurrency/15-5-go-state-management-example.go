/***
 concurrent in-memory counter service using a goroutine + channels pattern (aka actor model). This is one of the safest and most idiomatic ways to manage shared state in Go without using mutexes
 | Concept               | What It Does                                            |
| --------------------- | ------------------------------------------------------- |
| `CounterCommand`      | Encapsulates all possible state-changing actions        |
| `counterManager`      | Owns the counter state, processes commands sequentially |
| `chan CounterCommand` | Acts as a "message queue" to interact with state        |
| `chan int` (reply)    | Used for sending results back safely                    |


 
 */

package main
import ("fmt"
        "time"
        )


//Define the command types the state manager will accept
type CounterCommand struct{
  Action       string      //"inc", "get", "reset"
  Amount       int        //for "inc"
  Reply        chan int    // for "get"
}

func counterManger(cmds <-chan CounterCommand){
  counter :=0 
  for cmd := range cmds{
    switch cmd.Action{
      case "inc":
        counter += cmd.Amount
      case "get":
        cmd.Reply <- counter
      case "reset":
        counter =0
    }
  }
}

func main(){
  commands := make(chan CounterCommand)

  //start the state manager in a gorountine.
  go counterManager(commands)

  //simulate concurrent worker 
  for i:=0;i<5;i++{
    go func(id int){
      for j:=0;j<4;j++{
        commands <- CounterCommand{Action:"inc", Amount:1}
        time.Sleep(100*time.Millisecond)
      }
    }(i)
  }


  //Give some time for workers to finish.
  time.Sleep(1*time.Second)
  //request the current counter
  reply := make(chan int)
  commands <- CounterCommand{Action:"get", Reply:reply}
  fmt.Println("Final Counter value:", <-reply)

  //Rest counter.
  commands <- CounterCommand{Action:"reset"}

  //verify reset.
  commands <- CounterCommand{Action:"get", Reply:reply}
  fmt.Println("Counter After Reset", <-reply)
}







