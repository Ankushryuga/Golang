package main
import ("fmt"
"io"
"os"
)

//1 Assertion to Non-concrete interface..
/**
 * you can use type assertion to assert that an interface{} value 
 * implements another interface, not just a concrete type.
 */
 
func assertNonConcrete(){
  var i interface{} = os.Stdin  //os.Stdin is of type *os.File.
  
  //Assert that i implements io.Reader(Non-concrete interface).
  
  r, ok := i.(io.Reader)
  if ok{
    fmt.Println("I implements io.Reader")
    readBuf := make([]byte, 10)
    n, _ := r.Read(readBuf) //safe to use Read.
    fmt.Printf("Read %d bytes\n", n)
  }else{
    fmt.Println("i does not implement io.Reader")
  }
}

//safe type assertion handle..
func handlePanicInAssertion(){
  var val interface{}="Hello"
  // num := val.(int)  //panic..

  //safe
  str, ok := val.(string)
  if ok{
    fmt.Println("String value: ", str)
  }else{
    fmt.Println("Not a string")
  }

  num, ok := val.(int)
  if ok{
    fmt.Println("Integer value:", num)
  }else{
    fmt.Println("Not an int")
  }
}


func main(){
  assertNonConcrete()
}
