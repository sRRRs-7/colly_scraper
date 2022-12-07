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

	for page := 1; page <= 10; page++ {
		page_str := strconv.Itoa(page)
		url := "https://www.amazon.co.jp/s?k=anker&crid=3OXITLFR8UF3J&sprefix=anke%2Caps%2C247&ref=nb_sb_noss_" + page_str
		// methods
		info = scrape.GetAmazon(c, url, &product, &info)
	}

	scrape.SaveFileJson("scrape.json", &info)
}