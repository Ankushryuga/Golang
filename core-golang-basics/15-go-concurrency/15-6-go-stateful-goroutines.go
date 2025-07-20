/***
A stateful goroutine is a goroutine that encapsulates and manages its own internal state, often in response to external events or messages, usually through channels. 
Instead of sharing memory, the goroutine "owns" the state and updates it internally.

## 🧠 Why Use Them?
    => Stateful goroutines are used when:
      - You want to avoid locking (e.g., sync.Mutex) and shared memory.
      - You want to isolate state to reduce race conditions.
      - You want an actor-like model (somewhat inspired by Erlang or Akka).

📌 Benefits
No need for locks/mutexes.

The state is safely encapsulated.

Can be easier to reason about in complex concurrent systems.
*/

package main
import ("fmt")

type command struct{
  action      string
  value       int
  reply       chan int
}

func counter(cmds chan command){
  count :=0
  for cmd := range cmds{
    switch cmd.action{
      case "increment":
        count +=cmd.value
      case "get":
        cmd.reply <- count
    }
  }
}

func main(){
  cmds := make(chan command)
  go counter(cmds)

  //increment by 5.
  cmds <- command{action:"increment", value:5}
  // get current count.
  reply := make(chan int)
  cmd <- command{action:"get", reply:reply}
  fmt.Println("count", <-reply)

  //close when done.
  close(cmds)
}
