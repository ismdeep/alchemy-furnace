package request

type Task struct {
	Name        string `json:"name"`
	Cron        string `json:"cron"`
	BashContent string `json:"bash_content"`
	Description string `json:"description"`
}
