# Interfaces

Interfaces are one of Go's most powerful ideas, and also one the most misused.

- Beginner definition
  An Interface defines behavior

```go
type Speaker interface {
    Speak() string
}
```

Any type with a `Speak() string` method implements it automatically.

```go
type Dog struct{}

func (Dog) Speak() string {
	return "woof"
}
```

> [!NOTE]
> No explicit `implements` keyword.

### The beauty of implicit implementation

This reduces coupling.

```go
func SaySomething(s Speaker){
    fmt.Println(s.Speak())
}
```

And any type with right method workds.

### Small interfaces are best

Idiomatic Go prefers tiny interfaces.

Example from standard library:

- `io.Reader`
- `io.Writer`
- `fmt.Stringer`

```go
type Logger interface{
    Info(msg string)
}
```

- **accept interfaces, return structs.**

Why?

- callers can pass different implementations
- returning concrete types gives more power and clarity

```go
func NewService(repo UserRepository) *Service
```

This accepts an interface and returns a concrete service.

> [!CAUTION]
> **Don't define interface too early**

- BAD interface:

```go
type UserServiceInterface interface{
    CreateUser(ctx context.Context, req CreateUserRequest) (User, error)
    UpdateUser(ctx context.Context, req UpdateUserRequest) (USer, error)
    DeleteUser(ctx context.Context, id string) error
    ...
}
```

This is often unnecessary.

- Better Approach:
  - define interface at the point of use.
  - Keep them small
  - define them for testing or abstraction boundries

```go
type userGetter interface{
    GetById(ctx context.Context, id string) (User, error)
}
```

If a service only needs one method, ask for one method.

### Interface satisfaction example:

```go
type Notifier interface{
    Send(to, msg string) error
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(to, msg string) error{
    fmt.Println("email sent to", to)
    return nil
}
```

- Use it

```go
func notifyWelcom(n Notifier, email string) error {
    return n.Send(email, "welcome")
}
```

### Empty interface and "any".

Before Go 1.18 people used `interface{}`.
Now `any` is an alias for `interface{}`.

```go
func Print(v any){
    fmt.Println(v)
}
```

> [!CAUTION]
>
> - `any` often means loss of type safety
> - prefer concrete types or generic type parameters when possible.

### Type assertion

```go
var v any = "hello"

s, ok := v.(string)
if ok{
    fmt.Println(s)
}
```

without `ok`, it can panic.

```go
s := v.(string) // panic if not string
```

### Type Switch

```go
func Describe(v any){
    switch x := v.(type) {
        case string:
            fmt.Println("string:", x)
        case int:
            fmt.Println("int:", x)
        default:
            fmt.Println("unknown")
    }
}
```

Useful, not something to overuse.

## Nil Interface trap

> [!IMPORTANT]

```go
type CustomError struct{}

func (e *CustomError) Error() string {
    return "custom error"
}

func fail() error {
    var e *CustomError = nil
    return e
}
```

This returns a non-nil interface holding a nil pointer.
so `err != nil` becomes true.

> [!CAUTION]
> Be careful returning typed nils as interfaces.

> [!IMPORTANT]
> questions before creating designing interfaces.
>
> - What exact behavior do i need?
> - Can this interface be smallers?
> - Is this abstraction real or premature?
> - Is it defined by consumer instead of producer?
>
> **Best Patterns**:
>
> - Concrete types internally
> - Small consumer-owned interfaces at boundries
