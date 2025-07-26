package function

import (
	"fmt"
	"strings"

	"github.com/domainr/whois"

	"github.com/aliftech/jin/src/dto"
)

// fetchWhois performs a WHOIS lookup
func FetchWhois(host string) (dto.WhoisData, error) {
	data := dto.WhoisData{}
	req, err := whois.NewRequest(host)
	if err != nil {
		return data, fmt.Errorf("failed to create WHOIS request: %v", err)
	}
	resp, err := whois.DefaultClient.Fetch(req)
	if err != nil {
		return data, fmt.Errorf("failed to fetch WHOIS data: %v", err)
	}

	body := strings.ToLower(string(resp.Body))
	data.DomainName = parseWhoisField(body, "domain name")
	data.Registrar = parseWhoisField(body, "registrar")
	data.CreationDate = parseWhoisField(body, "creation date")
	data.ExpirationDate = parseWhoisField(body, "expiration date")
	data.Registrant = parseWhoisField(body, "registrant name")
	if data.Registrant == "" {
		data.Registrant = parseWhoisField(body, "registrant organization")
	}

	return data, nil
}

// parseWhoisField extracts a field from WHOIS response
func parseWhoisField(body, field string) string {
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if strings.HasPrefix(strings.ToLower(line), strings.ToLower(field)) {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}
