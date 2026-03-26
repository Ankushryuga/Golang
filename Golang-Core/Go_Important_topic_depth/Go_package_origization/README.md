# Organize package in Go

## What is a package?

A package is a directory of Go files with the same package name.

```go
package user
```

packages are Go's main unit of code organization.

## Core package design principles

1. Organize by domain, not by technical layer only

- Bad:
  - `models`
  - `utils`
  - `helpers`
  - `services`
  - `controllers`
    These become junk drawers

- Better:
  - `user`
  - `order`
  - `billing`
  - `auth`
    This aligns code with business concepts

2. Avoid package names like `common`, `util`, `misc`

3. Keep packages cohesive
   A package should have one main responsibilty.

- Bad:
  `user` package that handles:
  - `HTTP`
  - `DB`
  - `email`
  - `JWT`
  - `CLI`

Better to separate concerns

4. Avoid circular dependencies.
   Go forbids import cycles

This is a feature, not a limitation.
It forces cleaner boundaries.

If package A imports B and B imports A, your design is likely wrong.

Fix by:

- moving shared contract downward
- introducing smaller interfaces
- extracting neutral package
- rethinking ownership

## Typical project layout

A practical layout:

```bash
myapp/
  cmd/
    api/
      main.go
  internal/
    user/
      model.go
      service.go
      repository.go
      handler.go
    auth/
      service.go
      middleware.go
    platform/
      db/
        postgres.go
      log/
        logger.go
      config/
        config.go
  pkg/
  go.mod
  go.sum
```

Meaning of common directories

1. `cmd/`

Entry points
Each subdir builds a binary

Example:

- `cmd/api`
- `cmd/worker`
- `cmd/migrate`

2. `internal/`
   Private application code.
   Cannot be imported by external modules.

This is extremely useful for protecting internals

3. `pkg/`

Used for reusable public packages.
Use carefully
Many teams overuse `pkg`.

- use `internal/` by default
- use `pkg/` only for truly reusable public code.

**Package naming rules**

Good package names are:

- short
- lowercase
- singular usually
- meaningful

Good:

- user
- auth
- config
- db

Avoid:

- userService
- utils
- helpers
- manager

Package name should help code read naturally.

Example:

```go
user.NewService(...)
auth.ValidateToken(...)
config.Load()
```

### Avoid stutter.

Bad:

```go
package user

type UserService struct{}
```

Usage:

```go
user.UserService
```

Better:

```go
package user
type Service struct{}
```

Usage:

```go
user.Service
```

Similarly:

- user.Repository
- auth.Middleware
- config.Config

### Interfaces belong to the consumer.

This is very important for package design.

Bad:
repository package exports giant interface for itself.

Better:
consumer defines the tiny interface it needs.

```go
package service

type userGetter interface{
    GetById(ctx context.Context, id string) (user.User, error)
}
```

This reduces coupling and keeps abstractions honest.

### Internal layering example:

A clean flow:

1. `handler` accepts HTTP request
2. `service` applies business logic
3. `repository` talks to DB
4. `model` defines core domain types.

Example structure:

```bash
internal/user/
  entity.go
  service.go
  repository.go
  postgres_repository.go
  http_handler.go
```

Or split by subpackage if large:

```bash
internal/user/
  user.go
  service/
  repository/
  transport/http/
```

> [!IMPORTANT]
> Choose based on size.
> don’t over-engineer too early.
> Let structure grow with complexity.

## package architecture patterns:

### Pattern 1: Domain-oriented packages

```bash
internal/
  user/
  order/
  invoice/
```

Each domain owns:

1. types
2. business rules
3. repository contracts
4. use cases
   This often works very well.

### Pattern 2: Platform packages

Shared technical capabilities:

```bash
internal/platform/
  db/
  log/
  config/
  email/
```

These support domains but do not own business logic.

### Pattern 3: Clear dependency direction

A healthy dependency direction often looks like:

- transport depends on service
- service depends on repository interface
- repository implementation depends on DB
  domain

## Putting it together: example mini backend design

```go
package user

import (
    "context"
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("user not found")

type User struct{
    ID      string
    Name    string
    Email   string
}

func (u User) Validate() error {
    if u.ID == "" {
        return fmt.Errorf("id is required")
    }
    if u.Email == "" {
        return fmt.Errorf("email is required")
    }
    return nil
}

type Repository interface {
    Save(ctx context.Context, user User) error
    GetById(ctx context.Context, id string) (User, error)
}

type Service struct {
    repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, user User) error {
    if err := user.Validate(); err != nil {
        return fmt.Errorf("validate user: %w", err)
    }
    if err := s.repo.Save(ctx user); err!=nil {
        return fmt.Errorf("save user: %w", err)
    }
    return nil
}

func (s *Service) Get(ctx context.Context, id string) (User, error){
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return User{}, fmt.Errorf("get user: %w", err)
    }
    return user, nil
}
```

## Common Mistakes

**Mistake 1: Huge interfaces**
Correction:

- use small, focused interfaces

**Mistake 2: Package names like utils**

Correction:

- package by domain or clear capability

**Mistake 3: Returning stringly typed errors**

Correction:

- wrap errors and use errors.Is / errors.As

**Mistake 4: Overusing pointers**

Correction:

- use values when type is small and immutable-like

**Mistake 5: Reusing one struct for DB, API, validation, and business logic**

Correction:

- separate models where semantics differ

**Mistake 6: Ignoring nil slice vs empty slice**

Correction:

- define API semantics intentionally

**Mistake 7: Relying on map iteration order**

Correction:

- sort keys when order matters

**Mistake 8: Java-style inheritance mindset**

Correction:

- prefer composition, simple structs, small interfaces

**Mistake 9: Premature abstractions**

Correction:

- start concrete, abstract only where needed

**Mistake 10: Logging errors at every layer**

Correction:

- wrap through layers, log at boundary

## Additional informations.

need strong command over:

- `context.Context`
- concurrency with goroutines/channels
- mutexes and race conditions
- testing and table-driven tests
- dependency injection in Go style
- JSON encoding/decoding
- HTTP servers and middleware
- database/sql or ORM tradeoffs
- profiling and memory allocations
- benchmarking
- clean API design
- observability: logs, metrics, tracing
- graceful shutdown
- configuration management
- code review quality
- backward compatibility in public packages

These build on top of structs, methods, interfaces, slices, maps, errors, modules, and package design.

## Design Checklist

### Structs

- Is this struct representing one clear concept?
- Are fields exposed only when needed?
- Does zero value make sense?

### Methods

- Does this behavior truly belong on the type?
- Should receiver be pointer or value?
- Is method set consistent?

### Interfaces

- Is the interface small?
- Is it defined by the consumer?
- Is abstraction justified?

### Slices/maps

- Am I aware of sharing underlying memory?
- Am I handling missing keys properly?
- Do I need stable iteration order?

### Errors

- Am I adding context?
- Am I using errors.Is / errors.As?
- Am I avoiding string comparisons?
- Am I logging only at the right layer?

### Modules/packages

Are package boundaries clear?
Are imports one-directional?
Did I avoid junk-drawer packages?
Is this layout easy for a new engineer to navigate?
