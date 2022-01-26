package model

import (
	"github.com/ismdeep/alchemy-furnace/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func init() {
	for {
		instance, err := gorm.Open(mysql.Open(config.Config.DB.DSN), &gorm.Config{})
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		DB = instance
		break
	}

	if err := DB.AutoMigrate(
		&User{},
		&Task{},
		&Run{},
	); err != nil {
		panic(err)
	}
}
