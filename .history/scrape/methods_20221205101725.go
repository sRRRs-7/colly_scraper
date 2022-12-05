package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetHTMLTitle(c *colly.Collector, url string) {
	// info articles struct instance
	i := &info{}
	products := make([]productInfo, 0, 4)
	// html element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		i.Title = e.Text
	})
	// product scraping
	c.OnHTML("div.sg-col-4-of-24 sg-col-4-of-12 s-result-item s-asin sg-col-4-of-16 AdHolder sg-col s-widget-spacing-small sg-col-4-of-20", func(e *colly.HTMLElement) {
		product := productInfo {}
		e.ForEach("div.a-section aok-relative s-image-square-aspect", func(_ int, h *colly.HTMLElement) {
			fmt.Println(e.ChildText("img"))
		})
		products = append(products, product)
		i.Article = products
	})
	// request
	req := 0
	c.OnRequest(func(r *colly.Request) {
		req++
		i.URL = r.URL.String()
		fmt.Println("request: ", req)
	})
	// response
	res := 0
	c.OnResponse(func(r *colly.Response){
		res++
		i.StatusCode = r.StatusCode
		fmt.Println("response: ", res)
	})
	// error
	c.OnError(func(r *colly.Response, err error){
		fmt.Println("error: ", r.StatusCode, err)
	})
	// destination
	c.Visit(url)
	// wait all threads
	c.Wait()

	saveFileJson("scrape.json", i)
}


func GetHTML(c *colly.Collector, url string) {
	arr := []string{}

	// html element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		arr = append(arr, e.Text)
	})
	// request
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("request: ", r.URL.String())
	})
	// response
	c.OnResponse(func(r *colly.Response){
		// fmt.Println("response: ", r.StatusCode)
	})
	// error
	c.OnError(func(r *colly.Response, err error){
		// fmt.Println("error: ", r.StatusCode, err)
	})
	// destination
	c.Visit(url)
	// wait all threads
	c.Wait()

	outputFile("scrape.txt", arr)
}