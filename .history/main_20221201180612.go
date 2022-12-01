package main

import "github.com/sRRRs-7/colly_scraper/scrape"


func main() {
	url := "https://github.com/gocolly/colly"
	c := scrape.NewScraper()
	// methods
	scrape.GetHTMLTitle(c, url)
}