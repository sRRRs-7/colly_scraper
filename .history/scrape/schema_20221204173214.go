package scrape

type articleInfo struct {
	Name string `json:"title"`
	URL string `json:"url"`
	Name string `json:"name"`
}

type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
	Article []articleInfo `json:"article"`
}

