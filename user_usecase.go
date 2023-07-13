package main

import "context"

//go:generate moq -out mockuser_usecase.go . UserInputPort

type UserInputPort interface {
	Execute(context.Context, string) (*User, error)
}

type UserInteractor struct {
	UserRepositoryInterface UserRepositoryInterface
}

func NewUserUsecase(
	uri UserRepositoryInterface,
) UserInputPort {
	return &UserInteractor{
		UserRepositoryInterface: uri,
	}
}

func (i *UserInteractor) Execute(ctx context.Context, id string) (*User, error) {
	user, err := i.UserRepositoryInterface.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
