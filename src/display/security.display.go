package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displaySecurityHeaders prints security headers with colorful formatting
func DisplaySecurityHeaders(info *dto.WebsiteData, url string) {
	fmt.Printf("%s Security Headers for %s:\n", cyan("===="), url)
	if len(info.Security) > 0 {
		fmt.Printf("%s Security Headers Found:\n", green("✔"))
		for _, header := range info.Security {
			fmt.Printf("  %s: %s (%s)\n", green(header.Name), header.Value, header.Description)
		}
	} else {
		fmt.Printf("%s Security Headers: %s\n", yellow("⚠"), "None detected")
	}
}
