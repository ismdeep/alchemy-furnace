package model

type AccessToken struct {
	ID     uint `gorm:"primary_key"`
	UserID uint
	User   User
	Token  string
}
