# Receiver Interview questions:

- [What is a Receiver?](#what-is-a-receiver)


## What is a Receiver?
- A receiver is used to define a method on a type.
```go
type User struct {
    Name string
}

func (u User) SayHello() {
    fmt.Println("Hello", u.Name)
}
```
Here:
```go
(u User)    // is the receiver

// calling receiver:
user := User{Name: "Ankush"}
user.SayHello()
```

