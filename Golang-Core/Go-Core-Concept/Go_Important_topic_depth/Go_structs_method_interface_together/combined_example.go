package user

import (
	"context"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID    string
	Name  string
	Email string
}

func (u User) Validate() error {
	if u.ID == "" {
		return fmt.Errorf("id is required")
	}
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

type Repository interface {
	Save(ctx context.Context, user User) error
	GetByID(ctx context.Context, id string) (User, error)
	List(ctx context.Context) ([]User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, user User) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("validate user: %w", err)
	}
	if err := s.repo.Save(ctx, user); err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	return nil
}

func (s *Service) GetNameByDomain(ctx context.Context, domain string) (map[string][]string, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}
	result := make(map[string][]string)

	for _, user := range users {
		result[domain] = append(result[domain], user.Name)
	}
	return result, nil
}
