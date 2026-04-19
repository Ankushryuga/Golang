# Best tricky practice questions

Q1: Why does this not update?

```go
type User struct{ Age int }
func (u User) Grow() { u.Age++ }
```

Q2: Why does this interface assignment fail?

```go
type Speaker interface { Speak() }
type Dog struct{}
func (d *Dog) Speak() {}
var s Speaker = Dog {}
```

Q3: Why does this not compile?

```go
m := map[string]User{"a": {Age: 1}}
m["a"].Grow()
```

Q4: What is the difference b/w these?

```go
f1 := u.Show
f2 := User.Show
```


Q5: Why would you choose all pointer receiver on a type even for read-only methods?
Ans:
```bash
- Consistency
- Avoid confusion with interfaces
- Avoid copying large structs.
```