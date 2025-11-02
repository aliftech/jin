package scrape

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aliftech/jin/internal/core"
)

type HTTPServerScanner struct {
	client *http.Client
}

func NewHTTPServerScanner(timeout time.Duration) *HTTPServerScanner {
	return &HTTPServerScanner{
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, // For scanning only
				},
			},
		},
	}
}

func (s *HTTPServerScanner) Scan(ctx context.Context, rawURL string) (*core.ServerInfo, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}

	req, err := http.NewRequestWithContext(ctx, "GET", parsed.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("request setup failed: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	info := &core.ServerInfo{
		URL:         parsed.String(),
		StatusCode:  resp.StatusCode,
		Server:      strings.TrimSpace(resp.Header.Get("Server")),
		PoweredBy:   strings.TrimSpace(resp.Header.Get("X-Powered-By")),
		ContentType: strings.TrimSpace(resp.Header.Get("Content-Type")),
		Headers:     make(map[string]string),
	}

	for name, values := range resp.Header {
		key := http.CanonicalHeaderKey(name)
		info.Headers[key] = strings.Join(values, ", ")
	}

	if resp.TLS != nil {
		info.TLSVersion = tlsVersionName(resp.TLS.Version)
		info.TLSCipherSuite = tls.CipherSuiteName(resp.TLS.CipherSuite)
	}

	return info, nil
}

func tlsVersionName(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("Unknown (0x%04x)", v)
	}
}
