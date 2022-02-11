package response

// Task response schema
type Task struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Bash    string `json:"bash"`
	Cron    string `json:"cron"`
	LastRun *Run   `json:"last_run"`
}
