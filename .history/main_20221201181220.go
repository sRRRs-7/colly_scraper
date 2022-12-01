package main

import "github.com/sRRRs-7/colly_scraper/scrape"


func main() {
	url := "https://github.com/gocolly/colly"
	// instance
	c := scrape.NewScraper()
	// methods
	scrape.GetHTML()
	scrape.GetHTMLTitle(c, url)
}