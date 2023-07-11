package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

//go:generate moq -out mockuser_test.go . UserRepositoryInterface

type User struct {
	ID   string
	Name string
}

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

func main() {
	userRepository := &UserRepository{}

	// Getメソッドの使用例
	user, err := userRepository.Get(context.Background(), "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f")
	if err != nil {
		fmt.Println("Failed to get user:", err)
		return
	}
	fmt.Println("User:", user)
}
