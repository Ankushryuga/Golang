# Constants (Can't be modified) in Go:
    => it refer to fixed values that the program may not alter during its execution, These fixed values are also called literals.
      In Go language, the constants can be created using the const keyword.

# Constants can be of any of the basic data types such as:
    =>
      Integer constant
      Floating constant
      Character constant
      String literal

Example:
  func main(){
    const LENGTH int = 10
    const WIDTH int = 5
    var area int
    area = LENGTH * WIDTH;

    fmt.Printf("value of area : %d", area)
  }


