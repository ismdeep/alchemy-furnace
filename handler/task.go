package handler

import (
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/response"
)

type taskHandler struct {
}

var Task = &taskHandler{}

// List get task list
func (receiver *taskHandler) List() []*response.Task {
	results := make([]*response.Task, 0)
	for _, task := range config.Config.Tasks {
		results = append(results, &response.Task{
			ID:   task.ID,
			Name: task.Name,
			Bash: task.Bash,
			Cron: task.Cron,
		})
	}
	return results
}
