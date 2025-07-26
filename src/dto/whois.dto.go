package dto

// WhoisData holds WHOIS lookup data
type WhoisData struct {
	DomainName     string `json:"domain_name"`
	Registrar      string `json:"registrar"`
	CreationDate   string `json:"creation_date"`
	ExpirationDate string `json:"expiration_date"`
	Registrant     string `json:"registrant"`
}
