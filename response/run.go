package response

import (
	"github.com/ismdeep/alchemy-furnace/executor"
	"time"
)

type Run struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	ExitCode  int               `json:"exit_code"`
	Logs      []executor.ExeLog `json:"logs"`
	CreatedAt time.Time         `json:"created_at"`
}
