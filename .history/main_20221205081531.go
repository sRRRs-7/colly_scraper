package main

import "github.com/sRRRs-7/colly_scraper/scrape"


func main() {
	// input url for analytics
	url := "https://www.amazon.co.jp/-/en/s?k=%E4%BA%BA%E3%82%92%E3%83%80%E3%83%A1%E3%81%AB%E3%81%99%E3%82%8B%E3%81%84%E3%81%99&page=2&crid=WZOABIEQ8Z1Z&qid=1670141829&sprefix=%E4%BA%BA%E3%82%92%E3%83%80%E3%83%A1%E3%81%AB%E3%81%99%E3%82%8B%E3%81%84%E3%81%99%2Caps%2C199&ref=sr_pg_2"
	// instance
	c := scrape.NewScraper()
	// methods
	scrape.GetHTML(c, url)
	scrape.GetHTMLTitle(c, url)
}