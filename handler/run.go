package handler

import (
	"encoding/json"
	"github.com/ismdeep/alchemy-furnace/executor"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/response"
)

type runHandler struct{}

var Run = &runHandler{}

// List tasks
func (receiver *runHandler) List(taskID string, page int, size int) ([]*response.Run, int64, error) {
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
			ExitCode:  item.ExitCode,
			Logs:      []executor.ExeLog{},
			CreatedAt: item.CreatedAt,
		})
	}

	return results, total, nil
}

func (receiver *runHandler) Detail(taskID string, runID uint) (*response.Run, error) {
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
