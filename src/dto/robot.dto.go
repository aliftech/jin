package dto

// RobotsData holds robots.txt and sitemap.xml data
type RobotsData struct {
	DisallowedPaths []string `json:"disallowed_paths"`
	SitemapURLs     []string `json:"sitemap_urls"`
}
