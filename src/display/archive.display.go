package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displayArchive prints historical snapshot data with colorful formatting
func DisplayArchive(info *dto.WebsiteData, url string) {
	fmt.Printf("%s Historical Snapshots for %s:\n", cyan("===="), url)
	if len(info.Archive.Snapshots) > 0 {
		fmt.Printf("%s Snapshots Found:\n", green("✔"))
		for _, snapshot := range info.Archive.Snapshots {
			fmt.Printf("  %s: %s\n", green(snapshot.Timestamp), snapshot.URL)
		}
	} else {
		fmt.Printf("%s Snapshots: %s\n", yellow("⚠"), "None detected in Wayback Machine")
	}
}
