package response

import (
	"github.com/ismdeep/alchemy-furnace/executor"
	"time"
)

type Run struct {
	ID               uint              `json:"id"`
	Name             string            `json:"name"`
	TriggerName      string            `json:"trigger_name"`
	Status           int               `json:"status"`
	ExitCode         int               `json:"exit_code"`
	Logs             []executor.ExeLog `json:"logs"`
	CreatedAt        time.Time         `json:"created_at"`
	StartTime        time.Time         `json:"start_time"`
	EndTime          time.Time         `json:"end_time"`
	TimeElapseSecond uint              `json:"time_elapse_second"`
}
