package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/voikin/devan-distribution/internal/errs"
)

func (uc *UseCase) GenerateToken(ctx context.Context, username, password string) (string, string, error) {
	user, err := uc.userService.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}

	if uc.generatePasswordHash(password) != user.Password {
		return "", "", errs.NewErrorIncorrectPassword()
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.ID,
		user.Role.ID,
	})

	accessString, err := accessToken.SignedString([]byte(uc.cfg.JWT.SigningKey))
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

func (uc *UseCase) ParseToken(accessToken string) (int64, int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(uc.cfg.JWT.SigningKey), nil
	})
	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, claims.RoleID, nil
}

func (uc *UseCase) RefreshToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(uc.cfg.JWT.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		claims.UserID,
		claims.RoleID,
	})

	accessString, err := accessToken.SignedString([]byte(uc.cfg.JWT.SigningKey))
	if err != nil {
		return "", err
	}

	return accessString, nil
}

func (uc *UseCase) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(uc.cfg.JWT.Salt)))
}
