package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// GenerateJWT generates a JWT token
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 3 days
	})

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return "Bearer " + signedToken, nil
}

// ComparePassword compares the password and hash
func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ParseJWT parses the JWT token
func ParseJWT(tokenString string) (string, error) {
	// Remove the "Bearer " prefix
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract the username
	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("username claim is not a string")
	}

	return username, nil
}
