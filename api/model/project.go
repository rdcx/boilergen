package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `yaml:"name" json:"name"`

	User   *User `json:"user"`
	UserID uint  `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
