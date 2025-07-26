package function

import (
	"net"

	"github.com/aliftech/jin/src/dto"
)

// fetchDNSRecords queries DNS records for the given host
func FetchDNSRecords(host string) ([]dto.DNSRecord, error) {
	var records []dto.DNSRecord
	recordTypes := []string{"A", "AAAA", "MX", "TXT", "NS", "CNAME"}

	for _, rType := range recordTypes {
		var values []string
		switch rType {
		case "A":
			ips, err := net.LookupIP(host)
			if err == nil {
				for _, ip := range ips {
					if ip.To4() != nil {
						values = append(values, ip.String())
					}
				}
			}
		case "AAAA":
			ips, err := net.LookupIP(host)
			if err == nil {
				for _, ip := range ips {
					if ip.To16() != nil && ip.To4() == nil {
						values = append(values, ip.String())
					}
				}
			}
		case "MX":
			mxs, err := net.LookupMX(host)
			if err == nil {
				for _, mx := range mxs {
					values = append(values, mx.Host)
				}
			}
		case "TXT":
			txts, err := net.LookupTXT(host)
			if err == nil {
				values = append(values, txts...)
			}
		case "NS":
			nss, err := net.LookupNS(host)
			if err == nil {
				for _, ns := range nss {
					values = append(values, ns.Host)
				}
			}
		case "CNAME":
			cname, err := net.LookupCNAME(host)
			if err == nil && cname != "" && cname != host+"." {
				values = append(values, cname)
			}
		}
		if len(values) > 0 {
			records = append(records, dto.DNSRecord{
				Type:  rType,
				Value: values,
			})
		}
	}

	return records, nil
}
