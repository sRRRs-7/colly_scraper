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

	for page := 1; page <= 10; page++ {
		page_str := strconv.Itoa(page)
		url := "https://www.amazon.co.jp/s?k=%E5%8C%96%E7%B2%A7%E6%B0%B4&crid=128Y9RRPHWJQJ&sprefix=%E5%8C%96%E7%B2%A7%E6%B0%B4%2Caps%2C203&ref=nb_sb_noss_1"+ page_str
		// methods
		info = scrape.GetHTMLTitle(c, url, &product, &info, &id)
	}

	scrape.SaveFileJson("scrape.json", &info)
}