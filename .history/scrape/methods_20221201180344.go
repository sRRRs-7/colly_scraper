package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getHTMLTitle(c *colly.Collector, url string) {
	// html element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("element: ", e.Text)
	})
	// request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("request: ", r.URL.String())
	})
	// response
	c.OnResponse(func(r *colly.Response){
		fmt.Println("response: ", r.StatusCode)
	})
	// error
	c.OnError(func(r *colly.Response, err error){
		fmt.Println("error: ", r.StatusCode, err)
	})
	// destination
	c.Visit(url)
}