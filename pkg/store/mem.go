package store

import (
	"context"

	"github.com/google/uuid"
)

type memStore struct {
	users map[string]User
}

func NewMemStore() Store {
	return &memStore{
		users: make(map[string]User),
	}
}

func (s *memStore) CreateUser(ctx context.Context, name string, email string) (string, error) {
	id := uuid.NewString()
	s.users[id] = User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return id, nil
}

func (s *memStore) GetUser(ctx context.Context, id string) (*User, error) {
	if user, found := s.users[id]; found {
		return &user, nil
	}
	return nil, ErrNotFound
}

func (s *memStore) DeleteUser(ctx context.Context, id string) error {
	if _, found := s.users[id]; found {
		delete(s.users, id)
		return nil
	}
	return ErrNotFound
}

func (s *memStore) ListUsers(ctx context.Context) ([]User, error) {
	users := []User{}
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}
