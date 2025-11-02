// internal/bootstrap/app.go
package bootstrap

import (
	"time"

	"github.com/aliftech/jin/internal/module/db"
	info "github.com/aliftech/jin/internal/module/info"
	"github.com/aliftech/jin/internal/module/port"
	"github.com/aliftech/jin/internal/scrape"
)

func NewCLIHandler() *info.CLIHandler {
	scanner := scrape.NewHTTPServerScanner(10 * time.Second)
	return info.NewCLIHandler(scanner)
}

func NewPortScanHandler() *port.CLIHandler {
	scanner := scrape.NewTCPPortScanner(3 * time.Second) // faster per-port timeout
	return port.NewCLIHandler(scanner)
}

func NewDatabaseHandler() *db.CLIHandler {
	detector := scrape.NewHTTPDatabaseDetector()
	return db.NewCLIHandler(detector)
}
