package model

import "time"

// Run model
type Run struct {
	ID          uint `gorm:"primary_key"`
	TaskID      uint
	Task        Task
	TriggerType int    `gorm:"type:tinyint;not null;default 0;"` // 触发类型， 0定时任务自动触发 1手动执行
	Name        string `gorm:"type:varchar(255);not null"`
	ExitCode    int    `gorm:"type:tinyint;not null"`
	Content     string `gorm:"type:longtext;not null"`
	CreatedAt   time.Time
}
