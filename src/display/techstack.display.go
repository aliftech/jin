package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displayTechStack prints technology stack with colorful formatting
func DisplayTechStack(info *dto.WebsiteData, url string) {
	fmt.Printf("%s Technology Stack for %s:\n", cyan("===="), url)
	if len(info.TechStack) > 0 {
		fmt.Printf("%s Technologies Detected:\n", green("✔"))
		for _, tech := range info.TechStack {
			fmt.Printf("  %s (%s)\n", green(tech.Name), tech.Source)
		}
	} else {
		fmt.Printf("%s Technologies: %s\n", yellow("⚠"), "None detected")
	}
}
