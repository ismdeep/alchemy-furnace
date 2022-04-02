package request

type Task struct {
	Name        string `json:"name"`
	RunOn       string `json:"run_on"`
	BashContent string `json:"bash_content"`
	Description string `json:"description"`
}
