// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func flatten(input interface{})[]interface{}{
    var result []interface{}
    
    switch v := input.(type){
        case []interface{}:
            for _, item := range v{
                result = append(result, flatten(item)...)
            }
            default:
            result = append(result, v)
    }
    return result
}
func main() {
    
    nested := []interface{}{
        1,
        2,
        []interface{}{2,3},
    }
  fmt.Println("Try programiz.pro")
  
  fmt.Println(nested)
  
  fmt.Println(flatten(nested))
}
