package main

import "github.com/sRRRs-7/colly_scraper/scrape"


func main() {
	c := scrape.NewScraper("https://github.com/gocolly/colly")
		// method
	getHTMLTitle(c, url)
}