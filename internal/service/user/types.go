package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTTL = 24 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	RoleID int64 `json:"role_id"`
	UserID int64 `json:"user_id"`
}
