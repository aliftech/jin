package core

import "context"

type PortScanner interface {
	Scan(ctx context.Context, host string, ports []int) ([]PortInfo, error)
}

type PortInfo struct {
	Port    int    `json:"port"`
	Service string `json:"service"` // e.g., "http", "https"
	Status  string `json:"status"`  // "open" or "closed"
}
