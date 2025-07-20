# Go Unit-Testing and Benchmarking
- Unit Testing
- Benchmarking
- Table driven Test
- Parallel Testing


# Unit Testing in Go
# File structure: Create a test file alongside your code, named like:
      math_utils.go
      math_utils_test.go
      

# Benchmarking in go: Benchmark function must start with Benchmark, take a *testing.B, and use b.N to loop.
      func BenchmarkAdd(b *testing.B){
            for i:=0;i<b.N;i++{
                  Add(1,2)
            }
      }


## Run tests and benchmarks
      Run tests
            - go test
      Run with verbose o/p:
            go test -v
      Run Benchmark
            go test -bench=.


| Feature      | Usage                                       |
| ------------ | ------------------------------------------- |
| Unit Test    | `go test`                                   |
| Benchmark    | `go test -bench=.`                          |
| Table-driven | Struct slice + loop                         |
| Parallel     | `t.Parallel()`                              |
| Assertions   | `t.Errorf`, `t.Fatalf`, or `testify/assert` |
