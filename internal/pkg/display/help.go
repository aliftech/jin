package display

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	cyan  = color.New(color.FgCyan).SprintFunc()
	white = color.New(color.FgWhite).SprintFunc()
)

// PrintHelp displays usage instructions for JIN
func PrintHelp() {
	fmt.Println(white("Usage:"))
	fmt.Println("  " + green("jin <url>                          ") + white("→ Full server reconnaissance"))
	fmt.Println("  " + green("jin [command] -t <target>          ") + white("→ Run specific module"))
	fmt.Println()
	fmt.Println(white("Available Commands:"))
	fmt.Println("  " + cyan("ports") + white("      Scan for open ports (TCP connect scan)"))
	fmt.Println("  " + cyan("db") + white("         Detect backend database (MySQL, PostgreSQL, etc.)"))
	fmt.Println()
	fmt.Println(white("Flags:"))
	fmt.Println("  " + green("-j, --json") + white("   Output as JSON (for scrape command)"))
	fmt.Println("  " + green("-h, --help") + white("   Show this help message"))
	fmt.Println()
	fmt.Println(white("Examples:"))
	fmt.Println("  jin https://example.com")
	fmt.Println("  jin ports -t example.com")
	fmt.Println("  jin db -t example.com")
}
