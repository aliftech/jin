package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/aliftech/jin/src/display"
	"github.com/aliftech/jin/src/function"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func main() {
	fmt.Println(cyan(display.BANNER))
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("%s No command provided\n", red("✗"))
		display.PrintHelp()
		os.Exit(1)
	}

	command := strings.ToLower(args[0])
	if command == "help" {
		display.PrintHelp()
		return
	}

	if len(args) < 3 && (command == "scrape" || command == "dns" || command == "subdomains" || command == "whois" || command == "ssl" || command == "robots" || command == "archive" || command == "security" || command == "tech") {
		fmt.Printf("%s Invalid command format. Use %s %s -t <url>\n", red("✗"), blue("serverinfo"), command)
		display.PrintHelp()
		os.Exit(1)
	}

	if command == "scrape" && args[1] == "-t" {
		url := args[2]
		fmt.Printf("%s Fetching and scraping data for %s...\n", blue("⏳"), url)
		info, err := function.FetchServerInfo(url, true, false, false, false, false, false, false, false, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		if err := display.PrintJSON(info); err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		return
	}

	if command == "dns" && args[1] == "-t" {
		url := args[2]
		host := strings.TrimPrefix(url, "https://")
		host = strings.TrimPrefix(host, "http://")
		host = strings.Split(host, "/")[0]
		fmt.Printf("%s Fetching DNS records for %s...\n", blue("⏳"), host)
		info, err := function.FetchServerInfo(url, false, true, false, false, false, false, false, false, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplayDNSRecords(info, host)
		return
	}

	if command == "subdomains" && args[1] == "-t" {
		url := args[2]
		host := strings.TrimPrefix(url, "https://")
		host = strings.TrimPrefix(host, "http://")
		host = strings.Split(host, "/")[0]
		fmt.Printf("%s Enumerating subdomains for %s...\n", blue("⏳"), host)
		info, err := function.FetchServerInfo(url, false, false, true, false, false, false, false, false, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplaySubdomains(info, host)
		return
	}

	if command == "whois" && args[1] == "-t" {
		url := args[2]
		host := strings.TrimPrefix(url, "https://")
		host = strings.TrimPrefix(host, "http://")
		host = strings.Split(host, "/")[0]
		fmt.Printf("%s Fetching WHOIS data for %s...\n", blue("⏳"), host)
		info, err := function.FetchServerInfo(url, false, false, false, true, false, false, false, false, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplayWhois(info, host)
		return
	}

	if command == "ssl" && args[1] == "-t" {
		url := args[2]
		host := strings.TrimPrefix(url, "https://")
		host = strings.TrimPrefix(host, "http://")
		host = strings.Split(host, "/")[0]
		fmt.Printf("%s Fetching SSL certificate for %s...\n", blue("⏳"), host)
		info, err := function.FetchServerInfo(url, false, false, false, false, true, false, false, false, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplaySSL(info, host)
		return
	}

	if command == "robots" && args[1] == "-t" {
		url := args[2]
		fmt.Printf("%s Fetching robots.txt and sitemap.xml for %s...\n", blue("⏳"), url)
		info, err := function.FetchServerInfo(url, false, false, false, false, false, true, false, false, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplayRobots(info, url)
		return
	}

	if command == "archive" && args[1] == "-t" {
		url := args[2]
		fmt.Printf("%s Fetching historical data for %s...\n", blue("⏳"), url)
		info, err := function.FetchServerInfo(url, false, false, false, false, false, false, true, false, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplayArchive(info, url)
		return
	}

	if command == "security" && args[1] == "-t" {
		url := args[2]
		fmt.Printf("%s Analyzing security headers for %s...\n", blue("⏳"), url)
		info, err := function.FetchServerInfo(url, false, false, false, false, false, false, false, true, false)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplaySecurityHeaders(info, url)
		return
	}

	if command == "tech" && args[1] == "-t" {
		url := args[2]
		fmt.Printf("%s Detecting technology stack for %s...\n", blue("⏳"), url)
		info, err := function.FetchServerInfo(url, false, false, false, false, false, false, false, false, true)
		if err != nil {
			fmt.Printf("%s Error: %v\n", red("✗"), err)
			os.Exit(1)
		}
		display.DisplayTechStack(info, url)
		return
	}

	url := args[0]
	fmt.Printf("%s Fetching server info for %s...\n", blue("⏳"), url)
	info, err := function.FetchServerInfo(url, false, false, false, false, false, false, false, false, false)
	if err != nil {
		fmt.Printf("%s Error: %v\n", red("✗"), err)
		os.Exit(1)
	}

	display.DisplayServerInfo(info, url)
}
