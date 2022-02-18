package model

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(255);not null"`
	UserID      uint
	User        User
	RunOn       string `gorm:"type:varchar(512);not null"`
	BashContent string `gorm:"type:longtext;not null"`
	Description string `gorm:"type:text;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
