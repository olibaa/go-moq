package main

import (
	"context"

	"github.com/google/uuid"
)

//go:generate moq -out mockuser_repository.go . UserRepositoryInterface

type UserRepositoryInterface interface {
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, name string) (*User, error)
}

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (r *UserRepository) Get(ctx context.Context, id string) (*User, error) {
	user := &User{
		ID: id,
	}
	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, name string) (*User, error) {
	id := uuid.New()
	user := &User{
		ID:   id.String(),
		Name: name,
	}
	return user, nil
}
