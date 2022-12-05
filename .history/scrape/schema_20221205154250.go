package scrape

type Info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
	Article []ProductInfo `json:"product"`
}

type ProductInfo struct {
	ID int `json:"id"`
	Name string `json:"title"`
	URL string `json:"url"`
	Price string `json:"price"`
	Star string `json:"star"`
	Image string `json:"image"`
}



