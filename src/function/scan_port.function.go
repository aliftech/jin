package function

import (
	"fmt"
	"net"
	"time"
)

// scanPorts checks a predefined list of common ports
func ScanPorts(host string) []string {
	commonPorts := []struct {
		port string
		name string
	}{
		{"80", "HTTP"},
		{"443", "HTTPS"},
		{"22", "SSH"},
		{"21", "FTP"},
		{"3306", "MySQL"},
		{"5432", "PostgreSQL"},
		{"8080", "HTTP-Alternate"},
	}
	var openPorts []string

	for _, p := range commonPorts {
		addr := fmt.Sprintf("%s:%s", host, p.port)
		conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
		if err == nil {
			openPorts = append(openPorts, fmt.Sprintf("%s (%s)", p.port, p.name))
			conn.Close()
		}
	}

	return openPorts
}
