package user

import "context"

type Repo interface {
	CreateUser(ctx context.Context) error
}

type Service struct {
	repo Repo
}

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}
