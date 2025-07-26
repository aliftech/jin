package function

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"github.com/aliftech/jin/src/dto"
)

// scrapeWebsite extracts title, meta description, and links from the webpage
func ScrapeWebsite(resp *http.Response) (dto.ScrapedData, error) {
	data := dto.ScrapedData{
		Links: []string{},
	}

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				return data, nil
			}
			return data, fmt.Errorf("error parsing HTML: %v", tokenizer.Err())
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			switch token.Data {
			case "title":
				if tt := tokenizer.Next(); tt == html.TextToken {
					data.Title = strings.TrimSpace(tokenizer.Token().Data)
				}
			case "meta":
				var name, content string
				for _, attr := range token.Attr {
					if attr.Key == "name" {
						name = strings.ToLower(attr.Val)
					}
					if attr.Key == "content" {
						content = strings.TrimSpace(attr.Val)
					}
				}
				if name == "description" {
					data.Description = content
				}
			case "a":
				for _, attr := range token.Attr {
					if attr.Key == "href" && attr.Val != "" {
						data.Links = append(data.Links, attr.Val)
					}
				}
			}
		}
	}
}
