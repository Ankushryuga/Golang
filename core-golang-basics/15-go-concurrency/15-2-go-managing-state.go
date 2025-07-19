/**
Managing state in go involves maintaining and coordinating shared data across parts of a program.
## Choose right:
| Scenario                      | Use                                |
| ----------------------------- | ---------------------------------- |
| Single-threaded state         | Structs or methods                 |
| Shared between goroutines     | `sync.Mutex` / `sync.RWMutex`      |
| High-speed numeric counters   | `sync/atomic`                      |
| Fully serialized state access | Channels + goroutine (actor model) |
| Temporary read-only state     | Copy the data (safe in Go)         |


*/

package main
import "fmt"

///1 Basic state with structs:
type counter struct{
  value int
}

///2 State with methods(Encapsulations) : Use methods to encapsulate acces to the internal state.
func (c *Counter)Increment(){
  c.value++
}
func (c *Counter)Decrement(){
  c.value--
}
func (c *Counter) Value()int{
  return c.value
}

//3. State with Mutex (Safe sharing across goroutines): Use sync.Mutex or sync.RWMutex to protect state wheh shared b/w goroutines.
type SafeCounter struct{
  mu sync.Mutex
  value int
}
func (c *SafeCounter) Increment(){
  c.mu.Lock()
  defer c.mu.Unlock()
  c.value++
}

func (c *SafeCounter) Value()int{
  c.mu.Lock()
  defer c.mu.Unlock()
  return c.value
}

//4. State via channels(Message passing): Instead of sharing data with locks, you can shared ownership by funneling all state access through a goroutine and using channels.
//this approch avoids race conditions without needing locks.
type command struct{
  action string
  value int
  reply chan int
}

func stateManager(cmds <-chan command){
  var counter int
  for cmd := range cmds{
    switch cmd.action{
      case "inc":
        counter+=cmd.val
      case "get":
        cmd.reply<-counter
    }
  }
}


//5. State with sync/atomic(Fast counters): when performance matters and you're only modifying simple numeric value, use atomic operations.:

func AtomicStateExample(){
  var count int64
  atomic.AddInt64(&count, 1)
  val := atomic.LoadInt64(&count)

  
}

func main(){
  //1 basic state with structs
  c := Counter{}
  c.value++
  fmt.Println(c.Value)

  //4.State via channels.
  cmds := make(chan command)
  go stateManager(cmds)
  cmds <- command{action: "inc", val:1}
  reply := make(chan int)
  cmds <- command{action: "get", reply:reply}
  fmt.Println("value: ", <-reply)
}
