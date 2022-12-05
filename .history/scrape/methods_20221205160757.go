package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/sRRRs-7/colly_scraper/utils"
)

func GetHTMLTitle(c *colly.Collector, url string, products *[]ProductInfo, info *Info) Info {
	// info articles struct instance
	product := &ProductInfo{}
	// html element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		info.Title = e.Text
	})
	// product scraping
	id := 0
	c.OnHTML("div.a-spacing-base", func(e *colly.HTMLElement) {
		id ++
		product.ID = id
		product.Name = e.DOM.Find("h2").Children().Text()
		url, _ := e.DOM.Find("a").Attr("href")
		product.URL = "https://www.amazon.co.jp" + url
		product.Price = utils.RegexRetriever(`¥\d+,\d+`, e.DOM.Text())
		product.Star = utils.RegexRetriever(`\d*\.\d*\s*\w*\s*\w+\s*\d*\s*stars`, e.DOM.Text())
		img, _ := e.DOM.Find("img").Attr("src")
		product.Image = img

		products = append(products, product)
		if product.Name == "" {
			id--
		} else {
			info.Article = products
		}

	})
	// request
	c.OnRequest(func(r *colly.Request) {
		info.URL = r.URL.String()
	})
	// response
	c.OnResponse(func(r *colly.Response){
		info.StatusCode = r.StatusCode
	})
	// error
	c.OnError(func(r *colly.Response, err error){
		fmt.Println("error: ", r.StatusCode, err)
	})
	// destination
	c.Visit(url)
	// wait all threads
	c.Wait()

	return *info
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

	OutputFile("scrape.txt", arr)
}