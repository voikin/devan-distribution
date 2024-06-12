package auth

import (
	"context"

	"github.com/voikin/devan-distribution/internal/config"
	"github.com/voikin/devan-distribution/internal/entity"
)

type Repo interface {
	CreateUser(ctx context.Context, input entity.User) (int64, error)
	GetUser(ctx context.Context, username, password string) (*entity.User, error)
}

type Service struct {
	repo Repo
	cfg  *config.JWTConfig
}

func New(repo Repo, cfg *config.JWTConfig) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}
