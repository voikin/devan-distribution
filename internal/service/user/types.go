package user

import (
	"github.com/golang-jwt/jwt/v5"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}
