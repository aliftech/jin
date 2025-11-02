// internal/scrape/db_detector.go
package scrape

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aliftech/jin/internal/core"
)

type HTTPDatabaseDetector struct {
	client *http.Client
}

func NewHTTPDatabaseDetector() *HTTPDatabaseDetector {
	return &HTTPDatabaseDetector{
		client: &http.Client{
			Timeout: 8 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

func (d *HTTPDatabaseDetector) Detect(ctx context.Context, url string) (*core.DatabaseInfo, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return &core.DatabaseInfo{URL: url}, err
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return &core.DatabaseInfo{URL: url}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)
	bodyLower := strings.ToLower(bodyStr)

	server := resp.Header.Get("Server")
	powered := resp.Header.Get("X-Powered-By")
	headersLower := strings.ToLower(server + " " + powered)

	info := &core.DatabaseInfo{
		URL:        url,
		Detected:   false,
		Database:   "unknown",
		Evidence:   []string{},
		Confidence: "low",
	}

	// === MongoDB Detection ===
	if strings.Contains(bodyLower, "objectid(") ||
		strings.Contains(bodyLower, "bson") ||
		strings.Contains(bodyLower, "mongodb") ||
		strings.Contains(bodyLower, "nosql") ||
		strings.Contains(bodyLower, "com.mongodb") {
		info.Detected = true
		info.Database = "MongoDB"
		info.Confidence = "high"
		info.Evidence = append(info.Evidence, "MongoDB-related string found in response")
		return info, nil
	}

	// === MySQL Detection ===
	mysqlPatterns := []string{
		"mysql",
		"you have an error in your sql syntax",
		"mysql_fetch",
		"mysqlexception",
		"sql syntax.*mysql",
		"unknown column",
		"mysql server version",
	}
	for _, p := range mysqlPatterns {
		if strings.Contains(bodyLower, p) {
			info.Detected = true
			info.Database = "MySQL"
			info.Confidence = "high"
			info.Evidence = append(info.Evidence, fmt.Sprintf("MySQL pattern matched: %q", p))
			return info, nil
		}
	}

	// === PostgreSQL Detection ===
	pgPatterns := []string{
		"postgresql",
		"pg_",
		"unterminated quoted string",
		"psql:",
		"pgerror",
		"syntax error at or near",
	}
	for _, p := range pgPatterns {
		if strings.Contains(bodyLower, p) {
			info.Detected = true
			info.Database = "PostgreSQL"
			info.Confidence = "high"
			info.Evidence = append(info.Evidence, fmt.Sprintf("PostgreSQL pattern matched: %q", p))
			return info, nil
		}
	}

	// === Microsoft SQL Server ===
	mssqlPatterns := []string{
		"microsoft sql server",
		"mssql",
		"sql server",
		"unclosed quotation mark",
		"sqlstate",
	}
	for _, p := range mssqlPatterns {
		if strings.Contains(bodyLower, p) {
			info.Detected = true
			info.Database = "Microsoft SQL Server"
			info.Confidence = "high"
			info.Evidence = append(info.Evidence, fmt.Sprintf("MSSQL pattern matched: %q", p))
			return info, nil
		}
	}

	// === SQLite ===
	if strings.Contains(bodyLower, "sqlite") || strings.Contains(bodyLower, "sqlite_error") {
		info.Detected = true
		info.Database = "SQLite"
		info.Confidence = "high"
		info.Evidence = append(info.Evidence, "SQLite mentioned in response")
		return info, nil
	}

	// === Oracle ===
	if strings.Contains(bodyLower, "ora-") || strings.Contains(bodyLower, "oracle error") {
		info.Detected = true
		info.Database = "Oracle"
		info.Confidence = "high"
		info.Evidence = append(info.Evidence, "Oracle error code (ORA-) detected")
		return info, nil
	}

	// === Fallback: Check Headers (Medium Confidence) ===
	if strings.Contains(headersLower, "mysql") {
		info.Detected = true
		info.Database = "MySQL"
		info.Confidence = "medium"
		info.Evidence = append(info.Evidence, "MySQL mentioned in Server/X-Powered-By header")
	} else if strings.Contains(headersLower, "postgres") {
		info.Detected = true
		info.Database = "PostgreSQL"
		info.Confidence = "medium"
		info.Evidence = append(info.Evidence, "PostgreSQL mentioned in header")
	}

	return info, nil
}
