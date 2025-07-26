package function

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aliftech/jin/src/dto"
)

// fetchTechStack performs technology stack fingerprinting
func FetchTechStack(resp *http.Response) ([]dto.TechStack, error) {
	var techs []dto.TechStack
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return techs, fmt.Errorf("failed to read response body: %v", err)
	}

	bodyStr := strings.ToLower(string(body))
	patterns := []struct {
		name    string
		pattern string
		source  string
	}{
		{name: "WordPress", pattern: "wp-content", source: "HTML (wp-content path)"},
		{name: "Joomla", pattern: "<meta name=\"generator\" content=\"joomla", source: "HTML (meta tag)"},
		{name: "Drupal", pattern: "drupal", source: "HTML (Drupal reference)"},
		{name: "jQuery", pattern: "jquery.min.js", source: "HTML (script tag)"},
		{name: "Bootstrap", pattern: "bootstrap.min.css", source: "HTML (CSS link)"},
		{name: "React", pattern: "react.min.js", source: "HTML (script tag)"},
		{name: "Vue.js", pattern: "vue.min.js", source: "HTML (script tag)"},
	}

	for _, p := range patterns {
		if strings.Contains(bodyStr, p.pattern) {
			techs = append(techs, dto.TechStack{
				Name:   p.name,
				Source: p.source,
			})
		}
	}

	// Check headers and cookies
	for _, cookie := range resp.Cookies() {
		cookieName := strings.ToLower(cookie.Name)
		if strings.Contains(cookieName, "wordpress") {
			techs = append(techs, dto.TechStack{Name: "WordPress", Source: "Cookie"})
		} else if strings.Contains(cookieName, "django") {
			techs = append(techs, dto.TechStack{Name: "Django", Source: "Cookie"})
		} else if strings.Contains(cookieName, "laravel_session") {
			techs = append(techs, dto.TechStack{Name: "Laravel", Source: "Cookie"})
		}
	}

	if poweredBy := strings.ToLower(resp.Header.Get("X-Powered-By")); poweredBy != "" {
		if strings.Contains(poweredBy, "php") {
			techs = append(techs, dto.TechStack{Name: "PHP", Source: "X-Powered-By header"})
		} else if strings.Contains(poweredBy, "asp.net") {
			techs = append(techs, dto.TechStack{Name: "ASP.NET", Source: "X-Powered-By header"})
		}
	}

	return techs, nil
}
