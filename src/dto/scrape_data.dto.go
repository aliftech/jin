package dto

// ScrapedData holds the scraped website data
type ScrapedData struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
}
