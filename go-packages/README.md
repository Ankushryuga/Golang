# Packages in Go

A package in made up of Go files that live in the same directory and have the same package statement at the begining.

- Some packages are available through the Go standard library and are therefore installed with your Go installation. Others can be installed with Go's `go get` command.

- You can also build your own packages by creating Go files in the same directory across which you want to share code by using the necessary package statement.

- Package contain definitions of functions, `types`and `variables` that can then be used in other Go programs.

## Example

```bash
└── $GOPATH
    └── src
        └── github.com
            └── gopherguides
                └── greet
                    └── greet.go
                    └── greet_test.go

```

- Examples coming soon
- 
