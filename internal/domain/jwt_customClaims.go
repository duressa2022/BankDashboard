package domain

import "github.com/golang-jwt/jwt/v5"

// type for working with jwt custom claims

type Claims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.RegisteredClaims
}

// type for working with jwt custom refresh Claims
type RefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
