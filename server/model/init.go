package model

import (
	"fmt"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	instance, err := gorm.Open(sqlite.Open(fmt.Sprintf("%v/data/data.db", config.WorkDir)))
	if err != nil {
		panic(err)
	}
	DB = instance
	if err := DB.AutoMigrate(
		&User{},
		&Task{},
		&Run{},
		&Node{},
		&Trigger{},
		&Token{},
	); err != nil {
		log.Error("model", log.FieldErr(err))
	}
}
