/***
Generics: It allow you to write functions, types, and data structures that can operate on value of any type in a type-safe way.
They enable parametric polymorphism — meaning behavior is defined generically over a type parameter, without sacrificing compile-time type safety.

| Tip                       | Recommendation                                                       |
| ------------------------- | -------------------------------------------------------------------- |
| Start small               | Use generics where they reduce duplication                           |
| Use meaningful type names | Use `T`, `K`, `V` for short examples; descriptive names in real code |
| Avoid overengineering     | Don't genericize just for the sake of it                             |
| Combine with interfaces   | Use method-based constraints to express capability                   |


*/
package main
import "fmt"

//Basic Syntax:
func Print[T any](value T){
  fmt.Println(value)
}
/***
- T : is a type parameter
- any: is a built-in alias for interface{}
*/

/**
# Type Parameters in structs and types..
*/

type Box[T any] struct{
  value T
}

func (b Box[T]) Get() T{
  return b.value
}

func (b *Box[T]) Set(v T){
  b.value=v
}


//constraints: It limit what types can be passed as type parameters. They are specified using interfaces.
func sum[T int | float64] (a, b T) T{
  return a+b
}
//T  must be eitehr int or float64.


type Number interface {
    int | int64 | float64
}

func Add[T Number](a, b T) T {
    return a + b
}

//Using Custom Constraints

type Stringer interface {
    String() string
}

func PrintString[T Stringer](val T) {
    fmt.Println(val.String())
}


// Type Set:: A type set defines the set of types a type parameter can be. It's used in constraints.

type Ordered interface {
    ~int | ~float64 | ~string
}
/**
- ~T means all types whose underlying type is T.
- Useful for working with custom types based on primitives.
*/


//Generic methods:You cannot define generic methods on non-generic types yet.

type Pair[T any] struct {
    First, Second T
}

func (p Pair[T]) Swap() Pair[T] {
    return Pair[T]{First: p.Second, Second: p.First}
}



//uses
func main(){
  inbox:=Box[int]{value:100}
  stringBox := Box[string]{value:"hello"}
}
