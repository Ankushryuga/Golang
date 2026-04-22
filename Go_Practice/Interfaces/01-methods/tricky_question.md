# Go Methods 

## Basic questions

1. What is a method in Go?
- `A method is a function with a receiver attached to a type`.

- `A method is a function with a receiver. Receiver can be value or pointer receiver. Value receivers work on a copy, while pointer receivers can mutate the original and avoid copying. Receiver choice also affects method sets, which matters for interface satisfaction`.


2. What is receiver in Go?
- `The receiver is the parameter b/w "func" and the method name, such as "(u User)"`.

3. Can methods be defined on structs only?
- `No. They can be defined on any user-defined type.`

4. Can you define a method on `int` directly?
- `No. But you can define a new type based on "int" and add methods to that.`


## Value vs Pointer receiver.
5. What is the difference b/w `value receiver` and `pointer receiver`?
- Value receiver gets a copy.
- Pointer receiver can modify original value.

6. When should you use pointer receivers?
Use them when:
- you need mutation.
- the struct is large.
- you want consistent behavior.

7. Can a value call a pointer receiver method?
- Yes, if the value is addressable.

8. Can a map element call a pointer receiver method?
- No, because map elements are not addressable.

## Method Sets.
9. What is a method set?
- The set of methods a type has for interface satisfaction and method resolution.

10. What is the method set of `T`?
- It includes method with receiver `T`.

11. What is the method of `*T`?
- It includes methods with receiver `T` and `*T`.

12. Why does `T` not always satisfy an interface implemented by `*T` methods?
- Because `T`'s method set does not include pointer receiver methods.

## Interface and methods
13. How does Go know a type implements an interface?
- Implicitly, No implements keysword is needed.

14. If a method has pointer receiver, does the value type implements the interface?
- Usually no. The pointer type does.

15. Why are methods important for interfaces in Go?
- Because interfaces are satisfied by method sets.

## Advanced
16. What are method values?
- A method value binds the receiver:
```go
f := u.Hello
```

17. What are method expressions?
- A method expression treats method like a function:

```go
f := User.Hello
f(u, "Hi")
```

18. Can methods be defined on interfaces?
- No. Interfaces declare methods signatures, but do not define concrete methods like structs do.

19. Can methods be overloaded in Go?
- No. Go does not support method overloading.

20. Can methods be generic?
- Methods can use generic receiver type through generic types, but Go does not support type parameters declared on methods separately from the receiver type in the same way some languages do. 
=> Safe answer is:
- Go supports generics on types and functions.
- Method behavior depends on the generic type definition.
