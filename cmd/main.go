// cmd/main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aliftech/jin/internal/bootstrap"
	"github.com/aliftech/jin/internal/pkg/display"
	"github.com/fatih/color"
)

var (
	red = color.New(color.FgRed).SprintFunc()
	// green  = color.New(color.FgGreen).SprintFunc()
	// yellow = color.New(color.FgYellow).SprintFunc()
	blue = color.New(color.FgBlue).SprintFunc()
	cyan = color.New(color.FgCyan).SprintFunc()
)

func main() {
	fmt.Print(cyan(display.BANNER))
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("%s No target provided\n", red("✗"))
		display.PrintHelp()
		os.Exit(1)
	}

	// Handle direct URL: jin https://example.com
	if len(args) == 1 && !strings.HasPrefix(args[0], "-") {
		url := args[0]
		fmt.Printf("%s Fetching full server info for %s...\n", blue("⏳"), url)
		handler := bootstrap.NewCLIHandler()
		if err := handler.HandleScan(url, false); err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		return
	}

	// Handle commands: jin scrape -t <url>
	if len(args) >= 2 {
		command := strings.ToLower(args[0])
		if args[1] == "-t" && len(args) >= 3 {
			target := args[2]

			switch command {
			case "ports":
				host := cleanHost(target)
				fmt.Printf("%s Scanning open ports on %s...\n", blue("⏳"), host)
				handler := bootstrap.NewPortScanHandler()
				if err := handler.HandleScan(host, false, nil); err != nil {
					fmt.Printf("%s Error: %v\n", red("✗"), err)
					os.Exit(1)
				}

			case "db":
				fmt.Printf("%s Detecting backend database for %s...\n", blue("⏳"), target)
				handler := bootstrap.NewDatabaseHandler()
				if err := handler.HandleScan(target, false); err != nil {
					fmt.Printf("%s Error: %v\n", red("✗"), err)
					os.Exit(1)
				}

			default:
				fmt.Printf("%s Unknown command: %s\n", red("✗"), command)
				display.PrintHelp()
				os.Exit(1)
			}
			return
		}
	}

	fmt.Printf("%s Invalid usage\n", red("✗"))
	display.PrintHelp()
	os.Exit(1)
}

func cleanHost(input string) string {
	host := strings.TrimPrefix(input, "https://")
	host = strings.TrimPrefix(host, "http://")
	if i := strings.Index(host, "/"); i != -1 {
		host = host[:i]
	}
	return host
}
