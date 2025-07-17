package main

import (
  "encoding/json"
  "errors"
  "fmt"
  )

//step1 : Define the custom type..
type Status int

//step2 : Define enum values using iota
const (
  Unkonwn Status = iota
  Pending
  Approved
  Rejected
)
//step 3: Provide string representations.
var statusNames = [...]string{
  "Unknown",
  "Pending",
  "Approved",
  "Rejected",
}

func (s Status) String() string{
  if int(s) <0 || int(s) >= len(statusNames){
    return "Invalid"
  }
  return statusNames[s]
}

//step 4: Marshal as string to JSON..
func (s Status) MarshalJSON() ([]byte, error){
  return json.Marshal(s.String())
}

//step 5: unmarshal string to enum..
func (s *Status) UnmarshalJSON(data []byte) error{
  var str string
  if err := json.Unmarshal(data, &str); err != nil{
    return err
  }

  for i, name := range statusNames{
    if name == str{
      *s = Status(i)
      return nil
    }
  }
  return errors.New("Invalid status: " + str)
}


// step 6: Enum validation
func IsValidStatus(s Status) bool{
  return s>Unknown && s<= Rejected
}

//step 7: Gel all enum values..
func AllStatuses() []Status{
  return []Status{Pending, Approved, Rejected}
}










  
