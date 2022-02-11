package handler

import (
	"encoding/json"
	"errors"
	"github.com/ismdeep/alchemy-furnace/executor"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/response"
	"github.com/ismdeep/alchemy-furnace/runner"
	"github.com/ismdeep/log"
	"time"
)

type runHandler struct{}

var Run = &runHandler{}

// List tasks
func (receiver *runHandler) List(taskID uint, page int, size int) ([]*response.Run, int64, error) {
	items := make([]*model.Run, 0)
	var total int64
	conn := model.DB.Model(&items).Where("task_id=?", taskID)
	if err := conn.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := conn.Order("id desc").Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	results := make([]*response.Run, 0)
	for _, item := range items {
		results = append(results, &response.Run{
			ID:        item.ID,
			Name:      item.Name,
			Status:    item.Status,
			ExitCode:  item.ExitCode,
			Logs:      []executor.ExeLog{},
			CreatedAt: item.CreatedAt,
		})
	}

	return results, total, nil
}

func (receiver *runHandler) Detail(taskID uint, runID uint) (*response.Run, error) {
	item := &model.Run{}
	if err := model.DB.Where("id=? AND task_id=?", runID, taskID).First(item).Error; err != nil {
		return nil, err
	}

	logs := make([]executor.ExeLog, 0)
	_ = json.Unmarshal([]byte(item.Content), &logs)

	return &response.Run{
		ID:        item.ID,
		Name:      item.Name,
		ExitCode:  item.ExitCode,
		Logs:      logs,
		CreatedAt: item.CreatedAt,
	}, nil
}

func (receiver *runHandler) Start(taskID uint) error {
	// 1. check taskID
	var cnt int64
	if err := model.DB.Model(&model.Task{}).Where("id=?", taskID).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt <= 0 {
		return errors.New("task is not exists")
	}

	// 2. write info
	executorID := executor.GenerateExecutor()
	run := &model.Run{
		ExecutorID:  executorID,
		TaskID:      taskID,
		TriggerType: 0,
		Name:        "",
		ExitCode:    0,
		Content:     "",
	}
	if err := model.DB.Create(run).Error; err != nil {
		return err
	}

	// 3. start to run task
	go func(runID uint, executorID string) {
		run := &model.Run{}
		model.DB.Where("id=?", runID).First(run)
		startTime := time.Now()
		run.StartTime = startTime
		run.EndTime = time.Now()
		run.Status = model.RunEnumsStatusRunning
		model.DB.Save(run)

		exitCode, err := runner.Run(runID, executorID)
		if err != nil {
			log.Error("Run", log.FieldErr(err))
		}
		run.Status = model.RunEnumsStatusDone
		run.ExitCode = exitCode
		run.EndTime = time.Now()
		run.CmdLog, err = executor.DumpLog(executorID)
		if err != nil {
			log.Error("Run", log.FieldErr(err))
		}
		model.DB.Save(run)

	}(run.ID, executorID)

	return nil
}
