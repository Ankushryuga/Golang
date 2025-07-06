# Initialize your go Project:
    =>
    - mkdir myapp
    - cd myapp
    - go mof init github.com/username/myapp


# Directory structure:
    =>
    myapp/
    ├── go.mod
    ├── go.sum
    ├── cmd/              # Main application entry points (one folder per CLI or binary)
    │   └── myapp/
    │       └── main.go
    ├── internal/         # Private application and library code
    │   └── utils/        # Example: helper functions
    ├── pkg/              # Public library code (optional)
    ├── api/              # API definitions (e.g. OpenAPI, gRPC)
    ├── configs/          # Configuration files (YAML, JSON, etc.)
    ├── scripts/          # Scripts for devops, builds, etc.
    └── README.md


  # Install Dependencies:
    =>
    go get github.com/gorilla/mux

# Build & Run:
    =>
    go build -o bin/myapp ./cmd/myapp
    ./bin/myapp
    
    or just run without building..
    go run ./cmd/myapp

# Run Tests:
    =>
    go test ./...

# Format & Tidy code:
    =>
    go fmt ./...
    go mod tidy    //cleanup unused dependencies..


## 🧰 Tools You Might Use

| Tool             | Purpose                       |
| ---------------- | ----------------------------- |
| `go fmt`         | Format code                   |
| `go test`        | Run unit tests                |
| `go mod`         | Dependency management         |
| `golangci-lint`  | Code linting                  |
| `air` or `fresh` | Hot reload for local dev      |
| `godotenv`       | Load `.env` files for configs |



## summary:
| Step        | Command / Action                    |
| ----------- | ----------------------------------- |
| Init module | `go mod init <module-name>`         |
| Add dep     | `go get <package>`                  |
| Run app     | `go run ./cmd/myapp`                |
| Build app   | `go build -o bin/myapp ./cmd/myapp` |
| Structure   | Use `cmd/`, `internal/`, `pkg/`     |
| Clean deps  | `go mod tidy`                       |

