package handler

import (
	"errors"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/response"
)

type taskHandler struct {
}

var Task = &taskHandler{}

func (receiver *taskHandler) Create(userID uint, req *request.Task) (uint, error) {
	if req == nil {
		return 0, errors.New("req is nil")
	}

	item := &model.Task{
		Name:        req.Name,
		UserID:      userID,
		Cron:        req.Cron,
		BashContent: req.BashContent,
		Description: req.Description,
	}

	model.DB.Create(item)

	return item.ID, nil
}

// List get task list
func (receiver *taskHandler) List(userID uint) []*response.Task {
	tasks := make([]*model.Task, 0)
	if err := model.DB.Preload("User").Where("user_id=?", userID).Find(&tasks).Error; err != nil {
		return nil
	}

	results := make([]*response.Task, 0)
	for _, task := range tasks {
		results = append(results, &response.Task{
			ID:   task.ID,
			Name: task.Name,
			Bash: task.BashContent,
			Cron: task.Cron,
		})
	}
	return results
}
