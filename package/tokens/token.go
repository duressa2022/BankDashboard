package tokens

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"working.com/bank_dash/internal/domain"
)

// method fr creating access token
func CreateAccessToken(user domain.UserResponse, secret string, exp int) (string, error) {
	time := time.Now().Add(time.Duration(exp))
	claims := domain.Claims{
		Username: user.UserName,
		ID:       user.Id.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// method for creating refresh token
func CreateRefreshToken(user domain.UserResponse, secret string, exp int) (string, error) {
	time := time.Now().Add(time.Duration(exp))
	claims := domain.RefreshClaims{
		ID: user.Id.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// method for verfying access and refresh token
func VerifyToken(token string, secret string) (bool, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("error of method usage")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, nil
	}
	if !tokenString.Valid {
		return false, nil
	}
	return true, nil
}

// method for getting username from the token
func GetUserName(token string, secret string) (string, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error for method checking")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok && !tokenString.Valid {
		return "", err
	}
	return claims["username"].(string), err

}

// method for getting id from the token
func GetUserId(token string, secret string) (string, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("your method for signing token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok && !tokenString.Valid {
		return "", errors.New("invalid tokn for working")
	}
	return claims["id"].(string), nil
}
