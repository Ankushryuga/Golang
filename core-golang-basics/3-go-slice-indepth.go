/**
- In Go slices are designed to hold elements of a single, specific type. This means you cannot directly create a slice that contains elements of different types.
# However, there are workarounds to achieve similar functionality.

# An uninitialized slice equals to nil and has length 0.
# To create a slice with non-zero length, use the buildin make.
# By default a new slice's capacity is equal to its length;


# Slice support a "slice" operator with the syntax slice[low:high].

- 1. Using interface{}
- 2. Using struct{}
- 3. Map
*/
/***
- Using interface{} sacrifices type safety, so you need to use type assertions or type switch to work with the values.
- structs and map provide more structure but may require additional effor to manage.
*/

package main
import "fmt"

//# Slice support a "slice" operator with the syntax slice[low:high].
func sliceOperators(){
  var s[]string
  fmt.Println("uninit: ", s, s==nil, len(s)==0)

  s = make([]string, 3)
  fmt.Println("s: ", s, "len:", len(s), "caps:" cap(s))

  s[0]="a"
  s[1]="b"
  s[2]="c"

  fmt.Println("set: ", s)
  s.append(s, "d")
  s.append(s, "e")

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("copy:", c)

  l:=s[2:5]  //low:2, and high:5
  fmt.Println("slice1", l)

  l=s[:5]  //slices up to (but excluding)
  fmt.Println("slice2", l)

  l =s[2:]  //slice up from (and including)
  fmt.Println("slice3", l)

  twoD := make([][]int 3)
  for i:=0;i<3;i++{
    innerLen := i+1
    twoD[i]=make([]int, innerLen)
    for j:=0;j<innerLen;j++{
      twoD[i][j]=i+j  
    }
  }
  fmt.Println("2d", twoD)
}


//1-Using interfaces
func interfaceSlice(){
  mixedSlice := []interface{}{42, "hello", 3.14, true}
  for _, v:= range mixedSlice{
    fmt.Printf("Value: %v, Type: %T\n", v, v)
  }
}

//2 Using struct
type Mixed struct{
  IntVal      int
  StringVal   string
  FloatVal    float64
}

func structSlice(){
  mixedSlice := []Mixed{
    {IntVal: 42, StringVal: "Hello", FloatVal: 3.14}
    {IntVal: 43, StringVal: "World", FloatVal: 3.14223}
  }

  for _, v:= range mixedSlice{
    fmt.Println(v)
  }
}


//3 Using Map:
func mapSlice(){
  mixedSlice := []map[string]interface{}{
    {"id":1, "name":"Ankush", "Score":11.3},
    {"id":2, "name":"Banti", "Score":121.3},
    {"id":3, "name":"AN", "Score":11.3}
  }
  for _, v:=range mixedSlice{
    fmt.Println(v)
  }
}



