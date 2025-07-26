package function

import (
	"net/http"

	"github.com/aliftech/jin/src/dto"
)

// fetchSecurityHeaders analyzes security headers
func FetchSecurityHeaders(resp *http.Response) []dto.SecurityHeader {
	headers := []dto.SecurityHeader{
		{Name: "Strict-Transport-Security", Description: "Enforces HTTPS connections"},
		{Name: "X-Content-Type-Options", Description: "Prevents MIME type sniffing"},
		{Name: "X-XSS-Protection", Description: "Enables XSS filtering in browsers"},
		{Name: "Content-Security-Policy", Description: "Controls resources the browser can load"},
		{Name: "X-Frame-Options", Description: "Prevents clickjacking"},
		{Name: "Referrer-Policy", Description: "Controls referrer information sent"},
	}

	var result []dto.SecurityHeader
	for _, header := range headers {
		if value := resp.Header.Get(header.Name); value != "" {
			result = append(result, dto.SecurityHeader{
				Name:        header.Name,
				Value:       value,
				Description: header.Description,
			})
		}
	}

	return result
}
