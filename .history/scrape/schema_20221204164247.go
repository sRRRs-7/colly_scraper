package scrape

type article struct {}
type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
}

