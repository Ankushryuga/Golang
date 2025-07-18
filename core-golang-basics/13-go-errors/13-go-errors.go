/***
Go don't have exception handling, instead it have errors 
########
Go: the error type is just an interface..

type error interface{
  Error() string
}
Any type that implements the Error() string method satisfies the error interface.
## Go encourange you to handle errors immediately, rather than letting them propagate invisibly like exceptions.


errors.Is(err, target)    -- checks if an error is or wraps another.
errors.As(err, &target)   -- Checks if error is of a specific type.
## Error Wrapping::
use %w to wrap errors..

| Feature     | `error`         | `panic`                          |
| ----------- | --------------- | -------------------------------- |
| Use for     | Expected errors | Unexpected bugs (e.g. nil deref) |
| Recoverable | Yes             | Only with `recover()`            |
| Idiomatic   | ✅ Yes           | ❌ Avoid unless fatal             |



| Concept            | Description                                     |
| ------------------ | ----------------------------------------------- |
| `error` type       | Interface with `Error() string` method          |
| `errors.New`       | Basic error with message                        |
| `fmt.Errorf`       | Error with formatting and wrapping (`%w`)       |
| `errors.Is`        | Check for known error value (including wrapped) |
| `errors.As`        | Check for specific error types                  |
| Custom error types | Add context (e.g. error codes, messages)        |
| Wrapping           | Add context while preserving original error     |

*/
package main
import ("errors"
        "fmt"
        )
//Creating errors..
var errNotFound = errors.New("Item not found")

func find(id int) error{
  return errNotFound
}

func divide(a, b int) (int, error){
  if b==0{
    return 0, fmt.Errorf("Cannot divide by zero")
  }
  return a/b, nil
}


//Handling errors:::
func handleErrors() error {
  res, err := doSomething()
  if err!=nil{
    return err
  }
  return nil
}

//Custom error types...
//create structured, typed errors...
type MyError struct{
  Code int
  Msg  string
}

func (e *MyError) Error() string{
  return fmt.Sprintf("Code %d: %s", e.Code, e.Msg)
}


//Retrying on errors..
func retryErrors(){
  for i:=0;i<3;i++{
    err:=doNetworkCall()
    if err == nil{
      break
    }
    time.Sleep(time.Second)
  }
}


///Example 
func readConfig(path string) (string, error){
  data, err := os.ReadFile(path)
  if err!=nil{
    return "", fmt.Errorf("read config failed: %w", err)
  }
  return string(data), nil
}

func main(){
  _, err := readConfig("missing.conf")
  if err!=nil{
    if errors.Is(err, os.ErrNotExist){
      fmt.Println("config file not found..")
    }else{
      fmt.Println("Error: ", err)
    }
  }
}

