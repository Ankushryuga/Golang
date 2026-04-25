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

This receives a **copy** of the value. Use it when the method does not need to modify the original struct.

2. Pointer receiver
```go
func (u *User) Rename(name string){
    u.Name = name
}
```

This receives the address of the value. Use it when the method needs to modify the original struct or avoid copying a large struct.

Example: 
```go
type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func main() {
    counter := Counter{Value: 1}
    counter.Increment()

    fmt.Println(counter.Value) // 2
}
```
