package model

import (
	"time"

	"gorm.io/gorm"
)

type Option struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	ParentID    uint    `json:"parent"`
	Parent      *Option `gorm:"foreignKey:ParentID"`
	KeyName     string  `yaml:"key_name" json:"key_name"`
	DisplayName string  `yaml:"name" json:"name"`
	Description string  `yaml:"description" json:"description"`

	Values []Option `json:"values" gorm:"foreignKey:ParentID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
