# Stage 1: Go Basic:
    =>
    1. What is Go?, setup and tools (go run, go build, go mod)
    2. Basic syntax: main, package, import
    3. Variables, Constants, Types.
    4. Conditions: if, switch
    5. Loops: for, range
    6. Functions, multiple return values
    7. Error handling: error type

    ## Practice:
    a. Calculator CLI
    b. Temperature Converter



# Stage 2: Go Intermediate:
    =>📌 Goal: Learn Go's powerful type system and struct-based design.
    1. Arrays, Slices, Maps, Custom collections.
    2. structs and receivers
    3. Pointer in Go
    4. Interfaces and type assertions
    5. Methods vs Functions
    6. Packages and moduler code (go mod)
    7. Built-in packages: fmt, os, io, bufio, strconv, net/http
    
    ## Practice:
    a. Todo app
    b. CLI file reader with stats
    c. Build REST API using net/http


# Stage 3: Concurreny and channels:
    =>📌 Goal: Master Go’s standout feature: goroutines and channels.
    1. Goroutines (go keyword)
    2. Channels (unbuffered, buffered)
    3. Select statement
    4. Mutex, WaitGroups, sync package
    5. Deadlocks and race conditions

    ## Practice:
    a. Parallel File downloader
    b. Worker pool
    c. Chat server (basic)


# Stage 4: web API and Frameworks
    => 📌 Goal: Build robust APIs using Go frameworks and tools.
    1. REST API: using net/http
    2. Popular frameworks: Gin, Echo, Fiber
    3. Middleware, CORS, validation
    4. JSON marshalling/unmarshalling
    5. struct tags (json:"...")
    6. Dependency Injection
    7. Environment config (viper, .env)

    ## Practice:
    a. User management API
    b. JWT authentication with middleware
    c. Blogging platform


# Stage 5: Advanced Go
    => 📌 Goal: High-performance, scalable, production-ready Go apps.
    1. Testing: testing package, table tests, mocks
    2. Benchmarking & profiling
    3. Custom errors and wrapping
    4. Interfaces & composition (not inheritance)
    5. gRPC in Go
    6. Reflection
    7. Advanced error handling (errors, fmt.Errorf)
    8. Context package
    9. REST VS gRPC in microservices

    ## Practice:
    a. Microservice with gRPC + protobuf
    b. File upload service
    c. Rate Limiter using channels


# Stage 6: DevOps & Deployment:
    => 📌 Goal: Containerize, deploy, and monitor Go apps.
    1. Dockerize Go apps
    2. Build minimal images with multi-stage builds
    3. Use go install, go build properly
    4. Go with kafka, RabbitMQ
    5. Prometheus for monitoring
    6. CI/CD pipeline
    7. Deploy to: AWS (EC2, Lambda), Kubernetes.

    ## Practice:
    a. Dockerize your go microservices
    b. Deploy gRPC + REST services with Docker compose.


## BONUS:
    => 
    1. Generics in Go
    2. Go + PostgreSQL (with sqlx, gorm, pgx)
    3. Clean architecture in go.
    4. Event-driven design
    5. go mod for dependency management
    6. go test, benchstat, pprof
    7. Gin, GORM, protobuf, grpc-go, Docker

