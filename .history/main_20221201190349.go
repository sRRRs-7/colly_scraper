package main

import "github.com/sRRRs-7/colly_scraper/scrape"


func main() {
	url := "https://news.yahoo.com/?guccounter=1&guce_referrer=aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8&guce_referrer_sig=AQAAAEwp8STpRCql7Hmo2i1F0e2ThdrOJdWtmjCFifQmcMZbE-8PwyowfIyYRtwgNKp6bYHssuU0vpSM8WmiF-D57EBo-NXJ6-bAGbhz0MYd72fJ1tUytKwDlu6KM3aUuPxbI1M3JUjxBVQxSMCMpgyQBzBkp0tLZmJ7yB3ftzSg3X3v"
	// instance
	c := scrape.NewScraper()
	// methods
	scrape.GetHTMLTitle(c, url)
	scrape.GetHTML(c, url)
}