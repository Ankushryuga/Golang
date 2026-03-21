# Maps

Maps are hash tables: key-value storage.

- Basic

```go
ages := map[string]int{
    "Sohan": 27,
    "Ramesh": 23,
}

fmt.Println(ages["Sohan"])
```

## Reading from map

If key does not exist, you get zero value.

```go
m := map[string]int{}
fmt.Println(m["missing"])       // 0
```

To distinguish missing vs present:

```go
v, ok := m["missing"]
if !ok {
    fmt.Println("not found")
}
```

This is essential

## Writing to map

```go
m := make(map[string]int)
m["a"]=1
m["b"]=2
```

## Delete from map

```go
delete(m, "a")
```

- Safe even if key does not exist.

## Nil map behavior

```go
var m map[string]int
fmt.Println(m["x"])     // 0
m["x"]=1                // Panic
```

You can read from nil map, but cannot write to it.

> [IMPORTANT]
> Initialize maps before writing

## Range over map

```go
for k, v := range m {
    fmt.Println(k, v)
}
```

Order is not guaranted.

> [!IMPORTANT]
> Never rely on map iteration order.
>
> If stable order is needed:
>
> - collect keys
> - sort keys
> - iterate in sorted order

## Map of struct vs map of pointers

Consider:

```go
type User struct{
    Name string
}

m := map[string]User{
    "1": {Name: "Ankush"},
}
```

This won't work:

```go
u := m["1"]
u.Name = "New"
```

It modifies a copy, not map value.

Need reassignment:

```go
u := m["1"]
u.Name = "New"
m["1"]=u
```

Alternative:

```go
m := map[string]*user{
    "1": {Name: "Ankush"},
}
m["1"].Name = "New"
```

> [!CAUTION]
>
> 1. **map of values:** safer, simpler ownership
> 2. **map of pointers**: easier mutation, more shared state complexity

## Using struct keys

Maps can use comparable structs as keys.

```go
type Key struct{
    UserID  int
    Env     string
}

m := map[Key]string{}
m[Key{UserId: 1, Env: "prod"}]="active"
```

- Useful for compound keys
