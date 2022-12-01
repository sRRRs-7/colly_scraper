package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getHTMLTitle(c *colly.Collector) {
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("element: ", e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("request: ", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response){
		fmt.Println("response: ", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error){
		fmt.Println("error: ", r.StatusCode, err)
	})
}