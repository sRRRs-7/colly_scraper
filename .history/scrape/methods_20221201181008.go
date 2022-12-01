package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetHTMLTitle(c *colly.Collector, url string) {
	i := &info{}

	// html element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		i.Title = e.Text
		fmt.Println("element: ", e.Text)
	})
	// request
	c.OnRequest(func(r *colly.Request) {
		i.URL = r.URL.String()
		fmt.Println("request: ", r.URL.String())
	})
	// response
	c.OnResponse(func(r *colly.Response){
		i.StatusCode = r.StatusCode
		fmt.Println("response: ", r.StatusCode)
	})
	// error
	c.OnError(func(r *colly.Response, err error){
		fmt.Println("error: ", r.StatusCode, err)
	})
	// destination
	c.Visit(url)
	// wait all threads
	c.Wait()

	saveFileJson("scrape.json", p)
}