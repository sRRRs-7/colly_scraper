package scrape

import (
	"encoding/json"
	"log"
	"os"
)

type info struct {
	StatusCode int `json:"status_code"`
	URL string `json:"url"`
	Title string `json:"title"`
}

func saveJson(fName string, p *info) {
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal("cannot create file: ", err)
		return
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(p)
	if err != nil {
		log.Fatal("cannot encode json file: ", err)
	}
}