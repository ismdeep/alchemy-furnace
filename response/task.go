package response

// Task response schema
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Bash string `json:"bash"`
	Cron string `json:"cron"`
}
