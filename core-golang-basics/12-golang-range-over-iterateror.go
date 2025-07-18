/****
Go, range is a keyword used to iterate over the elements of
- Arrays and Slices
- Maps
- Strings
- Channels

| Use Case           | Can Use `range` | Requires Goroutine | Comment                       |
| ------------------ | --------------- | ------------------ | ----------------------------- |
| Slice/Array        | ✅ Yes           | ❌ No               | Native support                |
| Map                | ✅ Yes           | ❌ No               | Order not guaranteed          |
| String             | ✅ Yes           | ❌ No               | Iterates over Unicode runes   |
| Channel            | ✅ Yes           | ✅ Often            | Great for streaming values    |
| Custom Iterator    | ❌ No (directly) | ❌ No (direct)      | Must use `Next()` manually    |
| Iterator → Channel | ✅ Yes           | ✅ Yes              | Best for custom `range` usage |

*/

package main

//range over slice:
func rangeOverSlice(){
  num := []int{1,2,3,4}
  for i, v:= range num{
    fmt.Println(i, v)
  }
}

//over a map
func rangeOverMap(){
  nums := map[int]string{
    1:"Value1",
    2:"Value2",
  }
  for k, v:= range nums{
    fmt.Println(k, v)
  }
}

// over a channel
func rangeOverChannel(){
  ch := make(chan int)
  go func(){
    for i:=0;i<5;i++{
      ch <- i  //sending
    }
    close(ch)  //closing channel 
  }()
  //ranging over channel.
  for v := range ch{
    fmt.Println(v)
  }
  }


//range over a custom iterator via channel..
func Fibonacci(n int)  <-chan int{
  ch := make(chan int)
  go func(){
    a, b:=0,1
    for i:=0;i<n;i++{
      ch<-i
      a, b=b, a+b
    }
    close(ch)
  }()
  return ch
}


//Manual Iterator pattern:
type Iterator struct{
  data   []int
  index  int
}

func (it *Iterator) Next() (int, bool){
  if it.index >= len(it.data){
    return 0, false
  }
  val := it.data[it.index]
  it.index++
  return val, true
}





func main(){
  for num := range Fibonacci(5){
    fmt.Println(num)
  }
  //using manual iterator pattern..
  it := &Iterator{data:[]int{10,20,30}}

  for {
    val, ok := it.Next()
    if !ok{
      break
    }
    fmt.Println(val)
  }
}

