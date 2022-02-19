package model

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Key       string
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
