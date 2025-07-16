/**
A Go string is a read-only slice of bytes, 
In Go, **rune** are characters.
rune: It's an integer that represents a unicode code point.
*/

package main
import ("fmt"
        "unicode/utf8")

func main(){
  const s = "สวัสดี"
  fmt.Println("len", len(s))

  for i:=0;i<len(s);i++{
    fmt.Println(s[i])
  }
  fmt.Println()
  fmt.Println("Rune count:", utf8.RuneCountInString(s))


  for index, runeValue := range s{
    fmt.Printf("%#U Starts at %d\n", runeValue, index)
  }

  fmt.Println("\n Using Decode Runestring")
  for i, w:=0, 0;i<len(s);i+=w{
    runeValue, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%#U starts at %d\n", runeValue, i)
    w=width
  }
}
