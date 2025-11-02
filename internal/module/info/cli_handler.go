package info

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aliftech/jin/internal/core"
	"github.com/aliftech/jin/internal/module/info/dto"
)

type CLIHandler struct {
	scanner core.ServerScanner
}

func NewCLIHandler(scanner core.ServerScanner) *CLIHandler {
	return &CLIHandler{scanner: scanner}
}

func (h *CLIHandler) HandleScan(url string, outputJSON bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := h.scanner.Scan(ctx, url)
	if err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}

	if outputJSON {
		return h.printJSON(info)
	}
	return h.printHumanReadable(info)
}

func (h *CLIHandler) printJSON(info *core.ServerInfo) error {
	response := dto.FromDomain(info)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(response)
}

func (h *CLIHandler) printHumanReadable(info *core.ServerInfo) error {
	fmt.Printf("🌐 URL:          %s\n", info.URL)
	fmt.Printf("✅ Status:       %d\n", info.StatusCode)
	fmt.Printf("🖥️  Server:      %s\n", info.Server)
	fmt.Printf("⚡ Powered By:   %s\n", info.PoweredBy)
	fmt.Printf("📄 Content-Type: %s\n", info.ContentType)

	if info.TLSVersion != "" {
		fmt.Printf("🔒 TLS Version:  %s\n", info.TLSVersion)
		fmt.Printf("🔐 Cipher Suite: %s\n", info.TLSCipherSuite)
	}

	fmt.Println("\n📋 Headers:")
	for key, value := range info.Headers {
		fmt.Printf("  %s: %s\n", key, value)
	}
	return nil
}
