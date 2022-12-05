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
	// article scraping
	cnt := 0
	c.OnHTML("div", func(e *colly.HTMLElement) {
		cnt++
		name := e.DOM.Find("h2").Text()
		url, _ := e.DOM.Find("a").Attr("href")
		price, _ := e.DOM.Find("h2").Attr("span")
		description := e.DOM.Find("").AddClass("a-price-whole").Text()
		star := e.DOM.Find("i").AddClass("a-icon-alt").Text()
		product := productInfo {
			Name: name,
			URL: url,
			Price: price,
			Description: description,
			Star: star,
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