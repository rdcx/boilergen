package util

import (
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[RandomInt(len(letters))]
	}
	return string(result)
}

func RandomStringWithPrefix(prefix string) string {
	return prefix + RandomString(10)
}

func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ParseClaims(token string) (jwt.MapClaims, error) {

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {

		if token.Method.Alg() != "HS256" {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
