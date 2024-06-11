package controller

import "context"

type UseCase interface {
	CreateUser(ctx context.Context) error
}

type Controller struct {
	usecase UseCase
}

func New(usecase UseCase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}
