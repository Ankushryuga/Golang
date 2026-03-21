# Structs + Methods + Interfaces together

```go
type User struct{
    ID      string
    Name    string
    Email   string
}

func (u User) Validate() error {
    if u.ID == ""{
        return fmt.Errorf("Id is required")
    }
    if u.Email == ""{
        return fmt.Errorf("email is required")
    }
    return nil
}

type UserRepository interface {
    Save(ctx context.Context, user User) error
    GetByID(ctx context.Context, id string) (User, error)
}

type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user User) error {
    if err := user.Validate(); err != nil {
        return err
    }
    return s.repo.Save(ctx, user)
}
```

This shows:

- struct as data
- method for domain behavior
- interface for dependency boundry
- concrete service struct

This is much closer to real backend Go.
