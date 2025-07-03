package main

import "fmt"

func main(){
  //1: declare and initialize with values (composite literal):
    m := map[string]int {
      "apple":5,
      "banana":6,
    }

  //2: Declare and initialize empty map using:
    m2 := make(map[string]int)

  //3: with cap
  m3 := make(map[string]int, 100)  //preallocate space for 100 keys..

  //4: using var: just delcaring with var gives you a nil map, you must initialize it with make before writing.
  var m4 map[string]int
  m4 = make(map[string]int)

  //5: delcareas var and initialize with literal:
  var m5 = map[string]int{"a":304,"b":404}

  //6: nil map:
  var m6 map[string]int  //nil map
  fmt.Println(m6==nil)

  //7: map of slices or struct:
  team := map[string][]string{
    "devs":{"user1", "user2"},
    "test":{"test1", "test2"}
  }

  people := map[int]struct{
    Name string
    Age int
  }{
    1:{"Alice",23},
    2:{"Bob",40},
  }
  
}
