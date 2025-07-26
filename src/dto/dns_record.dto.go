package dto

// DNSRecord holds DNS record data
type DNSRecord struct {
	Type  string   `json:"type"`
	Value []string `json:"value"`
}
