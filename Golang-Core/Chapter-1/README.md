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
