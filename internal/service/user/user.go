package user

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/voikin/devan-distribution/internal/DTO"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/voikin/devan-distribution/internal/config"
	"github.com/voikin/devan-distribution/internal/entity"
	"github.com/voikin/devan-distribution/internal/errs"
)

type Repo interface {
	CreateUser(ctx context.Context, user DTO.CreateUser) (int64, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserByID(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (bool, error)
	DeleteUser(ctx context.Context, userId int64) (bool, error)
	GetRoles(ctx context.Context) ([]DTO.Role, error)
}

type Service struct {
	repo   Repo
	config config.JWTConfig
}

func New(repo Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, user DTO.CreateUser) (int64, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(ctx, user)
}

func (s *Service) GetUserByUsernameWithPassCheck(ctx context.Context, userName string, password string) (*entity.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, userName)
	if err != nil {
		return nil, err
	}

	if s.generatePasswordHash(password) != user.Password {
		return nil, errs.NewErrorIncorrectPassword()
	}

	return user, nil
}

func (s *Service) GetUserByTokenWithCheck(ctx context.Context, token string) (*entity.User, error) {
	userId, err := s.verifyToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetRoles(ctx context.Context) ([]DTO.Role, error) {
	roles, err := s.repo.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *Service) GenerateTokens(ctx context.Context, user *entity.User) (string, string, error) {
	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *Service) generateAccessToken(user *entity.User) (string, error) {
	accessToken, err := s.generateToken(user, s.config.AccessTokenTTL)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *Service) generateRefreshToken(user *entity.User) (string, error) {
	refreshToken, err := s.generateToken(user, s.config.RefreshTokenTTL)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (s *Service) verifyToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.config.SigningKey), nil
	})
	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func (s *Service) generateToken(user *entity.User, TTL time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
	})

	tokenString, err := token.SignedString([]byte(s.config.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.config.Salt)))
}
