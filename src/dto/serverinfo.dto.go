package dto

// ServerInfo holds the collected server information
type ServerInfo struct {
	WebServer    string            `json:"web_server"`
	OS           string            `json:"os"`
	Database     string            `json:"database"`
	PoweredBy    string            `json:"powered_by"`
	Framework    string            `json:"framework"`
	OpenPorts    []string          `json:"open_ports"`
	OtherHeaders map[string]string `json:"other_headers"`
}
