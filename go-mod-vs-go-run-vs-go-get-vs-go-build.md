# go mod:   Module Management
        => manages project dependencies and module settings.
        | Command                | Description                                   |
| ---------------------- | --------------------------------------------- |
| `go mod init <module>` | Initializes a new module                      |
| `go mod tidy`          | Adds missing or removes unused dependencies   |
| `go mod download`      | Downloads all modules to the local cache      |
| `go mod verify`        | Checks that downloaded modules are unmodified |
| `go mod graph`         | Shows dependency graph                        |


# go run:  quick compile and execute without creating a binary.


# go build: compile the program into an executable binary

# go get: add remote dependencies
    =>
    example: go get github.com/gin-gonic/gin

# go install: Build and Install Binary.
    =>
    Purpose:
    Builds and installs a binary tool into your $GOBIN or $GOPATH/bin.
    # example:
    go install github.com/some/tool@latest


# go test: Run unit test:
    =>
    1. go test ./...

    2. go test -v myfile_test.go
    Add -race to check for race conditions.



#  go fmt — Format Go Code
    =>
    Automatically formats your .go files to standard Go style.
    # example:
    go fmt ./...


# go doc:
    =>
    go doc fmt.Println



| Command       | Purpose                              | Creates Binary? | Modifies go.mod? | Use Case             |
| ------------- | ------------------------------------ | --------------- | ---------------- | -------------------- |
| `go mod init` | Initialize module                    | ❌               | ✅                | Start new Go project |
| `go mod tidy` | Clean up dependencies                | ❌               | ✅                | Maintain go.mod      |
| `go run`      | Compile & run (no binary saved)      | ❌               | ❌                | Quick test run       |
| `go build`    | Build binary                         | ✅               | ❌                | Build apps locally   |
| `go get`      | Add dependency                       | ❌               | ✅                | Add module packages  |
| `go install`  | Build + globally install binary tool | ✅               | ❌                | Install CLI tools    |
| `go test`     | Run tests                            | ❌               | ❌                | Testing              |
| `go fmt`      | Auto-format code                     | ❌               | ❌                | Format files         |
| `go doc`      | Show docs                            | ❌               | ❌                | Learn APIs           |

