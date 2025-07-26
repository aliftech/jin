package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displayServerInfo prints server information with colorful formatting
func DisplayServerInfo(info *dto.WebsiteData, url string) {
	fmt.Printf("%s Server Information for %s:\n", cyan("===="), url)
	if info.ServerInfo.WebServer != "" {
		fmt.Printf("%s Web Server: %s\n", green("✔"), info.ServerInfo.WebServer)
	} else {
		fmt.Printf("%s Web Server: %s\n", yellow("⚠"), "Not detected (Server header missing or generic)")
	}

	if info.ServerInfo.OS != "" {
		fmt.Printf("%s Operating System: %s\n", green("✔"), info.ServerInfo.OS)
	} else {
		fmt.Printf("%s Operating System: %s\n", yellow("⚠"), "Not detected (insufficient header information)")
	}

	if info.ServerInfo.Database != "" {
		fmt.Printf("%s Database: %s\n", green("✔"), info.ServerInfo.Database)
	} else {
		fmt.Printf("%s Database: %s\n", yellow("⚠"), "Not detected (no database-specific headers or endpoints)")
	}

	if info.ServerInfo.PoweredBy != "" {
		fmt.Printf("%s Powered By: %s\n", green("✔"), info.ServerInfo.PoweredBy)
	} else {
		fmt.Printf("%s Powered By: %s\n", yellow("⚠"), "Not detected (X-Powered-By header missing)")
	}

	if info.ServerInfo.Framework != "" {
		fmt.Printf("%s Framework: %s\n", green("✔"), info.ServerInfo.Framework)
	} else {
		fmt.Printf("%s Framework: %s\n", yellow("⚠"), "Not detected (no framework-specific headers or endpoints)")
	}

	if len(info.ServerInfo.OpenPorts) > 0 {
		fmt.Printf("%s Open Ports:\n", green("✔"))
		for _, port := range info.ServerInfo.OpenPorts {
			fmt.Printf("  %s\n", green(port))
		}
	} else {
		fmt.Printf("%s Open Ports: %s\n", yellow("⚠"), "None detected (ports may be firewalled)")
	}

	if len(info.ServerInfo.OtherHeaders) > 0 {
		fmt.Printf("%s Other Headers:\n", blue("📋"))
		for key, value := range info.ServerInfo.OtherHeaders {
			fmt.Printf("  %s: %s\n", blue(key), value)
		}
	}
}
