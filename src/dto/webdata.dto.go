package dto

// WebsiteData combines all collected data
type WebsiteData struct {
	ServerInfo  ServerInfo       `json:"server_info"`
	ScrapedData ScrapedData      `json:"scraped_data"`
	DNSRecords  []DNSRecord      `json:"dns_records"`
	Subdomains  []Subdomain      `json:"subdomains"`
	Whois       WhoisData        `json:"whois"`
	SSL         SSLData          `json:"ssl"`
	Robots      RobotsData       `json:"robots"`
	Archive     ArchiveData      `json:"archive"`
	Security    []SecurityHeader `json:"security_headers"`
	TechStack   []TechStack      `json:"tech_stack"`
}
