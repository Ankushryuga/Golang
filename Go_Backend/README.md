# Backend Mastering with Go

## Error handling

- `errors` in go are usually the last arguments in the return list.

```go
func processTruck(truck Truck) (string, float64, error){}
```

> [!NOTE]
> In go, we handle errors in a declarative manner.

```go
// declarative way to handle error.
err := processTruck(truck)
if err != nil {
    // some code.
}


// another way to handle error.
if err := processTruck(truck); err != nil { // benefit of it is that err variable desn't exist, it only exists here, garbage collector is just going to pick this err variable and going to delete them.
    // some code.
}
```


## Inteface 
