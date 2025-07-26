package display

import "fmt"

func PrintHelp() {
	fmt.Printf("%s Website Server Info CLI\n", cyan("===="))
	fmt.Println("Usage:")
	fmt.Printf("  %s <url>              - Fetch server information and scan ports for the given URL\n", blue("jin"))
	fmt.Printf("  %s scrape -t <url>    - Scrape website data and return server info in JSON format\n", blue("jin"))
	fmt.Printf("  %s dns -t <url>       - Fetch and display DNS records for the given URL\n", blue("jin"))
	fmt.Printf("  %s subdomains -t <url> - Enumerate subdomains for the given URL\n", blue("jin"))
	fmt.Printf("  %s whois -t <url>     - Fetch and display WHOIS data for the given URL\n", blue("jin"))
	fmt.Printf("  %s ssl -t <url>       - Fetch and display SSL/TLS certificate data for the given URL\n", blue("jin"))
	fmt.Printf("  %s robots -t <url>    - Fetch and display robots.txt and sitemap.xml data\n", blue("jin"))
	fmt.Printf("  %s archive -t <url>   - Fetch and display historical snapshots from Wayback Machine\n", blue("jin"))
	fmt.Printf("  %s security -t <url>  - Analyze and display security headers\n", blue("jin"))
	fmt.Printf("  %s tech -t <url>      - Detect and display technology stack\n", blue("jin"))
	fmt.Printf("  %s help               - Show this help message\n", blue("jin"))
}
