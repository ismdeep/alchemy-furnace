package response

type Trigger struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Cron        string `json:"cron"`
	Environment string `json:"environment"`
}
