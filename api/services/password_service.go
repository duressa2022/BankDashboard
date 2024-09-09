package services

import "golang.org/x/crypto/bcrypt"

// method for encrypting the password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// method for comparing password with hashed
func ComparePassword(password string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
