package dto

// Snapshot holds a single Wayback Machine snapshot
type Snapshot struct {
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
}
