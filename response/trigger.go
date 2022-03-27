package response

type Trigger struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Cron        string `json:"cron"`
	Environment string `json:"environment"`
	RecentRuns  []Run  `json:"recent_runs"`
}
