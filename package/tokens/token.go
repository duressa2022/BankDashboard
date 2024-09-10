package tokens

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"working.com/bank_dash/internal/domain"
)

// method for creating access token
func CreateAccessToken(user *domain.User, secret string, exp int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(exp) * time.Second)
	claims := domain.Claims{
		Username: user.UserName,
		ID:       user.Id.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// method for creating refresh token
func CreateRefreshToken(user *domain.User, secret string, exp int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(exp) * time.Second)
	claims := domain.RefreshClaims{
		ID: user.Id.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// method for verfying access and refresh token
func VerifyToken(token string, secret string) (bool, error) {

	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}
	if !tokenString.Valid {
		return false, errors.New("invalid token")
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

// method for getting the claims
func GetUserClaims(token string, secret string) (jwt.MapClaims, error) {
	tokenString, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("your method for signing token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tokenString.Claims.(jwt.MapClaims)
	if !ok && !tokenString.Valid {
		return nil, errors.New("invalid tokn for working")
	}
	return claims, nil
}
