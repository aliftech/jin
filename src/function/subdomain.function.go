package function

import (
	"fmt"
	"net"

	"github.com/aliftech/jin/src/dto"
)

// fetchSubdomains attempts to enumerate subdomains using a simple wordlist
func FetchSubdomains(host string) ([]dto.Subdomain, error) {
	var subdomains []dto.Subdomain
	commonSubdomains := []string{
		"www", "api", "blog", "mail", "dev", "test", "shop", "admin", "login", "staging",
	}

	for _, sub := range commonSubdomains {
		fqdn := fmt.Sprintf("%s.%s", sub, host)
		ips, err := net.LookupIP(fqdn)
		if err == nil {
			for _, ip := range ips {
				if ip.To4() != nil || ip.To16() != nil {
					subdomains = append(subdomains, dto.Subdomain{
						Name: fqdn,
						IP:   ip.String(),
					})
				}
			}
		}
	}

	return subdomains, nil
}
