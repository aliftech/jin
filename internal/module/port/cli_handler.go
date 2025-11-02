package port

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aliftech/jin/internal/core"
)

type CLIHandler struct {
	scanner core.PortScanner
}

func NewCLIHandler(scanner core.PortScanner) *CLIHandler {
	return &CLIHandler{scanner: scanner}
}

// DefaultPorts are commonly scanned ports
var DefaultPorts = []int{21, 22, 23, 25, 53, 80, 110, 143, 443, 993, 995, 3306, 5432, 6379, 27017}

func (h *CLIHandler) HandleScan(host string, outputJSON bool, customPorts []int) error {
	ports := DefaultPorts
	if len(customPorts) > 0 {
		ports = customPorts
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := h.scanner.Scan(ctx, host, ports)
	if err != nil {
		return fmt.Errorf("port scan failed: %w", err)
	}

	if outputJSON {
		return h.printJSON(results)
	}
	return h.printHumanReadable(results)
}

func (h *CLIHandler) printJSON(results []core.PortInfo) error {
	type Response struct {
		OpenPorts []core.PortInfo `json:"open_ports"`
	}
	openOnly := filterOpen(results)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(Response{OpenPorts: openOnly})
}

func (h *CLIHandler) printHumanReadable(results []core.PortInfo) error {
	fmt.Println("🔍 Open ports:")
	openFound := false
	for _, p := range results {
		if p.Status == "open" {
			fmt.Printf("  %d/%s\n", p.Port, p.Service)
			openFound = true
		}
	}
	if !openFound {
		fmt.Println("  No open ports found.")
	}
	return nil
}

func filterOpen(ports []core.PortInfo) []core.PortInfo {
	var open []core.PortInfo
	for _, p := range ports {
		if p.Status == "open" {
			open = append(open, p)
		}
	}
	return open
}
