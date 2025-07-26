package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displaySubdomains prints subdomains with colorful formatting
func DisplaySubdomains(info *dto.WebsiteData, host string) {
	fmt.Printf("%s Subdomains for %s:\n", cyan("===="), host)
	if len(info.Subdomains) > 0 {
		fmt.Printf("%s Subdomains Found:\n", green("✔"))
		for _, sub := range info.Subdomains {
			fmt.Printf("  %s (%s)\n", green(sub.Name), sub.IP)
		}
	} else {
		fmt.Printf("%s Subdomains: %s\n", yellow("⚠"), "None detected")
	}
}
