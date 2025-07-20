/****
panic is Go's built-in mechanism for handling unexpected runtime errors. when a panic occurs, it immediately stops execution of the current function and begins 
unwinding the stack, running any defer statements along the way. If the panic reaches the top of goroutine and isn't recovered, the program crashes.

## rules of thumb.
| Situation                    | Use `panic`?       |
| ---------------------------- | ------------------ |
| Logic bug / impossible state | ✅ Yes              |
| Input error or user behavior | ❌ No (use `error`) |
| Library misuse by caller     | ✅ Maybe            |
| General flow control         | ❌ Never            |

*/

package main
import "fmt"

func main(){
  //Recovering from panic
  defer func(){
    if r:=recover(); r!=nil{
      fmt.Println("Recovered from panic", r)
    }
  }()
  fmt.Println("About to panic")
  panic("crash")
  fmt.Println("This will not run")
}
