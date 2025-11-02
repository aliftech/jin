package dto

import "github.com/aliftech/jin/internal/core"

type ServerInfoResponse struct {
	URL            string            `json:"url"`
	StatusCode     int               `json:"status_code"`
	Server         string            `json:"server"`
	PoweredBy      string            `json:"powered_by"`
	ContentType    string            `json:"content_type"`
	TLSVersion     string            `json:"tls_version,omitempty"`
	TLSCipherSuite string            `json:"tls_cipher_suite,omitempty"`
	Headers        map[string]string `json:"headers"`
	Error          string            `json:"error,omitempty"`
}

func FromDomain(info *core.ServerInfo) *ServerInfoResponse {
	return &ServerInfoResponse{
		URL:            info.URL,
		StatusCode:     info.StatusCode,
		Server:         info.Server,
		PoweredBy:      info.PoweredBy,
		ContentType:    info.ContentType,
		TLSVersion:     info.TLSVersion,
		TLSCipherSuite: info.TLSCipherSuite,
		Headers:        info.Headers,
	}
}
