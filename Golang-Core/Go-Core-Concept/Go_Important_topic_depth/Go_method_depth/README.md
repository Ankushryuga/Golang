# Methods

A method is a funtion attached to a type.

```go
type User struct{
    Name    string
}

func (u User) Greet() string{
    return "Hello, "+ u.Name
}
```

- Usage:

```go
u := User{Name: "Ankush"}
fmt.Println(u.Greet())
```

## Value receiver vs Pointer receiver

### Value receiver

```go
func (u User) Rename (name string){
    u.Name = name
}
```

This changes only the copy.

### Pointer receiver

```go
func (u *User) Rename(name string){
    u.Name  = name
}
```

This changes the original.

> [!NOTE]
> **Rule for receivers**
>
> If any method on a type needs pointer receiver, usually all methods should use pointer receiver for consistency.

```go
func (u User) FullName() string
func (u *User) Rename(name string)
```

Sometimes it is okay, but consistency matters.

Recommended when type is mutable:

```go
func (u *User) FullName() string
func (u *User) Rename(name string)
```

## Method sets

This becomes important for interfaces.

If a method has pointer receiver:

```go
func (u *User) Save() errror
```

Then only `*User` implements interfaces requiring `Save`, not `User`.

### Methods are behavior, not just helpers

> [!NOTE]
>
> - Put behavior close to the type when it truly belongs there
> - Don't create clases mentally
> - Don't force **OOP** into Go.

Good:

```go
type Money struct{
    Amount      int64
    Currency    string
}

func (m Money)  Add(other Money) (Money, error){
    if m.Currency != other.Currency {
        return Money{}, fmt.Errorf("currency mismatch")
    }

    return Money {
        Amount:     m.Amount + other.Amount,
        Currency:   m.Currency,
    }, nil
}
```

This behavior naturally belogns to `Money`.

Not everything needs methods.
Many functions are better as plain package functions.
