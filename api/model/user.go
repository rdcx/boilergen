package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`

	Projects []Project `json:"projects"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
