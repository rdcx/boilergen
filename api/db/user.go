package db

import (
	"encoding/json"
	"errors"
	"lightban/api/model"
	"lightban/util"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func (db *DB) GetUser(id uint) (*model.User, error) {
	var user model.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) CreateUser(email, username, password string) (*model.User, error) {
	var user model.User

	if err := db.Where("email = ?", email).First(&user).Error; err == nil {
		return nil, ErrUserAlreadyExists
	}

	pass, err := util.HashPassword(password)

	if err != nil {
		return nil, err
	}

	user = model.User{
		Email:    email,
		Username: username,
		Password: pass,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) Login(username, password string) (string, error) {

	var user model.User

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	key := os.Getenv("JWT_SECRET")

	if key == "" {
		panic("JWT_SECRET not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
		"user_id": json.Number(strconv.FormatUint(uint64(user.ID), 10)),
	})

	return token.SignedString([]byte(key))
}
