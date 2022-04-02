package model

import "time"

// Run model
type Run struct {
	ID          uint `gorm:"primary_key"`
	ExecutorID  string
	TaskID      uint
	Task        Task
	TriggerID   uint
	Trigger     Trigger
	TriggerName string
	TriggerType int       `gorm:"type:tinyint;not null;default 0;"` // 触发类型， 0定时任务自动触发 1手动执行
	Name        string    `gorm:"type:varchar(255);not null"`
	Status      int       `gorm:"type:tinyint;not null"` // 状态， 0Pending 1Running 2Done 3Abort
	ExitCode    int       `gorm:"type:tinyint;not null"`
	CmdLog      string    `gorm:"type:longtext"`
	Content     string    `gorm:"type:longtext;not null"`
	StartTime   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	EndTime     time.Time `gorm:"type:timestamp;default:current_timestamp"`
	CreatedAt   time.Time
}

const RunEnumsTriggerTypeCronAuto = 0 // 定时任务自动触发
const RunEnumsTriggerTypeManual = 1   // 手动执行

const RunEnumsStatusPending = 0 // Pending
const RunEnumsStatusRunning = 1 // Running
const RunEnumsStatusDone = 2    // Done
const RunEnumsStatusAbort = 3   // Abort
