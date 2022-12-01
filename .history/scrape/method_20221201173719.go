packge scrape

func getHTML(c *colly.Collector) {
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("scrape element: ", e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visit url: ", r.URL.String())
	})
}