package store

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("NotFound")
)

type Store interface {
	CreateUser(ctx context.Context, name string, email string) (string, error)
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]User, error)
}

type User struct {
	ID    string
	Name  string
	Email string
}
