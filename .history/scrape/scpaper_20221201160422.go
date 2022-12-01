package scrape

import (
	"github.com/gocolly/colly"
	"github.com/sRRRs-7/colly_scraper/utils"
)

func Scraper() {
	c := colly.NewCollector(
		colly.UserAgent("user-agent"),
		colly.AllowURLRevisit(),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", utils.RandomString())
	})


}