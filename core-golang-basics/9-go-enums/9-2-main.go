package main
import (
  "encoding/json"
  "fmt"
  )

func main(){
  //Enum usage
  var current Status = Approved

  fmt.Println("Current Status: ", current)  //o/p: Approved
  fmt.Println("Is valid? ", IsValidStatus(current))  //o/p: true

  //Loop through all..
  fmt.Println("\nAll Status: ")
  for _, s:= range AllStatuses(){
    fmt.Println("-", s)
  }

  //Json example::
  data, _ := json.Marshal(current)
  fmt.Println("\nJSON encoded: ", string(data))

  //JSON decoded.
  var decoded Status
  _ = json.Unmarshal([]byte(`"Pending"`), &decoded)
  fmt.Println("Decoded from JSON:" decoded)                    //O/P: Pending
}
