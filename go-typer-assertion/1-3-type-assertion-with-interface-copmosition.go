package main

import "fmt"

type Shape interface{
  Area() float64
}

type Circle struct{
  Radious float64
}


func (c Circle) Area() float64{
  return 3.14 * c.Radious * c.Radious
}
func DescribeShape(s Shape){
  c, ok := c.(Circle)  //type assertion::
  if ok{
    fmt.Println("Circle with radious: ", c.Radious)
  }
}
