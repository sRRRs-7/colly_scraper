package main

import (
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.UserAgent("user-agent"),
		colly.AllowURLRevisit()
	)
	c.OnHTML()
}