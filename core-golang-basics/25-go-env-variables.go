/**
To set a key/value pair, use os.Setenv.
To get a value for a key use os.Getenv: this will return an empty string if the key isn't present in the env.
To list all key/value pairs os.Environ.
*/

package main
import (
  "fmt"
  "os"
  "strings"
  )
func main(){
  os.Setenv("Foo","1")
  fmt.Println("Foo", os.Getenv("Foo"))

  for _, e:= range os.Envion{
    pair := string.SplitN(e, "=", 2)
    fmt.Println(pair[0])
  }
}
