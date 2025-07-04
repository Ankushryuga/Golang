// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"
type Person struct{
    Name string
    Tag []string
}

func deepCopy(p Person) Person{
    copiedTags := make([]string, len(p.Tag))
    copy(copiedTags, p.Tag)

    return Person{
        Name: p.Name,
        Tag: copiedTags,
    }

}
func main() {
    
    p1:= Person{Name:"Ankur", Tag: []string{"g1", "g2"}}
    ps := p1    //shallow
    p2 := deepCopy(p1)
    fmt.Println(ps)
    fmt.Println(p2)
  fmt.Println("Try programiz.pro")
}
