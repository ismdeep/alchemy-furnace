package model

import (
	"time"
)

type Trigger struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	TaskID      uint
	Task        Task
	Cron        string `gorm:"type:varchar(255);not null"`
	Environment string `gorm:"type:longtext"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
