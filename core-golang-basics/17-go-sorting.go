/**
Go's sort package support sorting:
1. Primitive slices: []int, []string, []float64
2. custom structs using interfaces 


## Reverse Sorting.: use sort.Reverse
sort.Sort(sort.Reverse(ByAge(people)))
*/


package main
import (
  "cmp"
  "fmt"
  "slices"
  "sort"
  
  )

type Person struct{
  Name string
  Age int
}

type ByAge[]Person

func (a ByAge) Len()int    {return len(a)}
func (a ByAge) Swap(i, j int) {return a[i], a[j]=a[j], a[i]}
func (a ByAge) Less(i, j int) bool {return a[i].Age < a[j].Age}

func main(){
  people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 20}
  }
  

  sort.Sort(ByAge(people))
  fmt.Println("Sorted by age: ", people)
  

  //sorting by functions.
  fruits := []string{"peach", "apple", "banana"}
  lenCmp := func(a, b string) int{
    return cmd.Compare(len(a), len(b))
  }

  slices.SortFunc(fruits, lenCmp)
  fmt.Println(fruits)

  type peo struct{
    name     string
    age      int
  }

  peos := []peo{
    peo{name:"jax", age:20]
  }

    slices.SortFunc(people, func(a, b peo) int{return cmd.Compare(a.age, b.age)})
    fmt.Println(peos)
}


