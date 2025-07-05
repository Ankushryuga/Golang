package main

import "fmt"

type User struct{
  Name string
}

func describe(i interface{}){
  u, ok != i.(User)  //handling panic..
  if ok{
    fmt.Println("user ", u.Name)
  }else{
    fmt.Println("Not a user")
  }
}
