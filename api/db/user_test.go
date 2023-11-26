package db_test

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUse(t *testing.T) {
	db := testDB()

	email := gofakeit.Email()
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, false, false, 10)

	user, err := db.CreateUser(email, username, password)

	if err != nil {
		t.Error(err)
	}

	if user.Email != email {
		t.Errorf("Expected email to be %s, got %s", email, user.Email)
	}

	if user.Username != username {
		t.Errorf("Expected username to be %s, got %s", username, user.Username)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		t.Error(err)
	}
}

func TestLogin(t *testing.T) {

	os.Setenv("JWT_SECRET", "secret")

	db := testDB()

	email := gofakeit.Email()
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, false, false, 10)

	u, err := db.CreateUser(email, username, password)
	if err != nil {
		t.Error(err)
	}

	token, err := db.Login(username, password)

	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("Expected token to not be empty")
	}

	claims := jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {

		if token.Method.Alg() != "HS256" {
			t.Errorf("Expected alg to be HS256, got %s", token.Method.Alg())
		}

		return []byte("secret"), nil
	})

	if err != nil {
		t.Error(err)
	}

	if claims["exp"] == nil {
		t.Error("Expected exp to not be nil")
	}

	if claims["user_id"] == nil {
		t.Error("Expected user_id to not be nil")
	}

	v, ok := claims["user_id"].(float64)
	if !ok {
		t.Errorf("Expected user_id to be float64, got %T", claims["user_id"])
	}

	if uint(v) != u.ID {
		t.Errorf("Expected user_id to be %d", u.ID)
	}
}
