// internal/core/server_scanner.go
package core

import "context"

// ServerScanner is the port for gathering web server info
type ServerScanner interface {
	Scan(ctx context.Context, url string) (*ServerInfo, error)
}

// ServerInfo holds the gathered data
type ServerInfo struct {
	URL            string            `json:"url"`
	StatusCode     int               `json:"status_code"`
	Server         string            `json:"server"`
	PoweredBy      string            `json:"powered_by"`
	ContentType    string            `json:"content_type"`
	TLSVersion     string            `json:"tls_version,omitempty"`
	TLSCipherSuite string            `json:"tls_cipher_suite,omitempty"`
	Headers        map[string]string `json:"headers"`
}
