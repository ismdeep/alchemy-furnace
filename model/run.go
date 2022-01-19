package model

import "gorm.io/gorm"

// Run model
type Run struct {
	gorm.Model
	TaskID   string `gorm:"type:varchar(255)"`
	Name     string `gorm:"type:varchar(255)"`
	ExitCode int    `gorm:"type:tinyint"`
	Content  string `gorm:"type:longtext"`
}
