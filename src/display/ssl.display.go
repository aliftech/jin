package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displaySSL prints SSL/TLS certificate data with colorful formatting
func DisplaySSL(info *dto.WebsiteData, host string) {
	fmt.Printf("%s SSL/TLS Certificate for %s:\n", cyan("===="), host)
	if info.SSL.Issuer != "" {
		fmt.Printf("%s Issuer: %s\n", green("✔"), info.SSL.Issuer)
	} else {
		fmt.Printf("%s Issuer: %s\n", yellow("⚠"), "Not detected")
	}
	if info.SSL.Subject != "" {
		fmt.Printf("%s Subject: %s\n", green("✔"), info.SSL.Subject)
	} else {
		fmt.Printf("%s Subject: %s\n", yellow("⚠"), "Not detected")
	}
	if info.SSL.Expiration != "" {
		fmt.Printf("%s Expiration: %s\n", green("✔"), info.SSL.Expiration)
	} else {
		fmt.Printf("%s Expiration: %s\n", yellow("⚠"), "Not detected")
	}
	if len(info.SSL.SANs) > 0 {
		fmt.Printf("%s Subject Alternative Names:\n", green("✔"))
		for _, san := range info.SSL.SANs {
			fmt.Printf("  %s\n", green(san))
		}
	} else {
		fmt.Printf("%s Subject Alternative Names: %s\n", yellow("⚠"), "None detected")
	}
}
