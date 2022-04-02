package request

// Node request model
type Node struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	SSHKey   string `json:"ssh_key"`
}
