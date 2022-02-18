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
		RunOn:       req.RunOn,
		BashContent: req.BashContent,
		Description: req.Description,
	}

	model.DB.Create(item)

	return item.ID, nil
}

func (receiver *taskHandler) Update(taskID uint, req *request.Task) error {
	task := &model.Task{}
	if err := model.DB.Where("id=?", taskID).First(task).Error; err != nil {
		return err
	}

	task.Name = req.Name
	task.BashContent = req.BashContent
	if err := model.DB.Save(task).Error; err != nil {
		return err
	}

	return nil
}

// List get task list
func (receiver *taskHandler) List(userID uint) []*response.Task {
	tasks := make([]*model.Task, 0)
	if err := model.DB.Preload("User").Where("user_id=?", userID).Find(&tasks).Error; err != nil {
		return nil
	}

	results := make([]*response.Task, 0)
	for _, task := range tasks {

		result := &response.Task{
			ID:      task.ID,
			Name:    task.Name,
			Bash:    task.BashContent,
			LastRun: nil,
		}

		// 获取最后一次运行记录
		lastRuns := make([]model.Run, 0)
		model.DB.Where("task_id=?", task.ID).Order("id desc").Find(&lastRuns)
		if len(lastRuns) <= 0 {
			result.LastRun = nil
		} else {
			result.LastRun = &response.Run{
				ID:        lastRuns[0].ID,
				Name:      "",
				Status:    lastRuns[0].Status,
				ExitCode:  lastRuns[0].ExitCode,
				CreatedAt: lastRuns[0].CreatedAt,
				StartTime: lastRuns[0].StartTime,
				EndTime:   lastRuns[0].EndTime,
			}
		}

		results = append(results, result)
	}

	return results
}

func (receiver *taskHandler) Detail(taskID uint) (*response.Task, error) {
	task := &model.Task{}
	if err := model.DB.Where("id=?", taskID).First(task).Error; err != nil {
		return nil, err
	}

	return &response.Task{
		ID:   task.ID,
		Name: task.Name,
		Bash: task.BashContent,
	}, nil
}

func (receiver *taskHandler) Delete(taskID uint) error {
	var tasks []model.Task
	if err := model.DB.Where("id=?", taskID).Find(&tasks).Error; err != nil {
		return err
	}

	if len(tasks) <= 0 {
		return errors.New("task not found")
	}

	task := tasks[0]
	if err := model.DB.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
