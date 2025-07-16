/***
To create an empty map, use the builtin make:
make(map[key-type]value-type)

# Set key/value pairs using typical name[key]=val syntax.
# Get a value for a key ith name[key].


Key Types
Keys must be comparable types:

Allowed:
  - Basic types: int, float64, string, bool
  - Pointers
  - Structs (if all fields are comparable)
  - Arrays (elements must be comparable)
Not allowed:
  - Slices
  - Maps
  - Functions

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



import (
  "fmt"
  "sort"
  "sync"
)

func main() {
  // Create map
  capitals := map[string]string{
    "USA":    "Washington",
    "France": "Paris",
    "Japan":  "Tokyo",
  }

  // Insert/update
  capitals["Germany"] = "Berlin"
  capitals["France"] = "Lyon" // update value

  // Delete
  delete(capitals, "Japan")

  // Existence check
  if capital, ok := capitals["USA"]; ok {
    fmt.Println("Capital of USA:", capital)
  }

  // Iterate in sorted order
  keys := make([]string, 0, len(capitals))
  for k := range capitals {
    keys = append(keys, k)
  }
  sort.Strings(keys)
  for _, k := range keys {
    fmt.Printf("%s -> %s\n", k, capitals[k])
  }

  // Map as a set
  set := make(map[string]struct{})
  set["apple"] = struct{}{}
  if _, exists := set["apple"]; exists {
    fmt.Println("apple is in the set")
  }

  // Concurrent safe map with sync.Map
  var sm sync.Map
  sm.Store("foo", 42)
  if val, ok := sm.Load("foo"); ok {
    fmt.Println("sync.Map foo:", val)
  }
}

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

  //Maps declaration & Initialization..
  var mapp map[string]int       // nil map, unusable until initialized
  mapp = make(map[string]int)   // initialized empty map
  
  // or use shorthand:
  mapp := make(map[string]int)
  
  // Or initialize with literals
  mapp := map[string]int{"a": 1, "b": 2}
  //Nil maps behave like empty maps on lookups but panic on writes.
  //Always use make or literals before adding keys.
  
}
