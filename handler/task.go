package handler

import (
	"errors"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/response"
	"github.com/ismdeep/log"
	"time"
)

type taskHandler struct {
}

var Task = &taskHandler{}

func (receiver *taskHandler) Create(req *request.Task) (uint, error) {
	if req == nil {
		return 0, errors.New("req is nil")
	}

	item := &model.Task{
		Name:        req.Name,
		RunOn:       req.RunOn,
		BashContent: req.BashContent,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
	task.UpdatedAt = time.Now()
	if err := model.DB.Save(task).Error; err != nil {
		return err
	}

	return nil
}

// List get task list
func (receiver *taskHandler) List() []response.Task {
	tasks := make([]*model.Task, 0)
	if err := model.DB.Find(&tasks).Error; err != nil {
		return nil
	}

	results := make([]response.Task, 0)
	for _, task := range tasks {

		// 获取触发器列表
		var _triggers []model.Trigger
		model.DB.Where("task_id=?", task.ID).Find(&_triggers)
		triggers := make([]response.Trigger, 0)
		for _, _trigger := range _triggers {
			t := response.Trigger{
				ID:          _trigger.ID,
				Name:        _trigger.Name,
				Cron:        _trigger.Cron,
				Environment: _trigger.Environment,
			}

			// get last run info of trigger
			var _runs []model.Run
			if err := model.DB.Where("task_id=? AND trigger_id=?", task.ID, _trigger.ID).Order("id desc").Limit(5).Find(&_runs).Error; err != nil {
				log.Error("get task list", log.FieldErr(err))
			}
			for _, v := range _runs {
				t.RecentRuns = append(t.RecentRuns, response.Run{
					ID:          v.ID,
					Name:        v.Name,
					TriggerName: v.TriggerName,
					Status:      v.Status,
					ExitCode:    v.ExitCode,
					Logs:        nil,
					CreatedAt:   v.CreatedAt,
					StartTime:   v.StartTime,
					EndTime:     v.EndTime,
				})
			}
			triggers = append(triggers, t)
		}

		result := response.Task{
			ID:       task.ID,
			Name:     task.Name,
			Bash:     task.BashContent,
			Triggers: triggers,
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
