# Chapter 1 of golang Core

- [Chapter 1 of golang Core](#chapter-1-of-golang-core)
  - [Variables and Data Types](#variables-and-data-types)
    - [Variables](#variables)
  - [Constants](#constants)
  - [Data Types](#data-types)
    - [String](#string)
    - [Bool](#bool)
      - [Operators](#operators)
    - [Numeric types](#numeric-types)
      - [Signed and Unsigned integers](#signed-and-unsigned-integers)
      - [Byte and Rune](#byte-and-rune)
      - [Floating point](#floating-point)
      - [Operators](#operators-1)
      - [Complex](#complex)
    - [Zero values](#zero-values)
    - [Type Conversion](#type-conversion)
    - [Alias types](#alias-types)
    - [Defined types](#defined-types)
  - [String Formatting](#string-formatting)
  - [Flow control](#flow-control)
    - [If/Else](#ifelse)
    - [Compact if or short if](#compact-if-or-short-if)
    - [Switch](#switch)
    - [Loop](#loop)
      - [For loop](#for-loop)
      - [Break and continue](#break-and-continue)
      - [Forever loop](#forever-loop)
  - [Functions](#functions)
    - [Simple declaration](#simple-declaration)
    - [Returning the value](#returning-the-value)
    - [Multiple returns](#multiple-returns)
    - [Named return values](#named-return-values)

## Variables and Data Types

### Variables

- Declaration without initialization:

```go
var name string
```

- Declaration with initialization:

```go
var name string = "Test""
```

- Multiple declaration:

```go
var var1, var2 string = "Hello", "world""
// OR
var (
    var1 string = "Hello""
    var2 string = "World"
)
```

- Type is omitted but will be inferred:

```go
var var1 = "what is type of this variable?"
```

- Here we omit `var` keyword and type is always implicit. This is how we will see variables being declared most of the time.
- We also use `:=` for declaration plus assignment.

```go
foo := "Shorthand!"
```

> [!NOTE]
> Shorthand only works inside `function` bodies

## Constants

Constant can be declare with the `const` keyword.

```go
const constant = "This is a constant" // value cannot be reassigned.
```

> [!NOTE]
> Only constants can be assigned to other constants.

```go
const a = 10
const b = a // ✅ Works

var a = 10
const b = a // ❌ a (variable of type int) is not constant (InvalidConstInit)
```

## Data Types

### String

- String = It's a sequence of bytes. They are declared either using double quotes or backticks which can span multiple lines.

```go
var name string = "My name is Go""

var bio string = `This is golang string type.
                Its a sequence of bytes.`
```

### Bool

- `bool` which is used to store boolean values. It can have two possible values - `true` or `false`.

```go
var value bool = false
var isItTrue bool = true
```

#### Operators

- we can use following operators on boolean types
  - Logical = `&&`, `||`, `!`
  - Equality = `==`, `!=`

### Numeric types

#### Signed and Unsigned integers

- Go has several built-in integer types of varying sizes for storing signed and unsigned integers.

- The size of generic `int` and `uint` types are platform-dependent. This means it is 32-bits wide on a 32-bit system and 64-bits wide on a 64-bit system.

```go
var i int = 404                     // Platform dependent
var i8 int8 = 127                   // -128 to 127
var i16 int16 = 32767               // -2^15 to 2^15 - 1
var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1
```

- Similar to signed integers, we have unsigned integers.

```go
var ui uint = 404                     // Platform dependent
var ui8 uint8 = 255                   // 0 to 255
var ui16 uint16 = 65535               // 0 to 2^16
var ui32 uint32 = 2147483647          // 0 to 2^32
var ui64 uint64 = 9223372036854775807 // 0 to 2^64

var uiptr uintptr                     // Integer representation of a memory address
```

> [!NOTE]
> There's also an unsigned integer pointer `uintptr` type, which is an integer representation of a memory address. It is not recommended to use this.

> [!IMPORTANT]
> It is recommended that whenever we need an integer value, we should just use `int` unless we have a specific reason to use a `sized` or `unsigned` integer type.

#### Byte and Rune

- Golang has two additional types called `byte` and `rune` that are aliases for `uint8` and `int32` data types respectivly.

```go
type byte = uint8
type rune = int32
```

- A `rune` represents a unicode code point.

```go
var b byte = 'a'
var r rune = '🍕'
```

#### Floating point

- These are used to store numbers with a decimal component.
- Go has 2 floating point types `float32` and `float64`. Both type follows the IEEE-754 standard.

> [!NOTE]
> The default type for floating point values is float64.

```go
var f32 float32 = 1.7812 // IEEE-754 32-bit
var f64 float64 = 3.1415 // IEEE-754 64-bit
```

#### Operators

- Go providers several operators for performing operations on numeric types.

- **Arithmetic**: `+`, `-`, `*`, `/`, `%`
- **Comparison**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Bitwise**: `&`, `|`, `ˆ`, `<<`, `>>`
- **Incre/Decre**: `++`, `--`
- **Assignment**: `=`, `+=`, `-=`, `*=`, `/=`, `%=`, `<<=`, `>>=`, `&=`, `|=`, `ˆ=`

#### Complex

There are 2 complex types in Go, `complex128` where both real and imaginary parts are `float64` and `complex64` where real and imaginary parts are `float32`.

we can define complex numbers either using the built-in complex function or as literals.

```go
var c1 complex128 = complex(10, 1)
var c2 complex64 = 12 + 4i
```

### Zero values

- In go, any variable declared without an explicit initial value is given its zero value.

```go
var i int
var f float64
var b bool
var s string

fmt.Printf("%v %v %v %q\n", i, f, b, s)
```

- `int` and `float`are assigned as `0`, `bool` as false, and `string` as an `empty string`.

### Type Conversion

```go
i := 32
f := float64(i)
u := uint(f)

fmt.Printf("%T %T", f, u)
```

> [!NOTE]
> This is different from parsing

### Alias types

- It allow you to interchangeably with the underlying type.

```go
package main
import "fmt"

type MyAlias = string
func main(){
    var str MyAlias = "I am an alias"
    fmt.Printf("%T - %s", str, str) // Output: string - I am an alias.
}
```

### Defined types

- Unlike `alias` types do not use an equals sign.

```go
package main
import "fmt"

type MyDefined string
func main() {
    var str MyDefined = "I am defined"
    fmt.Printf("%T -%s", str, str)  // Output: main.MyDefined - I am defined.
}
```

> [!NOTE]
> `defined types` do more than just give a name to a type.
> It first defines a new named type with an underlying type. However, this defined type is different from any other type. including its underlying type.
>
> Hence, It cannot be used interchangeably with the underlying type like alias types.

```go
// Example of defined type.
package main
import "fmt"

type MyAlias = string
type MyDefined string

func main(){
    var alias MyAlias
    var def MyDefined

    // ✅ Works
    var copy1 string = alias

    // ❌ Cannot use def (variable of type MyDefined) as string value in variable
    var copy2 string = def
    fmt.Println(copy1, copy2)
}

// as we can see, we cannot use the defined type interchangely with the underlying type, unlike alias types.
```

## String Formatting

- String formating or sometimes also known as templating.
- `fmt` package contains lots of functions. So to save time, we will discuss the most frequently used functions.

```go
// fmt.Print=> "Print" does not format anything, it simply takes a string and prints it.
fmt.Print("What", "is", "your", "name?")
fmt.Print("My", "name", "is", "golang")

// Println => It same as "Print", but it adds a new line at the end and also inserts space b/w the arguments.
fmt.Println("What", "is", "your", "name?")
fmt.Println("My", "name", "is", "golang")

// Printf => also knonw as "Print Formatter", which allows us to format numbers, strings, booleans and much more.
name:= "golang"
fmt.Println("What is your name?")
fmt.Printf("My name is %s", name)

// "%s" was substituted with name variable
// "%s", its called annotation verbs and they tell the function how to format the arguments. we can control things like width, types and precision with these and there are lots of them

percent := (7.0 / 9) * 100
fmt.Printf("%f", percent)       // o/p: 77.777778

// lets say we want just 77.78 which is 2 points precision, we can do that as well by using ".2f". Also, to add an actual percent sign, we will need to escape it.

fmt.Printf("%.2f %%", percent)
```

- `Sprint`, `Sprintln`, and `Sprintf`, these are basically the same as the print functions, the only difference being they can return the string instead of printing it.

```go
// example
s := fmt.Sprintf("hex:%x bin:%b", 10, 10)
fmt.Println(s)

// Sprintf formats our integer as hex or binary and returns it as a string.

// multiline string literals
msg := `
Hello from
multiline
`

fmt.Println(msg)
```

> [!NOTE]
> These String formatting examples are just tip so please checkout golang string formating [https://pkg.go.dev/fmt](https://pkg.go.dev/fmt)

## Flow control

### If/Else

- Works same as other programming languages, but the expression doesn't need to be surrounded by parentheses `()`.

```go
func main(){
  x := 10
  if x>5 {
    fmt.Println("x is gt 5")
  }else if x>10 {
    fmt.Println("x is gt 10")
  }else {
    fmt.Println("else case")
  }
}
```

### Compact if or short if

- we can also compact our if statements.

> [!NOTE]
> Go doen't have a special "short if" syntax or ternary operator (`? :`) like others, the idiomatic way to write concise unconditionals is to use the standard `if` statement with an optional **short statement** preceding the condition.

This format is commonly used for error handling. allowing a variable declaration (using `:=`) and a condition check in a single line. The variable's scope is limited to the `if` and corresponding `else` blocks.

```go
// syntax:
/**
if initialization; condition {
  // code block if condition is true
}else{
  // code block if condition is false
}
*/

func main(){
  if x:=10; x>5{
    fmt.Println("x is gt 5")
  }
}
```

```go
// error handling
import (
  "fmt"
  "os"
)

func main(){
  // The short statement opens a file and assigns the file handle and potential error.
  if file, err := os.open("example.txt"); err != nil {
    fmt.Println("Error opening file", err)
    // err is in scope here
  }else{
    fmt.Println("File opened successfully")
    file.close()
    // file and err are in scope here
  }
  // file and err are NOT in scope here and cannot be used.
}
```

### Switch

- In Go, the switch case only runs the first case whose value is equal to the condition expression and not all the cases that follow. Hence, unlike other languages, `break` statement is automatically added at the end of each case.

- It evaluates cases from top to bottom, stopping when a case succeeds.

```go
func main(){
  day:="monday"

  switch day{
    case "monday":
      fmt.Println("time to work!")
    case "friday":
      fmt.Println("let's party")
    default:
      fmt.Println("browse memes")
  }
}
```

- Go also support shorthanded declaration like this.

```go
switch day := "monday"; day{
  case "monday":
    fmt.Println("time to work!")
  case "friday":
    fmt.Println("let's party")
  default:
    fmt.Println("browser memes")
}
```

> [IMPORTANT]
>
> - `fallthrough` keyword to transfer control to the next case even though the current case might have matched.

```go
switch day:= "monday"; day{
  case "monday":
    fmt.Println("time to work!")
    fallthrough
  case "friday":
    fmt.Println("let's party")
  default:
    fmt.Println("browse memes")
}
// if we run this, we'll see that after the first case matches the switch statement continues to the next case because of the `fallthrough`keyword.
// o/p: time to work!
// o/p: let's party
```

- we can also use it without any condition, which is same as `switch true`

```go
x := 10
switch {
  case x > 5:
    fmt.Println("x is greater")
  default:
    fmt.Println("x is not greater")
}
```

### Loop

- Go, have only one type of loop which is `for` loop. And it doesn't need any parenthesis `()` unlike others.

#### For loop

- `init statement`: which is executed before the first iteration.
- `condition expression`: which is evaluated before every iteration.
- `post statement`: which is executed at the end of every iteration.

```go
func main(){
  for i:=0; i<10; i++{
    fmt.Println(i)
  }
}
```

#### Break and continue

- Go also supports both `break` and `continue` statements for loop control.

```go
func main(){
  for i:=0; i<10; i++{
    if i<2 {
      continue
    }
    fmt.Println(i)
    if i > 5{
      break
    }
  }
  fmt.Println("we broke out!")
}
```

So, the `continue` statement is used when we want to skip the remaining portion of the loo, and `break` statement is used when we want to break out of the loop.

> [!NOTE]
> `Init` and `post` statements are optional, hence we can make out `for` loop behave like a while loop as well.

```go
func main(){
  i:=0
  for ;i < 10;{ // we can also remove the additional semi-colons to make it a little cleaner.
    i += 1
  }
}
```

#### Forever loop

- If we omit the loop condition, it loops forever, so an infinite loop can be compactly expressed. This is also known as the forever loop.

```go
func main(){
  for {
    // do stuff here
  }
}
```

## Functions

### Simple declaration

```go
 `func myFunction() {}`
```

- And we can call or execute it as follows.

```go
myFunction()
```

```go
func main(){
  myFunction("Hello")
}
func myFunction(p1 string){
  fmt.Println(p1)
}
```

- when consecutive parameters have the same type.

```go
func myNextFunction(p1, p2 string) {}
```

### Returning the value

```go
func main(){
  s := myFunction("Hello")
  fmt.Println(s)
}

func myFunction(p1 string) string {
  msg := fmt.Sprintf("%s function", p1)
  return msg
}
```

### Multiple returns

```go
func main(){
  s, i := myFunction("Hello")
  fmt.Println(s, i)
}

func myFunction(p1 string) (string, int) {
  msg := fmt.Sprintf("%s function", p1)
  return msg, 10
}
```

### Named return values

- Go's return value may be named. If so, they are treated as variables defined at the top of the function.

- These names should be used to document the meaning of the return values.

- A `return` statement without arguments the named return values.

> [!NOTE]
> Naked return statements should be used only in short functions.

```go
package main
import "fmt"

func split(sum int)(x, y int){
  x = sum * 4 / 9
  y = sum - x
  return
}

func main(){
  fmt.Println(split(17))
}
```
