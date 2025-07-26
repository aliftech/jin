package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displayDNSRecords prints DNS records with colorful formatting
func DisplayDNSRecords(info *dto.WebsiteData, host string) {
	fmt.Printf("%s DNS Records for %s:\n", cyan("===="), host)
	recordTypes := []string{"A", "AAAA", "MX", "TXT", "NS", "CNAME"}
	found := make(map[string][]string)

	for _, record := range info.DNSRecords {
		found[record.Type] = record.Value
	}

	for _, rType := range recordTypes {
		if values, exists := found[rType]; exists && len(values) > 0 {
			fmt.Printf("%s %s Records:\n", green("✔"), rType)
			for _, value := range values {
				fmt.Printf("  %s\n", green(value))
			}
		} else {
			fmt.Printf("%s %s Records: %s\n", yellow("⚠"), rType, "None detected")
		}
	}
}
