package model

import "time"

// Run model
type Run struct {
	ID          uint `gorm:"primary_key"`
	TaskID      uint
	Task        Task
	TriggerType int    `gorm:"type:tinyint;not null;default 0;"` // 触发类型， 0定时任务自动触发 1手动执行
	Name        string `gorm:"type:varchar(255);not null"`
	Status      int    `gorm:"type:tinyint;not null"` // 状态， 0Pending 1Running 2Done 3Abort
	ExitCode    int    `gorm:"type:tinyint;not null"`
	Content     string `gorm:"type:longtext;not null"`
	CreatedAt   time.Time
}

const RunEnumsTriggerTypeCronAuto = 0 // 定时任务自动触发
const RunEnumsTriggerTypeManual = 1   // 手动执行

const RunEnumsStatusPending = 0 // Pending
const RunEnumsStatusRunning = 1 // Running
const RunEnumsStatusDone = 2    // Done
const RunEnumsStatusAbort = 3   // Abort
