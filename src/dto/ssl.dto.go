package dto

// SSLData holds SSL/TLS certificate data
type SSLData struct {
	Issuer     string   `json:"issuer"`
	Subject    string   `json:"subject"`
	Expiration string   `json:"expiration"`
	SANs       []string `json:"sans"`
}
