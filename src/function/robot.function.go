package function

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"github.com/aliftech/jin/src/dto"
)

// fetchRobots fetches and parses robots.txt and sitemap.xml
func FetchRobots(url string, client *http.Client) (dto.RobotsData, error) {
	data := dto.RobotsData{}

	// Fetch robots.txt
	resp, err := client.Get(url + "/robots.txt")
	if err == nil && resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err == nil {
			lines := strings.Split(string(body), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(strings.ToLower(line), "disallow:") {
					path := strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
					if path != "" {
						data.DisallowedPaths = append(data.DisallowedPaths, path)
					}
				}
				if strings.HasPrefix(strings.ToLower(line), "sitemap:") {
					sitemapURL := strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
					if sitemapURL != "" {
						data.SitemapURLs = append(data.SitemapURLs, sitemapURL)
					}
				}
			}
		}
	}

	// Fetch sitemap.xml
	resp, err = client.Get(url + "/sitemap.xml")
	if err == nil && resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err == nil {
			tokenizer := html.NewTokenizer(strings.NewReader(string(body)))
			for {
				tt := tokenizer.Next()
				if tt == html.ErrorToken {
					break
				}
				if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
					token := tokenizer.Token()
					if token.Data == "loc" {
						if tt := tokenizer.Next(); tt == html.TextToken {
							url := strings.TrimSpace(tokenizer.Token().Data)
							if url != "" {
								data.SitemapURLs = append(data.SitemapURLs, url)
							}
						}
					}
				}
			}
		}
	}

	return data, nil
}
