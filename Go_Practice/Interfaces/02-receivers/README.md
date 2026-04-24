# Receivers in Golang

A `receiver` is the special parameter that lets you attach a method to a type.

Example:
```go
type User struct {
    Name string
}

func (u User) Greet() {
    fmt.Println("Hello", u.Name)
}
```
Here, `u User` is the receiver, and `Greet` is a method of `User`.

- You call it like this:

```go
user := User{Name: "Ankush"}
user.Greet()
```
- There are two main receiver types:

1. Value receiver
```go
func (u User) Greet() {
    fmt.Println(u.Name)
}
```