package main_test

import (
	"context"
	"errors"
	moq "olibaa/go-moq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserInteractor_Execute(t *testing.T) {
	// GetFuncでMockを作る。Get Methodで取得できるようになる。
	mockRepo := &moq.UserRepositoryInterfaceMock{
		GetFunc: func(ctx context.Context, id string) (*moq.User, error) {
			if id == "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f" {
				return &moq.User{ID: "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f", Name: "Alice"}, nil
			}
			return nil, errors.New("user not found")
		},
	}

	usecase := moq.NewUserUsecase(mockRepo)

	user, err := usecase.Execute(context.Background(), "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f")
	assert.NoError(t, err)
	assert.Equal(t, "1f2c3625-9bc2-5e1b-b808-cf96a85aee1f", user.ID)
	assert.Equal(t, "Alice", user.Name)

	_, err = usecase.Execute(context.Background(), "test-uuid")
	assert.Error(t, err)
}
