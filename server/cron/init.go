package cron

import (
	"fmt"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/log"
	"github.com/robfig/cron/v3"
)

type Job struct {
	TaskID    uint
	TriggerID uint
}

func (receiver Job) Run() {
	if err := handler.Run.Start(receiver.TaskID, receiver.TriggerID); err != nil {
		log.Error("Cron", log.Any("trigger_id", receiver.TriggerID), log.FieldErr(err))
	}
}

var cronIDs map[uint]cron.EntryID

func Push(trigger model.Trigger) {
	// 1. 检查是否有cronID
	if id, ok := cronIDs[trigger.ID]; ok {
		fmt.Println("remove old cron job")
		c.Remove(id)
	}

	if _, err := cron.ParseStandard(trigger.Cron); err != nil {
		return
	}

	id, err := c.AddJob(trigger.Cron, Job{TaskID: trigger.TaskID, TriggerID: trigger.ID})
	if err != nil {
		log.Error("cron", log.FieldErr(err))
	}
	cronIDs[trigger.ID] = id
}

var c *cron.Cron

func init() {
	cronIDs = make(map[uint]cron.EntryID)
	c = cron.New()

	var triggers []model.Trigger
	if err := model.DB.Find(&triggers).Error; err != nil {
		panic(err)
	}

	for _, trigger := range triggers {
		Push(trigger)
	}
	c.Start()

	go func() {
		for {
			t := <-handler.TriggerChan
			Push(t)
		}
	}()

}
