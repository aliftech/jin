package function

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aliftech/jin/src/dto"
	"github.com/fatih/color"
)

// Colors for output formattin
var (
	blue = color.New(color.FgBlue).SprintFunc()
)

// fetchServerInfo makes HTTP requests and extracts server information
func FetchServerInfo(url string, scrapeMode, dnsMode, subdomainsMode, whoisMode, sslMode, robotsMode, archiveMode, securityMode, techMode bool) (*dto.WebsiteData, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Don't follow redirects
		},
	}

	// Ensure URL has protocol
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	// Extract host from URL
	host := strings.TrimPrefix(url, "https://")
	host = strings.TrimPrefix(host, "http://")
	host = strings.Split(host, "/")[0]

	// Normalize URL (remove trailing slash)
	url = strings.TrimSuffix(url, "/")

	info := &dto.WebsiteData{
		ServerInfo: dto.ServerInfo{
			OtherHeaders: make(map[string]string),
		},
	}

	// Fetch DNS records
	if dnsMode {
		fmt.Printf("%s Fetching DNS records for %s...\n", blue("⏳"), host)
		records, err := FetchDNSRecords(host)
		if err != nil {
			return info, fmt.Errorf("failed to fetch DNS records: %v", err)
		}
		info.DNSRecords = records
		return info, nil
	}

	// Fetch subdomains
	if subdomainsMode {
		fmt.Printf("%s Enumerating subdomains for %s...\n", blue("⏳"), host)
		subdomains, err := FetchSubdomains(host)
		if err != nil {
			return info, fmt.Errorf("failed to enumerate subdomains: %v", err)
		}
		info.Subdomains = subdomains
		return info, nil
	}

	// Fetch WHOIS data
	if whoisMode {
		fmt.Printf("%s Fetching WHOIS data for %s...\n", blue("⏳"), host)
		whoisData, err := FetchWhois(host)
		if err != nil {
			return info, fmt.Errorf("failed to fetch WHOIS data: %v", err)
		}
		info.Whois = whoisData
		return info, nil
	}

	// Fetch SSL/TLS certificate
	if sslMode {
		if !strings.HasPrefix(url, "https://") {
			return info, fmt.Errorf("SSL analysis requires an HTTPS URL")
		}
		fmt.Printf("%s Fetching SSL certificate for %s...\n", blue("⏳"), host)
		sslData, err := FetchSSL(host)
		if err != nil {
			return info, fmt.Errorf("failed to fetch SSL certificate: %v", err)
		}
		info.SSL = sslData
		return info, nil
	}

	// Fetch robots.txt and sitemap.xml
	if robotsMode {
		fmt.Printf("%s Fetching robots.txt and sitemap.xml for %s...\n", blue("⏳"), url)
		robotsData, err := FetchRobots(url, client)
		if err != nil {
			return info, fmt.Errorf("failed to fetch robots data: %v", err)
		}
		info.Robots = robotsData
		return info, nil
	}

	// Fetch archive data
	if archiveMode {
		fmt.Printf("%s Fetching historical data for %s...\n", blue("⏳"), url)
		archiveData, err := FetchArchive(url, client)
		if err != nil {
			return info, fmt.Errorf("failed to fetch archive data: %v", err)
		}
		info.Archive = archiveData
		return info, nil
	}

	// Main request to the root URL
	resp, err := client.Get(url)
	if err != nil {
		return info, fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	// Scrape website data
	if scrapeMode {
		scrapedData, err := ScrapeWebsite(resp)
		if err != nil {
			return info, fmt.Errorf("failed to scrape website: %v", err)
		}
		info.ScrapedData = scrapedData
		// Reopen response body for server info
		resp.Body.Close()
		resp, err = client.Get(url)
		if err != nil {
			return info, fmt.Errorf("failed to fetch URL for server info: %v", err)
		}
		defer resp.Body.Close()
	}

	// Fetch security headers
	if securityMode {
		fmt.Printf("%s Analyzing security headers for %s...\n", blue("⏳"), url)
		info.Security = FetchSecurityHeaders(resp)
		return info, nil
	}

	// Fetch technology stack
	if techMode {
		fmt.Printf("%s Detecting technology stack for %s...\n", blue("⏳"), url)
		techs, err := FetchTechStack(resp)
		if err != nil {
			return info, fmt.Errorf("failed to detect technology stack: %v", err)
		}
		info.TechStack = techs
		return info, nil
	}

	// Scan ports
	fmt.Printf("%s Scanning ports for %s...\n", blue("⏳"), host)
	info.ServerInfo.OpenPorts = ScanPorts(host)

	// Extract Server header
	if server := resp.Header.Get("Server"); server != "" {
		info.ServerInfo.WebServer = server
		serverLower := strings.ToLower(server)
		if strings.Contains(serverLower, "win") || strings.Contains(serverLower, "iis") {
			info.ServerInfo.OS = "Windows"
		} else if strings.Contains(serverLower, "linux") || strings.Contains(serverLower, "unix") || strings.Contains(serverLower, "apache") || strings.Contains(serverLower, "nginx") {
			info.ServerInfo.OS = "Linux/Unix (likely)"
		}
	}

	// Extract X-Powered-By header
	if poweredBy := resp.Header.Get("X-Powered-By"); poweredBy != "" {
		info.ServerInfo.PoweredBy = poweredBy
		poweredByLower := strings.ToLower(poweredBy)
		if strings.Contains(poweredByLower, "php") {
			info.ServerInfo.Database = "MySQL (likely, based on PHP)"
		} else if strings.Contains(poweredByLower, "asp.net") {
			info.ServerInfo.Database = "Microsoft SQL Server (likely, based on ASP.NET)"
		}
	}

	// Check for framework-specific headers or cookies
	if wp := resp.Header.Get("X-Generator"); strings.Contains(strings.ToLower(wp), "wordpress") {
		info.ServerInfo.Framework = "WordPress"
		info.ServerInfo.Database = "MySQL (likely, based on WordPress)"
	}
	for _, cookie := range resp.Cookies() {
		cookieName := strings.ToLower(cookie.Name)
		if strings.Contains(cookieName, "wordpress") {
			info.ServerInfo.Framework = "WordPress"
			info.ServerInfo.Database = "MySQL (likely, based on WordPress)"
		} else if strings.Contains(cookieName, "django") {
			info.ServerInfo.Framework = "Django"
			info.ServerInfo.Database = "PostgreSQL or SQLite (likely, based on Django)"
		} else if strings.Contains(cookieName, "laravel_session") {
			info.ServerInfo.Framework = "Laravel"
			info.ServerInfo.Database = "MySQL (likely, based on Laravel)"
		}
	}

	// Collect other interesting headers
	interestingHeaders := []string{
		"X-Frame-Options",
		"Content-Security-Policy",
		"X-AspNet-Version",
		"X-Runtime",
		"X-Version",
		"X-Drupal-Cache",
		"X-Pingback",
		"X-Served-By",
		"Via",
	}
	for _, header := range interestingHeaders {
		if value := resp.Header.Get(header); value != "" {
			info.ServerInfo.OtherHeaders[header] = value
			if header == "X-Drupal-Cache" {
				info.ServerInfo.Framework = "Drupal"
				info.ServerInfo.Database = "MySQL or PostgreSQL (likely, based on Drupal)"
			}
		}
	}

	// Basic HTML content analysis (if not in scrape mode)
	if !scrapeMode {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			bodyStr := strings.ToLower(string(body))
			if strings.Contains(bodyStr, "<meta name=\"generator\" content=\"wordpress") {
				info.ServerInfo.Framework = "WordPress"
				info.ServerInfo.Database = "MySQL (likely, based on WordPress meta tag)"
			} else if strings.Contains(bodyStr, "<meta name=\"generator\" content=\"joomla") {
				info.ServerInfo.Framework = "Joomla"
				info.ServerInfo.Database = "MySQL (likely, based on Joomla)"
			} else if strings.Contains(bodyStr, "rails") {
				info.ServerInfo.Framework = "Ruby on Rails"
				info.ServerInfo.Database = "PostgreSQL or MySQL (likely, based on Rails)"
			}
		}
	}

	// Probe common endpoints to detect technologies
	techTests := []struct {
		path      string
		framework string
		database  string
	}{
		{path: "wp-login.php", framework: "PHP/WordPress", database: "MySQL"},
		{path: ".aspx", framework: "ASP.NET", database: "Microsoft SQL Server"},
		{path: "index.php", framework: "PHP", database: "MySQL"},
		{path: ".jsp", framework: "Java/JSP", database: "MySQL or PostgreSQL"},
	}
	for _, test := range techTests {
		if info.ServerInfo.Framework != "" {
			continue
		}
		probeURL := url + "/" + test.path
		resp, err := client.Get(probeURL)
		if err != nil {
			continue
		}
		if resp.StatusCode == 200 || resp.StatusCode == 301 || resp.StatusCode == 302 {
			info.ServerInfo.Framework = test.framework
			info.ServerInfo.Database = test.database + " (inferred from endpoint)"
		}
		resp.Body.Close()
	}

	return info, nil
}
