package scrape

type productInfo struct {
	Name string `json:"title"`
	URL string `json:"url"`
	Price string `json:"name"`
	Description string `json:"destination"`

}

type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
	Article []articleInfo `json:"article"`
}

