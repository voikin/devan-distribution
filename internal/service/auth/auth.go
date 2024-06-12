package auth

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/voikin/devan-distribution/internal/config"
	"github.com/voikin/devan-distribution/internal/entity"
)

const (
	tokenTTL = 24 * time.Hour
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

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int      `json:"user_id"`
	Roles  []string `json:"roles"`
}

func (s *Service) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(ctx, user)
}

func (s *Service) GenerateToken(ctx context.Context, username, password string) (string, string, error) {
	user, err := s.repo.GetUser(ctx, username, s.generatePasswordHash(password))
	if err != nil {
		return "", "", err
	}

	// Создание access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
		user.Roles,
	})

	accessString, err := accessToken.SignedString([]byte(s.cfg.SigningKey))
	if err != nil {
		return "", "", err
	}

	// Создание refresh токена
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)), // Refresh токен действителен в течение недели
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
		user.Roles,
	})

	refreshString, err := refreshToken.SignedString([]byte(s.cfg.SigningKey))
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

func (s *Service) ParseToken(accessToken string) (int, []string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.SigningKey), nil
	})
	if err != nil {
		return 0, nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.Roles, nil
}

func (s *Service) RefreshToken(refreshToken string) (string, error) {
	// Разбор refresh токена
	token, err := jwt.ParseWithClaims(refreshToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	// Создание нового access токена
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		claims.UserId,
		claims.Roles,
	})

	accessString, err := accessToken.SignedString([]byte(s.cfg.SigningKey))
	if err != nil {
		return "", err
	}

	return accessString, nil
}

func (s *Service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.Salt)))
}
