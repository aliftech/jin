package display

import (
	"fmt"

	"github.com/aliftech/jin/src/dto"
)

// displayRobots prints robots.txt and sitemap.xml data with colorful formatting
func DisplayRobots(info *dto.WebsiteData, url string) {
	fmt.Printf("%s Robots and Sitemap Data for %s:\n", cyan("===="), url)
	if len(info.Robots.DisallowedPaths) > 0 {
		fmt.Printf("%s Disallowed Paths (robots.txt):\n", green("✔"))
		for _, path := range info.Robots.DisallowedPaths {
			fmt.Printf("  %s\n", green(path))
		}
	} else {
		fmt.Printf("%s Disallowed Paths (robots.txt): %s\n", yellow("⚠"), "None detected or robots.txt missing")
	}
	if len(info.Robots.SitemapURLs) > 0 {
		fmt.Printf("%s Sitemap URLs:\n", green("✔"))
		for _, sitemap := range info.Robots.SitemapURLs {
			fmt.Printf("  %s\n", green(sitemap))
		}
	} else {
		fmt.Printf("%s Sitemap URLs: %s\n", yellow("⚠"), "None detected or sitemap.xml missing")
	}
}
