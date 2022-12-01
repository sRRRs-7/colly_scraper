package scrape

import "os"

type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
}

func saveJson(fName string, p *info) {
	f, err := os.Create(fName)
	if err != nil {}
}