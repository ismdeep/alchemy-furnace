package model

type Task struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(255);not null"`
	UserID      uint
	User        User
	Cron        string `gorm:"type:varchar(255);not null"`
	BashContent string `gorm:"type:longtext;not null"`
	Description string `gorm:"type:text;not null"`
}
