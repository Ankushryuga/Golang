package main

import "fmt"

func handleDynamicInput(data interface{}){
  //safe assertion:
  if userId, ok := data.(int); ok{
    fmt.Println("User id: ", userId)
  }else{
    fmt.Println("Invalid type for user id")
  }
}


/**
✅ Final Rule of Thumb:
❌ Never do v := x.(T) unless you're 100% sure what type x holds
✅ Always do v, ok := x.(T) if the type might vary
*/
