package main_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	moq "olibaa/go-moq"
	"testing"
)

func TestGetUserByID(t *testing.T) {
	// GetFuncでMockを作る。Get Methodで取得できるようになる。
	mockRepo := &moq.UserRepositoryInterfaceMock{
		GetFunc: func(ctx context.Context, id string) (*moq.User, error) {
			if id == "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f" {
				return &moq.User{ID: "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f", Name: "Alice"}, nil
			}
			return nil, errors.New("user not found")
		},
	}

	user, err := mockRepo.Get(context.Background(), "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f")
	assert.NoError(t, err)
	assert.Equal(t, "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f", user.ID)
	assert.Equal(t, "Alice", user.Name)

	_, err = mockRepo.Get(context.Background(), "test-uuid")
	assert.Error(t, err)
}
