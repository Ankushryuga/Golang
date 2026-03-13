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
    - [Functions as values](#functions-as-values)
    - [Closures](#closures)
    - [Variadic Functions](#variadic-functions)
    - [Init](#init)
    - [Defer](#defer)
  - [Modules](#modules)
    - [What are modules?](#what-are-modules)
    - [Vendoring](#vendoring)
  - [Packages](#packages)
    - [What are packages?](#what-are-packages)
    - [External dependencies](#external-dependencies)
  - [Workspaces](#workspaces)
  - [Useful Commands](#useful-commands)
  - [Build](#build)

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
  return  // return statement without any arguments, this is also known as **naked returns**
}

func main(){
  fmt.Println(split(17))
}
```

### Functions as values

- In Go, we can use functions as values.

```go
func myFunction(){
  fn := func(){
    fmt.Println("Inside fn")
  }
  fn()
}

// Simplified version of above code.
func myFunction(){
  func(){
    fmt.Println("Inside fn")
  }()
}
// Notice how we execute it using the parenthesis at the end.
```

### Closures

- A closure is a function that references variables from outside of its body.

- Closures are lexically scoped, which means functions can access the values in scope when defining the function.

```go
func myFunction() func(int) int{
  sum := 0
  return func(v int) int {
    sum += v
    return sum
  }
}

...
add := myFunction()
add(5)
fmt.Println(add(10))
...
```

### Variadic Functions

which are functions that can take zero or multiple arguments using the `...` ellipses operator.

```go
func main(){
  sum := add(1,2,3,5)
  fmt.Println(sum)
}

func add(values ...int) int {
  sum := 0
  for _, v := range values{
    sum += v
  }
  return sum
}
```

> [!IMPORTANT]
> `fmt.Println` is a variadic function, that's how we were able to pass multiple values to it.

### Init

`init` is a special lifecycle function that is executed before the `main` function.

the `init` function does not take any arguments nor returns any value.

```go
package main
import "fmt"

func init(){
  fmt.Println("Before main!")
}

func main(){
  fmt.Println("Running main")
}
```

- `init` function was executed before the `main` function.

> [!IMPORTANT]
> There can be more than one `init` function in single or multiple files.
> For multiple `init` in a single file, their processing is done in the order of their declaration, while `init` functions declared in multiple files are processed according to the lexicographic filename order.

```go
package main

import "fmt"

func init(){
  fmt.Println("Before main!")
}

func init(){
  fmt.Println("Hello again")
}

func main(){
  fmt.Println("Running main")
}
```

> [!IMPORTANT]
> The `init` function is optional and is particularly used for any global setup which might be essential for our program, such as establishing a database connection, fetching configuration files, setting up environment variables, etc.

### Defer

the `defer` keyword, which lets us postpones the execution of a function until the surrounding function returns.

```go
func main(){
  defer fmt.Println("I am finished")
  fmt.Println("Doing some work....")
}
```

we can use multiple `defer` functions, using multiple `defer` brings us to what is known as `defer stack`.

```go
func main(){
  defer fmt.Println("I am finished")
  defer fmt.Println("Are you?")

  fmt.Println("Doing some work...")
}

//o/p:
// Doing some work...
// Are you?
// I am finished.

```

As we can `defer statements` are stacked and executed in LIFO manner.

So, Defer is incredibly useful and is commonly used for doing cleanup or error handling.

Functions can also be used with generics.

## Modules

### What are modules?

A module is a collection of [Go packages](https://go.dev/ref/spec#Packages) stored in a file tree with a `go.mod` file at its root, provided the directory is outside `$GOPATH/src`.

Go modules were introduced in Go 1.11, which brings native support for versions and modules.

> [!IMPORTANT]
> `GOPATH` is a variable that defines the root of your workspace and it contains the following folders:
>
> - **src**: contains Go source code orgnaized in a hierarchy.
> - **pkg**: contains compiled package code.
> - **bin**: contains compiled binaries and executables.
>   ![alt text](image.png)

- Create a new module using `go mod init` command which creates a new module and initializes the go.mod file that describes it.
  `go mod init example`

The important thing to note here is that Go module can correspond to a Github repository as well if you plan to publish this module. For example:
`go mod init example`

- `go.mod` which is the file that defines the module's module path and also the import path used for the root directory, and its dependency requirements.

```go
module <name>
go <version>

require (
  ...
)
```

And if we want to add a new dependency, we will use `go install` command:

- `go install github.com/rs/zerolog`

- `go.sum` file was also created. This file contains the expected hashes of the content of the new modules.

- `go list -m all`
- If the dependency is not used, we can simply remove it using `go mod tidy`

### Vendoring

Vendoring is the act of making your own copy of the 3rd party packages your project is using. Those copies are traditionally placed inside each project and then saved in the project repository.

This can be done through `go mod vendor` command.

- After the `go mod vendor` command is executed, a `vendor` directory will be created.

```bash
├── go.mod
├── go.sum
├── go.work
├── main.go
└── vendor
    ├── github.com
    │   └── rs
    │       └── zerolog
    │           └── ...
    └── modules.txt
```

## Packages

### What are packages?

A package is nothing but a directory containing one or more Go source files, or other Go packages.

This means every Go source file must belong to a package and package declaration is done at top of every source file as follows.

`package <package_name>`

so far, we've done everything inside of `package.main`. The `main` package should also contain a `main()` function which is a special function that acts as the entry point of an executable program.

> [!NOTE]
> Go also has a concept of imports and exports but it's very elegant.
> Basically, any value (like a variable or function) can be exported and visible from other packages if they have been defined with an upper case identifier.

- Example:

```go
package custom

var value int = 10    // Will not be exported.
var Value int = 20    // Will be exported.
```

- As we can see lower case identifiers will not be exported and will be private to the package it's defined in. In our case the `custom` package.

Here we can refer to it using the `module` we had initialized in our `go.mod` file earlier.

```go
----go.mod----
module example

go 1.18

----main.go----
package main
import "example/custom"

func main(){
  custom.Value
}
```

- Notice how the package name is the last name of import path.
- we can import multiple packages as well like this:

```go
package main

import (
  "fmt"
  "example/custom"
)

func main(){
  fmt.Println(custom.Value)
}
```

- we can also alias our imports to avoid collisions like this.

```go
package main

import (
  "fmt"
  abcd "example/custom"
)

func main(){
  fmt.Println(abcd.Value)
}

```

### External dependencies

- In go we can install external packages using `go install` command as we saw earlier.

`go install github.com/rs/zerolog`

```go
package main

import (
  "github.com/rs/zerolog/log"
  abcd "example/custom"
)

func main(){
  log.Print(abcd.Value)
}
```

> [IMPORTANT]
> Make sure to check out the go doc of packages you install, which is usally located in the project's readme file. go doc parses the source code and generates documentation in HTML format. Reference to it is usually located in readme files.

> [!NOTE]
> Go doesn't have a particular "folder structure" convention, always try your best to organize your packages in a simple and intuitive way.

## Workspaces

- multi-module workspaces were introduced in Go 1.18.
- Workspaces allow us to work with multiple modules simultaneously without having to edit `go.mod` files for each module. Each module within a workspace is treated as a root module when resolving dependencies.

```bash
# hello module
mkdir workspace && cd workspaces
mkdir hello && cd hello
go mod init hello
```

```go
package main

import (
  "fmt"
  "golang.org/x/example/stringutil"
)

func main(){
  result := stringutil.Reverse("Hello Workspace")
  fmt.Println(result)
}
```

```bash
$ go get golang.org/x/example
go: downloading golang.org/x/example v0.0.0-20220412213650-2e68773dfca0
go: added golang.org/x/example v0.0.0-20220412213650-2e68773dfca0


$ go run main.go
ecapskroW olleH
```

- WHat if we want to modify the `stringutil` module that our code depends on?
- Until now, we had to do it using the `replace` directive in the `go.mod` file, but now let's see how we can use `workspaces` directory.

- `go work init`

- This will create a `go.work` file.

```bash
cat go.work
go 1.18
```

- we will also add our `hello` module to the workspace.

- `go work use ./hello`

- This should update the `go.work` file with a reference to our `hello` module.

```bash
go 1.18
use ./hello
```

- download and modify the `stringutil` package and update the `Reverse` function implementation.

```bash
$ git clone https://go.googlesource.com/example
Cloning into 'example'...
remote: Total 204 (delta 39), reused 204 (delta 39)
Receiving objects: 100% (204/204), 467.53 KiB | 363.00 KiB/s, done.
Resolving deltas: 100% (39/39), done.
```

`example/stringutil/reverse.go`

```go
func Reverse(s string) string{
  return fmt.Sprintf("I can do whatever!! %s", s)
}
```

Finally, let's add `example` package to our workspace.

```bash
$ go work use ./example
$ cat go.work
go 1.18

use (
	./example
	./hello
)

# now if we run our hello module we will notice that the Reverse function has been modified.

go run hello
I can do whatever!! Hello Workspace
```

## Useful Commands

- `go fmt` which formats the source code and it's enforced by that language so that we can focus on how our code should work rather than how our code should look.

- `go vet` which reports likely mistakes in our packages.

- `go env` which simply prints all the go environment information.

- `go doc` which shows documentation for a package or symbol.
  - `go doc -src fmt Printf`

- `go help` command to see what other commands are available.

- `go fix` finds Go programs that use old APIs and rewrites them to use newer ones.
- `go generate` is usually used for code generation.
- `go install` compiles and installs packages and dependencies.

- `go clean` is used for cleaning files that are generated by compilers.

etc.

## Build

Building static binaries is one of the best features of Go which enables us to ship our code efficiently.

- we can do this very easily using the `go build` command.

```go
package main

import "fmt"
func main(){
  fmt.Println("I am a binary!")
}
```

- `go build` this should produce a binary with the name of our module.
- `go build -o app`

```bash
# Now to run this, we simply need to execute it.
./app
I am a binary!
```

> [I!IMPORTANT]
> Some important build time variables, starting with:
> `GOOS` and `GOARCH`
> These envionment variables help us build go programs for different operating systems and underlying processor architectures.

```bash
# we can list all the supported architecture using "go tool" command.

go tool dist list
android/amd64
ios/amd64
js/wasm
linux/amd64
windows/arm64
...
```

- An example for building a windows executable from macOS!
  - `GOOS=windows GOARCH=amd64 go build -o app.exe`
  - `CGO_ENABLED`
    This variable allows us to configure [CGO](https://go.dev/blog/cgo), which is a way in Go to call C code.
