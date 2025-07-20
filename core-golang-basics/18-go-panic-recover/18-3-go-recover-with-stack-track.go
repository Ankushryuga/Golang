/**
You can use the runtime/debug package to log the full stack trace during a panic.
| Pattern              | Use Case                            |
| -------------------- | ----------------------------------- |
| `recover()`          | Stop panic and continue             |
| Middleware + recover | Prevent HTTP server crash           |
| `debug.Stack()`      | Print full call trace               |
| Logging + re-panic   | Trace, then allow crash (if needed) |

*/

package main
import (
  "log"
  "runtime/debug"
  )

func RecoverWithStack(){
  if r:=recover(); r!=nil{
    log.Println("[Panic]", r)
    log.Println(string(debug.stack()))  //full stack trace..
  }
}


func main(){
  defer RecoverWithStack()
  fmt.Println("Starting app..")
  panic("critical failure")
}
