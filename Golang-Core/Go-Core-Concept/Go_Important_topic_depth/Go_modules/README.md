# Modules

Go modules are the dependency and versioning system for Go.

## What is a module?

A module is a collection of Go packages versioned together, defined by `go.mod`.

- Create one:

```bash
go mod init github.com/ankush/myapp
```

This creates:

```go
module github.com/ankush/myapp

go 1.24
```

## Why modules matter?

Modules solve:

1. dependency management
2. version pinning
3. reproducible builds
4. import path identity

## Basic commands

### Initialize module

```bash
go mod init github.com/ankush/myapp
```

- Add dependencies automatically

```bash
go get github.com/gin-gonic/gin
```

- Clean dependencies

```bash
go mod tidy
```

- Vendor dependencies

```bash
go mod vendor
```

- Download dependencies

```bash
go mod download
```

`go.mod` and `go.sum`

- `go.mod`
  Defines:
  1.  module path
  2.  Go version
  3.  required dependencies

- `go.sum`
  Contains cryptographic checksums for dependencies

> [!IMPORTANT]
> commit both files to version control

### Import paths module identity

- If module path is:

```bash
module github.com/ankush/myapp
```

Then packages are imported like:

```go
import "github.com/ankush/myapp/internal/user"
```

The module path is not just a label.
It defines package identity

### Replace directive

Useful for local development

```bash
replace github.com/company/sharedlib => ../sharedlib
```

Good for local multi-module work.

> [!CAUTION]
> don't accidiently ship broken local `replace` setups.

### Multi-module vs single-module repo

- start with a single module unless there is a strong reason not to
- multiple modules add complexity
- use multi-module only for independently versioned components or public libs

> [!IMPORTANT]
> **Dependency hygiene**
>
> 1. minimal dependencies
> 2. stable dependencies
> 3. security of dependencies
> 4. dependency ownership
>
> Rules:
>
> 1. Avoid adding heavy frameworks blindly
> 2. Prefer standard library when enough
> 3. audit critical dependencies
> 4. understand what you import
