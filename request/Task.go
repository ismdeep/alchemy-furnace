package request

type Task struct {
	Name        string `json:"name"`
	Cron        string `json:"cron"`
	RunOn       string `json:"run_on"`
	BashContent string `json:"bash_content"`
	Description string `json:"description"`
}
