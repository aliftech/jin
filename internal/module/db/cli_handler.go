// internal/module/db/cli_handler.go
package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aliftech/jin/internal/core"
)

type CLIHandler struct {
	detector core.DatabaseDetector
}

func NewCLIHandler(detector core.DatabaseDetector) *CLIHandler {
	return &CLIHandler{detector: detector}
}

func (h *CLIHandler) HandleScan(url string, outputJSON bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := h.detector.Detect(ctx, url)
	if err != nil {
		return fmt.Errorf("database detection failed: %w", err)
	}

	if outputJSON {
		return h.printJSON(info)
	}
	return h.printHumanReadable(info)
}

func (h *CLIHandler) printJSON(info *core.DatabaseInfo) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(info)
}

func (h *CLIHandler) printHumanReadable(info *core.DatabaseInfo) error {
	if !info.Detected {
		fmt.Println("❌ No database detected.")
		return nil
	}

	// Simple color functions (no external dep beyond what you have)
	green := func(s string) string { return "\033[32m" + s + "\033[0m" }
	yellow := func(s string) string { return "\033[33m" + s + "\033[0m" }

	fmt.Printf("✅ Database detected: %s\n", green(info.Database))
	fmt.Printf("📊 Confidence: %s\n", yellow(info.Confidence))
	fmt.Println("🔍 Evidence:")
	for _, e := range info.Evidence {
		fmt.Printf("  • %s\n", e)
	}
	return nil
}
