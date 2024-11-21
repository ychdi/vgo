package auth

import "github.com/golang-jwt/jwt/v5"

// AdminClaims AdminClaims
type AdminClaims struct {
	UserID uint64   `json:"user_id"`
	Role   []string `json:"role"`
	Super  int      `json:"super"`
	jwt.RegisteredClaims
}

// UserClaims UserClaims
type UserClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}
