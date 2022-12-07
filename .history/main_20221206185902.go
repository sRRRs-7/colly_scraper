package main

import (
	"strconv"

	"github.com/sRRRs-7/colly_scraper/scrape"
)


func main() {
	// instance
	c := scrape.NewScraper()
	// struct instance
	product := []scrape.ProductInfo{}
	info := scrape.Info{}
	id := 0

	for page := 1; page <= 1; page++ {
		page_str := strconv.Itoa(page)
		url := "https://www.amazon.co.jp/s?k=%E4%BA%BA%E3%82%92%E3%83%80%E3%83%A1%E3%81%AB%E3%81%99%E3%82%8B&crid=3PZWAG8WRMCCP&sprefix=%E4%BA%BA%E3%82%92%E3%83%80%E3%83%A1%E3%81%AB%E3%81%99%E3%82%8B%2Caps%2C239&ref=nb_sb_noss_" + page_str
		// methods
		info = scrape.GetAmazon(c, url, &product, &info, &id)
	}

	scrape.SaveFileJson("scrape.json", &info)
}