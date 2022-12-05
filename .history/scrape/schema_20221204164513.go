package scrape

type article struct {
	Title string `json:"title"`
	URL string `json:"url"`
}
type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
	Article article 'json:"article"'
}

