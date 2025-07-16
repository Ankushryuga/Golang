/**
# in go, a function are created with func keyword.
## NOTE: Methods are not functions.
- Function: Independent block of code.
- Method: A function with receiver (associated with a type).

## init() and main() functions.
1. main() function
    - Entry point of a go program.
    - Must be in package main.

2. init() function:
    - Runs before main(), used for setup.
    - Each file can have multiple init() functions.


## Function Internals 
  - Functions are implemented as function pointers with closures in memory.
  - variadic parameters are passed as slices.
  - Anonymous functions capture variables by reference, not value.

## Higher-order function:
  - Function that take or return other functions:

# 🔹 15. Best Practices
| Practice                       | Recommendation                          |
| ------------------------------ | --------------------------------------- |
| Keep functions short & focused | Improves readability and reusability    |
| Use named returns for clarity  | Only when names add value               |
| Avoid deep recursion           | Go doesn't optimize tail recursion      |
| Use closures carefully         | Don’t overuse; can lead to memory leaks |
| Handle errors explicitly       | Go prefers simple error handling        |
| Use `defer` for cleanup logic  | For files, locks, DB connections, etc.  |

# 🔹 16. Common Mistakes to Avoid
| Mistake                                       | Fix                                             |
| --------------------------------------------- | ----------------------------------------------- |
| Not initializing slices in variadic functions | Use `make([]T, 0)` if needed                    |
| Using a loop variable in closures incorrectly | Copy loop variable inside loop                  |
| Forgetting `return` in named-return functions | `return` must still be called                   |
| Assuming `defer` runs after panic             | `defer` runs **during** panic propagation       |
| Passing large structs by value                | Use pointer receivers or pointers in parameters |


# Summary Table:
| Feature            | Description                               |
| ------------------ | ----------------------------------------- |
| Basic function     | `func name(args...) returnType { ... }`   |
| Multiple returns   | `(a int, b string)`                       |
| Variadic           | `func(args ...int)`                       |
| Anonymous function | `func() { ... }()`                        |
| Closures           | Inner functions capturing outer variables |
| Function as value  | `var f func(int) int`                     |
| Methods            | `func (r Receiver) Method()`              |
| Error handling     | `defer`, `panic`, `recover`               |
| Recursion          | Functions calling themselves              |
*/

package main
import "fmt"

//1. single return example
func singleReturnExample(name string)string{}

//2. multiple return example.
func divide(a, b int) (int, int){}

//3. Named Return values.
func area(length, width int) (area int){
  area=length*width
  return//implicit return of named variable.
}
//4. Variadic functions: Allows passing any number of arguments of the same type.
func sum(nums ...int) int{}  //sum(1,2), sum(1,2,3,4,5) etc..

//5. Anonymous functions(Lambdas).
//func (){}()
//Assigned to a variable..
greet := func(name string){}

//6. Closures: A function defined inside another that captures variable from the outer function.
func counter() func() int{
  count:=0
  return func() int{
    count++
    return count
  }
}

//c := counter()
//fmt.Println(c())  //1
//fmt.Println(c()) //2

//Defer, Panic, Recover with functions.
//Go functions can use defer to delay a statement until the function end.

func deferExample(){
  defer fmt.Println("This runs last")
  fmt.Println("This runs first")
}


/7. Function Types and Passing Functions.
func apply(op func(int, int) int, a, b int) int{return op(a, b)}
func add(x, y int) int{return x+y}
apply(add, 3,5)


//########### Methods::
type Rectangle struct{
  width, height int
}
//Method with value receiver..
func (r Rectangle) Area() int{
  return r.width*r.height
}
//Method with pointer receiver..
func (r *Rectangle) Scale(factor float64){
  c.width *= factor
}

///############################# VVI Recvoer: Catches a panic inside a deferred function..
func safeDivide(a, b int){
  defer func(){
    if r:=recover(); r!=nil{
      fmt.Println("Recovered from: ", r)
    }
  }()
  fmt.Println(a/b)
}



func greet(name string) string {
    return "Hello, " + name
}

func divide(a, b int) (q, r int) {
    q = a / b
    r = a % b
    return
}

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(greet("Go"))
    
    q, r := divide(10, 3)
    fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
    
    fmt.Println("Sum:", sum(1, 2, 3, 4))

    // Anonymous function
    square := func(x int) int {
        return x * x
    }
    fmt.Println("Square:", square(5))

    // Closure
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }()
    fmt.Println("Counter:", counter())
    fmt.Println("Counter:", counter())
}
