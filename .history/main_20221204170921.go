package main

import "github.com/sRRRs-7/colly_scraper/scrape"


func main() {
	// input url for analytics
	url := "https://www.amazon.co.jp/-/en/"
	// instance
	c := scrape.NewScraper()
	// methods
	scrape.GetHTML(c, url)
	scrape.GetHTMLTitle(c, url)
}