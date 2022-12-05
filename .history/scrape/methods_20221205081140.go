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
	c.OnHTML("h2.a-size-mini a-spacing-none a-color-base s-line-clamp-2", func(e *colly.HTMLElement) {
		name := e.DOM.Text()
		url, _ := e.DOM.Find("a").Attr("href")
		price := e.Attr("a-price-whole")
		star := e.Attr("a-section a-spacing-none a-spacing-top-micro")
		image := e.ChildAttr("img", "s-image")
		product := productInfo {
			Name: name,
			URL: "https://www.amazon.co.jp" + url,
			Price: price,
			Star: star,
			Image: image,
		}
		products = append(products, product)
		i.Article = products
	})
	// request
	c.OnRequest(func(r *colly.Request) {
		i.URL = r.URL.String()
	})
	// response
	c.OnResponse(func(r *colly.Response){
		i.StatusCode = r.StatusCode
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