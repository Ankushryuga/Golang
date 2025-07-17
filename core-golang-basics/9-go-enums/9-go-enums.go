/***
Enumerated types (enums) are a distinct type that consists of a set of named constants. It improves type safety, readability and intention clarity.

In Go, there is no enum keyword - instead, we create enums using:
  - Named types(usually based on int)
  - const + iota

## Basic Enum pattern.
type Status int    // It defines a new named type based on int.
const (
  Pending Status = iota
  Approved
  Rejected
)  // defines constants
//iota : auto-increments starting from 0.
| Name     | Value |
| -------- | ----- |
| Pending  | 0     |
| Approved | 1     |
| Rejected | 2     |
*/

package main
import "fmt"

type Status int
const (
  Pending Status = iota
  Approved
  Rejected
  )
//using enums..
func process(status Status){
  if status==Approved{
    fmt.Println("Status: Approved")
  }
}

//Custom String() Method.
// Go's fmt.Println() won't print the enum name by default. You need to implement the String interface..
func (s Status) String() string{
  return [...]string{"Pending", "Approved", "Rejected"}[s]
}
//ensure the index are valid, to avoid panics..

//Enum Validation: You might want to validate if a value is part of the enum:

func isValidStatus(s Status) bool {
  switch s{
    case Pending, Approved, Rejected:
      return true
  }
  return false
}

//Custom Enum Values(Not Starting from 0): You can set specific values manually:
const (
  Low     Priority = 1
  Medium  Priority = 2
  High    Priority = 3
  ) // or use iota + offset:

const (
  One = iota +1 //1
  Two           // 2
  )

//Enums with JSON : If you are using enums with encoding/json. Its better to use strings for readability.

//Marshal to JSON:
func (s Status) MarshalJSON() ([]byte, error){
  return json.Marshal(s.String())
}
//Unmarshal from JSON:
func (s *Status) unmarshalJSON(data []byte) error{
  var str string
  if err := json.Unmarshal(data, &str); err!=nil{
    return err
  }

  switch str{
    case "Pending":
      *s = Pending
    case "Approved":
      *s = Approved
    case "Rejected":
      *s = Rejected
    default:
      return fmt.Errorf("Invalid status: %s", str)
  }
  return nil
}

// Enum Type Safety: You can define methods only applicable to your custom enum type..
func (s Status) IsFinal() bool{
    return s == Approved || s == Rejected
}

//Enum with flags(bit masking)  : If you want to represent combinations.
type Permission uint

const (
  Read Permission = 1<<iota  //1
  Write                      //2
  Execute                    //3
)

func Combine(){
  perm := Read | Write
  if perm&Read !=0{
    fmt.Println("Has read permission")
  }
}


//Enum Collections & Iteration: Iterate over enum values:
var allStatus = []Status{Pending, Approved, Rejected}

for _, s := range allStatus {
  fmt.Println(s)
}


//Common Pitfalls
/***
| Pitfall                            | Explanation                                                            |
| ---------------------------------- | ---------------------------------------------------------------------- |
| Using `int` instead of custom type | Loses type safety. Always define a custom type like `type Status int`. |
| Missing `String()` method          | Printing will show numbers unless `String()` is implemented.           |
| Out-of-bounds in `String()` array  | Panics if value isn't valid. Validate input or use map instead.        |
| Forgetting validation in JSON      | Malformed data can cause issues during unmarshaling.                   |

*/
