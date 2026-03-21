# Structs

A `struct` is Go's way to grouping related data together.

- Beginner view

```go
package main

import "fmt"

type User struct{
    ID      int
    Name    string
    Email   string
}

func main(){
    u := User{
        ID:     1,
        Name:   "Ankush",
        Email:  "ankush@example.com"
    }
    fmt.Println(u.Name)
}
```

> [!NOTE]
> A struct is like custom data type with named fields.
> **Struct are everywhere in Go:**

1. request/response models
2. database entities
3. configs
4. domain models
5. service dependencies
6. state containers

> [!NOTE]
> **Zero value behavior**
> One of the most important Go ideas: every type has a zero value.

- For a struct

```go
var u User
fmt.Println(u.ID)       // 0
fmt.Println(u.Name)     // ""
fmt.Println(u.Email)    // ""
```

Questions?

- Can this type be safely used without explicit initialization?
- Does the zero value represent something meaningful?
- Should construction be forced through a constructor?

## Exported vs Unexported fields

If a name starts with uppercase, it is exported outside the package.

```go
type User struct{
    ID      int
    Name    string
    email   string
}
```

Here:

- `ID`, `Name` are exported.
- `email` is private to the package

> [!IMPORTANT]
> Keep fields unexported unless callers truly need direct access.
> Expose behavior, not internal representation.

## Anonymous struct

Useful for small local cases.

```go
cfg := struct {
    Host string
    Port int
}{
    Host: "localhost"
    Port: 8080
}
```

> [!NOTE]
> **Anonymous structs** are Good for:
>
> 1. tests
> 2. local grouping
> 3. one-off marshaling/unmarshaling

> [!CAUTION]
> Not good for:
> reusable domain models

## Struct tags

Used for JSON, DB, validation, etc.

```go
type User struct{
    ID      int     `json:"id"`
    Name    string  `json:"name"`
    Email   string  `json:"email"`
}
```

> [!CAUTION]
>
> - tags are metadata, not business logic
> - do not overload one struct for every layer if meanings differ
>   **Example** Don't always use the same struct for:
>
> 1. DB model
> 2. API response
> 3. Internal domain object
>
> That creates tight coupling

## Embedded structs

Go supports composition by embedding.

```go
type Address struct{
    City        string
    Country     string
}

type User struct{
    Name    string
    Address
}
```

### Usage

```go
u := User{
    Name:   "Ankush",
    Address: Address{
        City:       "Chennai",
        Country:    "India",
    },
}

fmt.Println(u.City)
```

This is composition, not inheritance.

> [!NOTE]
>
> 1. Shared behavior or fields
> 2. Common metadata
> 3. Decorators or wrappers

> [!CAUTION]
>
> - embedding can make APIs confusing if overused.
> - explicit fields are often clearer

## Struct Comparision

Structs are comparable only if all fields are comparable.

```go
type Point struct{
    X int
    Y int
}

a := Point{1, 2}
b := Point{1, 2}

fmt.Println(a==b)           // true
```

But if struct contains slice/map/function, it is not comparable.

```go
type Bad struct{
    Item    []string
}
```

`Bad` cannot be compared with `==`.

## When to use pointer vs value for structs

This is one most important Go topics.

### Value

```go
func printUser(u User){
    fmt.Println(u.Name)
}
```

This copies the struct.
Use value when:

- struct is small
- function should not modify original
- type behaves like data/value object

```go
func renameUser(u *User, name string){
    u.Name = name
}
```

Use Pointer when:

- function/method modifies receiver
- struct is large and copying is undesirable
- you want shared mutable state
- method set requires pointer receiver

> [!NOTE]
> Choose intentionally, not randomly
> be consitent for a type's methods
