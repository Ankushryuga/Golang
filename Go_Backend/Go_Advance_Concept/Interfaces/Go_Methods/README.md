# Methods Definition:
```go
type User struct {
  ID          int
  Name        string
  Email       string
  CreatedBy   time.Time
}

type Database struct{
  Connection  string
}
``` 
In Go, **methods** are functions that are associated with a specific type. When we define a **struct** (a modeled type), we often need to define its behaviour as well. This is done by attaching functions, called **methods**, to the struct.


By associating methods with a struct, we define the actions that the struct can perform, making it more than just a collection of fields. This allows us to encapsulate both **data** and **behavior** withing a type.


Example: 

```go
// Method to validate a User
func (u *User) Validate() error {
  if u.Name == "" {
    return errors.New("name is required")
  }
  if !strings.Contains(u.Email, "@") {
    return errors.New("invalid email")
  }
  return nil
}

// Method to connect to a Database
func (db *Database) Connect() error {
  // Simulate connection logic
  fmt.Println("Connected to database:", db.Connection)
  return nil
}
```

