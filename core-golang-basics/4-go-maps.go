/***
To create an empty map, use the builtin make:
make(map[key-type]value-type)

# Set key/value pairs using typical name[key]=val syntax.
# Get a value for a key ith name[key].

# If the key doesn't exist, the zero value of the value type is returned.
# builtin len: returns the number of key/value pairs whn called on a map.
# builtin delete removes key/value pairs from a map.
# To remove all key/value pairs from a map, use the clear builtin.

## Invalid Key Types in Golang.
  - Slice
  - Maps
  - Functions

### NOTE:::::: Maps are not safe for concurrent writes.
# Use: sync.Map or a mutex if accessing maps from multiple goroutines.

✅ Best Practices
 - Use make() when building maps dynamically.
 - Always check if a key exists before using it.
 - Don’t rely on iteration order.
 - Avoid using maps with very large values as keys (due to performance).
*/

package main
import "fmt"

func main(){
  m:=make(map[string]int)
  m["k1"]=10
  m["k2"]=11
  m["k3"]=12
  m["k4"]=13

  m1 := map[string]int{"a":100}

  //Nested Map..
  m2 := map[string]map[string]int{
    "Alice":{
      "math":100,
      "science":100
    },
  }

  fmt.Println(m2["Alice"]["math"])
  
}
