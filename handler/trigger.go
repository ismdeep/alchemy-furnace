package handler

import (
	"errors"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/response"
)

type triggerHandler struct {
}

var Trigger = &triggerHandler{}

var TriggerChan chan model.Trigger

func init() {
	TriggerChan = make(chan model.Trigger)
}

func (receiver *triggerHandler) Add(userID uint, taskID uint, req *request.Trigger) (uint, error) {
	if req == nil {
		return 0, errors.New("req is nil")
	}

	trigger := model.Trigger{
		Name:        req.Name,
		TaskID:      taskID,
		Cron:        req.Cron,
		Environment: req.Environment,
	}

	if err := model.DB.Create(&trigger).Error; err != nil {
		return 0, err
	}

	TriggerChan <- trigger

	return trigger.ID, nil
}

func (receiver *triggerHandler) Update(userID uint, taskID uint, triggerID uint, req *request.Trigger) error {
	if req == nil {
		return errors.New("req is nil")
	}

	var triggers []model.Trigger
	if err := model.DB.Where("id=?", triggerID).Find(&triggers).Error; err != nil {
		return err
	}
	if len(triggers) <= 0 {
		return errors.New("record not found")
	}
	trigger := triggers[0]
	trigger.Name = req.Name
	trigger.Cron = req.Cron
	trigger.Environment = req.Environment
	if err := model.DB.Save(&trigger).Error; err != nil {
		return err
	}

	TriggerChan <- trigger

	return nil
}

func (receiver *triggerHandler) List(userID uint, taskID uint) ([]response.Trigger, error) {
	var tasks []model.Task
	if err := model.DB.Where("id=?", taskID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	var triggers []model.Trigger
	if err := model.DB.Where("task_id=?", taskID).Find(&triggers).Error; err != nil {
		return nil, err
	}

	var respTriggers []response.Trigger
	for _, trigger := range triggers {
		respTriggers = append(respTriggers, response.Trigger{
			ID:          trigger.ID,
			Name:        trigger.Name,
			Cron:        trigger.Cron,
			Environment: trigger.Environment,
		})
	}
	return respTriggers, nil
}

func (receiver *triggerHandler) Delete(taskID uint, triggerID uint) error {
	var triggers []model.Trigger
	if err := model.DB.Where("id=?", triggerID).Find(&triggers).Error; err != nil {
		return err
	}
	trigger := triggers[0]
	if err := model.DB.Delete(&trigger).Error; err != nil {
		return err
	}

	return nil
}
