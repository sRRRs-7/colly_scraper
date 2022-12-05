package scrape

type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
	Article []productInfo `json:"product"`
}

type productInfo struct {
	Name string `json:"title"`
	URL string `json:"url"`
	Price string `json:"price"`
	Stars string `json: "stars"`
}



