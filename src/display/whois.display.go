package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displayWhois prints WHOIS data with colorful formatting
func DisplayWhois(info *dto.WebsiteData, host string) {
	fmt.Printf("%s WHOIS Data for %s:\n", cyan("===="), host)
	if info.Whois.DomainName != "" {
		fmt.Printf("%s Domain Name: %s\n", green("✔"), info.Whois.DomainName)
	} else {
		fmt.Printf("%s Domain Name: %s\n", yellow("⚠"), "Not detected")
	}
	if info.Whois.Registrar != "" {
		fmt.Printf("%s Registrar: %s\n", green("✔"), info.Whois.Registrar)
	} else {
		fmt.Printf("%s Registrar: %s\n", yellow("⚠"), "Not detected")
	}
	if info.Whois.CreationDate != "" {
		fmt.Printf("%s Creation Date: %s\n", green("✔"), info.Whois.CreationDate)
	} else {
		fmt.Printf("%s Creation Date: %s\n", yellow("⚠"), "Not detected")
	}
	if info.Whois.ExpirationDate != "" {
		fmt.Printf("%s Expiration Date: %s\n", green("✔"), info.Whois.ExpirationDate)
	} else {
		fmt.Printf("%s Expiration Date: %s\n", yellow("⚠"), "Not detected")
	}
	if info.Whois.Registrant != "" {
		fmt.Printf("%s Registrant: %s\n", green("✔"), info.Whois.Registrant)
	} else {
		fmt.Printf("%s Registrant: %s\n", yellow("⚠"), "Not detected (may be redacted)")
	}
}
