## 1. create a new Go module in your project.

`go mod init github.com/Ankushryuga/url-shortener/Go_Backend`.


📦 What is a Go module?
  - A module is:
      - A collection of Go code (your project)
      - With its own dependencies
      - Managed using a go.mod file
      - Think of it like:
          - package.json in Node.js
          - requirements.txt / pyproject.toml in Python


Running go mod init enables:
- ✅ Dependency management (go get, go mod tidy)
- ✅ Version control of libraries
- ✅ Clean builds
- ✅ Reproducible environments



go mod init = “start a new project with dependency tracking”
go.mod = “project identity + dependencies”


## 2. go get <module-name>

👉 go get is used to download and add external dependencies (packages) to your Go project.

`go get github.com/gin-gonic/gin`
👉 Go will:
- Download the Gin package
- Add it to your go.mod
- Save checksums in go.sum

## 3. go tidy


## 4. Lint : golangci-lint run -h
