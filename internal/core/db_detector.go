// internal/core/db_detector.go
package core

import "context"

type DatabaseDetector interface {
	Detect(ctx context.Context, url string) (*DatabaseInfo, error)
}

type DatabaseInfo struct {
	URL        string   `json:"url"`
	Detected   bool     `json:"detected"`
	Database   string   `json:"database"`
	Evidence   []string `json:"evidence"`
	Confidence string   `json:"confidence"` // "high", "medium", "low"
}
