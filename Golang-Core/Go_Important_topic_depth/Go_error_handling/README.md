# Error handling

Go does not use exceptions for routine errors.
`Errors are values`.

- Basic Pattern:

```go
result, err := doSomething()
if err != nil {
    return err
}
```

This is the core pattern of Go.

- Create errors:

```go
errors.New("Something went wrong")

// or
fmt.Errorf("failed to fetch user: %w", err)
```

Use `%w` to wrap errors.

- Why wrapping matters?

```go
func loadUser(id string) error {
    err := dbQuery(id)
    if err != nil {
        return fmt.Errorf("load user %s: %w", id, err)
    }
    return nil
}
```

This adds context while preserving the original error.

> [!IMPORTANT]
> Always add useful context at boundaries.

- Bad: `return err`

- Better: `return fmt.Errorf("fetch profile from repo: %w", err)`

## Sentinel errors

`var ErrNotFound = errors.New("not found")`

Usage:

```go
if errors.Is(err, ErrNotFound){
    // handle not found
}
```

Use Sentinel errors for meaningfull program decisions.

Examples:

- not found
- unauthorized
- conflict
- timeout

But don't create dozens of vague sentinels

## Custom error types

```go
type ValidationError struct {
    Field       string
    Msg         string
}

func (e ValidationError) Error() string{
    return e.Field + ": " + e.Msg
}
```

Use:

```go
func validateEmail(email string) error {
    if email == ""{
        return ValidationError {
            Field: "emai",
            Msg:   "required",
        }
    }
    return nil
}
```

Then:

```go
var ve ValidationError
if errors.As(err, &ve){
    fmt.Println("validation field:", ve.Field)
}
```

Use custom types when callers need structured error information.

### errors.Is and errors.As

- `errors.Is`
  Checks for specific wrapped error.

```go
if errors.Is(err, ErrNotFound){
    // ...
}
```

- `errors.As`
  Checks for a specific error type.

```go
var ve ValidationError

if errors.As(err, &ve) {
    // ...
}
```

> [!NOTE]
> Never do like following:

#### Never do like this

```go
if err.Error() == "not found" {
    ...
}
```

### Error handling guidelines

1. Add context

```go
return fmt.Errorf("create order: %w", err)
```

2. Don't log and return the same error at every layer
   Avoid duplicate logs

Pick the right layer to log, often near process boundry:

- HTTP handler
- consumer loop
- job runner
- main

3. Use typed or sentinel errors for control flow

Example:

- `ErrNotFound`
- `ErrUnauthorized`
- `ErrConflict`

4. Do not panic for normal failures
   Panic is for programmar bugs or unrecoverable situations, not validation or DB misses.

5. Return early
   Go code is clearer with early returns

- Bad:

```go
if err == nil {
    // huge block
}
```

- Better:

```go
if err != nil {
    return err
}
// continue
```

Example: layered errors

```go
func (r *UserRepo) GetByID(ctx context.Context, id string) (User, error){
    user, err := r.queryUser(ctx, id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return User {}, ErrNotFound
        }
        return User{}, fmt.Errorf("query user by id %s: %w", id, err)
    }
    return user, nil
}

func (s *UserService) GetProfile(ctx context.Context, id string) (User, error){
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return User{}, fmt.Errorf("get profile: %w", err)
    }
    return user, nil
}
```

At HTTP layer:

```go
user, err := svc.GetProfile(ctx, id)
if err != nil {
    switch {
        case errors.Is(err, ErrNotFound):
            http.Error(w, "not found", http.StatusNotFound)
        default:
            http.Error(w, "internal error", http.StatusInternalServerError)
    }
    return
}
```
