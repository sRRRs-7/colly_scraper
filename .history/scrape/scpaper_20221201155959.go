package scrape

import "github.com/gocolly/colly"

func Scraper() {
	c := colly.NewCollector(
		colly.UserAgent("user-agent"),
		colly.AllowURLRevisit(),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", utilsRandomString())
	})
}