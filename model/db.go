package model

import (
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func init() {
	loadInstance := func() {
		log.Info("init", log.String("info", "started to load model.DB"))
		for {
			instance, err := gorm.Open(mysql.Open(config.Config.DSN))
			if err != nil {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			DB = instance
			break
		}
		log.Info("init", log.String("info", "model.DB loaded"))
	}

	loadInstance()
	if err := DB.AutoMigrate(
		&User{},
		&Task{},
		&Run{},
		&Node{},
	); err != nil {
		log.Error("model", log.FieldErr(err))
	}

	go func() {
		w := config.GenerateWatcher()
		for {
			<-w
			loadInstance()
		}
	}()

}
