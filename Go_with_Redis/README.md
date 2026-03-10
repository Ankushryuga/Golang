## Redis
Redis is an open-souce, in-memory data structure store that can be used as a cache, message broker, or a persistent `key-value` database.


### Redis Golang client:
  - `go get github.com/redis/go-redis/v9`.

### Setting up Redis connection:
To start using Redis in your Golang project, import the Redis client package, and establish connection with your `Redis server`.

```go
package main
import (
"context"
"fmt"
"log"
"time"

"github.com/redis/go-redis/v9"
)

var ctx = context.Backgroud()

func main() {
  // Connect to redis
  client := redis.Newclient(&redis.Options{
    Addr:        "localhost:6379"      // Redis server address.
    Password:    "",                   // No password for local development
    DB:          0,                    // Default DB
  })

  // Ping the redis server to check the connection
  pingServer, err := client.ping(ctx).Result()
  if err != nil {
    log.Fatal("Error connecting to Redis:", err)
  }
  fmt.Println("Connected to Redis:", pingServer)
}
```

- Another way is using connection string:
```go
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
  panic(err)
}

client := redis.NewClient(opt)
```

### Basic operations:

1. Setting and Getting values:
> [!NOTE]
> Redis command accepts a context that you can use to set [timeout](https://redis.uptrace.dev/guide/go-redis-debugging.html#timeouts) or propagate some information, for example, [tracing context](https://redis.uptrace.dev/guide/go-redis-monitoring.html).

```go
// Set a key-value pair
err := client.Set(ctx, "greeting", "Hello, Redis!", 0).Err()
if err!=nil {
  log.Fatal(err)
}

// Get the value for a key
value, err := client.Get(ctx, "greeting").Result()
if err!=nil{
  log.Fatal(err)
}
fmt.Println("Greetings:", value)
```

2. Working with Lists:
- Redis provides list data structures that can be useful for implementing queues.


```go
// LPush nsert values at the head of the list
err = client.Lpush(ctx, "tasks", "task 1", "task 2").Err()
if err != nil {
  log.Fatal(err)
}

// RPush insert values at the back/tail of the list
err = client.RPush(ctx, "tasks", "task 1", "task 2").Err()
if err != nil {
  log.Fatal(err)
}

// LPop remove the first element in the list
task, err := client.Lpop(ctx, "tasks").Result()
if err!=nil{
  log.Fatal(err)
}
fmt.Println("Popped tasks:", task)

// RPopr remove the last element in the list
task, err := client.RPop(ctx, "tasks").Result()
if err != nil {
  log.Fatal(err)
}
fmt.Println("Popped Tasks:", task)
```


3. Using Hashes
- Hashes in Redis allow you to store multiple field-value pair under a single key.

```go
// Set hash field-values
err = client.HSet(ctx, "user:1", map[string]interface{}{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25,
}).Err()

if err!=nil {
log.Fatal(err)
}

// Get hash field-values
userInfo, err := client.HGetAll(ctx, "user:1").Result()
if err!=nil {
  log.Fatal(err)
}
fmt.Println("User Info:", userInfo)

```

### Handling Errors with redis.Nil

- go-redis exports the `redis.Nil` error and returns it whenever Redis Server responds with (`nil`). You can use redis-cli to check what response Redis returns.

```go
val, err := client.Get(ctx, "key").Result()

switch {
  case err == redis.Nil:      // `redis.Nil` to distinguish an empty string reply and a nil reply (key does not exist)
    fmt.Println("Key does not exists")
  case err != nil:
    fmt.Println("Get failed", err)
  case val == "":
    fmt.Println("value is empty")
}

// GET is not the only command that returns nil reply, for example, **"BLPOP"** and **"ZSCORE"** can also return `redis.Nil`.

### Advanced Operations:

```
