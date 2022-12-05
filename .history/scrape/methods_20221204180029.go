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
	cnt := 0
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		cnt++
		name := e.DOM.Text()
		url, _ := e.DOM.Find("a").Attr("href")
		product := productInfo {
			Name: name,
			URL: url,
		}
		products = append(products, product)
		i.Article = products
	})
	c.OnHTML("div", func(e *colly.HTMLElement) {
		cnt++
		price := e.DOM.Find("span").AddClass("a-price-whole").Text()
		star := e.DOM.Find("i").AddClass("a-icon-alt").Text()
		product := productInfo {
			Price: price,
			Star: star,
		}
		products = append(products, product)
		i.Article = products
		fmt.Println("loop count: ", cnt)
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