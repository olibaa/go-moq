package main

import (
	"context"
	"fmt"
)

//go:generate moq -out mockuser_repository.go . UserRepositoryInterface

type User struct {
	ID   string
	Name string
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
