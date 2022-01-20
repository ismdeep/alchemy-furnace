package model

import "time"

// Run model
type Run struct {
	ID        uint   `gorm:"primary_key"`
	TaskID    string `gorm:"type:varchar(255);not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	ExitCode  int    `gorm:"type:tinyint;not null"`
	Content   string `gorm:"type:longtext;not null"`
	CreatedAt time.Time
}
