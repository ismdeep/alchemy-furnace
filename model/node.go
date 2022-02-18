package model

import (
	"gorm.io/gorm"
	"time"
)

// Node model
type Node struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	User      User
	Name      string `gorm:"type:varchar(255)"`
	Host      string `gorm:"type:varchar(255)"`
	Port      int    `gorm:"type:int"`
	Username  string `gorm:"type:varchar(255)"`
	SSHKey    string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
