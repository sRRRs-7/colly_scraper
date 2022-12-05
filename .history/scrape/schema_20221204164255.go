package scrape

type article struct {
	Title string `json:"title"`
}
type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
}

