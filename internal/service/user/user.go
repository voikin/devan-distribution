package user

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/voikin/devan-distribution/internal/entity"
	"github.com/voikin/devan-distribution/internal/errs"
)

type Repo interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserByID(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (bool, error)
	DeleteUser(ctx context.Context, userId int64) (bool, error)
}

type Service struct {
	repo Repo
}

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(ctx, user)
}

func (s *Service) ValidateUserByCreds(ctx context.Context, userName string, password string) (*entity.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, userName)
	if err != nil {
		return nil, err
	}

	if generatePasswordHash(password) != user.Password {
		return nil, errs.NewErrorIncorrectPassword()
	}

	return user, nil
}

func (s *Service) GenerateToken(ctx context.Context, user entity.User) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
		user.Role.ID,
	})

	accessString, err := accessToken.SignedString([]byte(s.cfg.JWT.SigningKey))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)), // Refresh токен действителен в течение недели
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
		user.Role.ID,
	})

	refreshString, err := refreshToken.SignedString([]byte(s.cfg.SigningKey))
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.Salt)))
}
