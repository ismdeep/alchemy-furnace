package model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(255);unique"`
	Digest    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type userStore struct {
}

var UserStore = &userStore{}

func (receiver *userStore) UserExists(username string) (bool, error) {
	var count int64
	if err := DB.Model(&User{}).Where("username=?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (receiver *userStore) GetUser(username string) (*User, error) {
	u := &User{}
	if err := DB.Where("username=?", username).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
