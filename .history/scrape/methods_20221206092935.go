package scrape

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/sRRRs-7/colly_scraper/utils"
)

func GetHTMLTitle(c *colly.Collector, url string, products *[]ProductInfo, info *Info, id *int) Info {
	// html element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		info.Title = e.Text
	})
	// request
	c.OnRequest(func(r *colly.Request) {
		info.URL = r.URL.String()
	})
	// response
	c.OnResponse(func(r *colly.Response){
		info.StatusCode = r.StatusCode
	})
	// product scraping
	product := ProductInfo{}
	c.OnHTML("div.a-spacing-base", func(e *colly.HTMLElement) {
		*id ++
		product.ID = *id
		product.Name = e.DOM.Find("h2").Children().Text()
		url, _ := e.DOM.Find("a").Attr("href")
		product.URL = "https://www.amazon.co.jp" + url
		p := e.DOM.Find("span.a-price-whole").Text()
		product.Price = utils.PriceReplace(p)
		product.Star = e.DOM.Find("span.a-icon-alt").Text()
		img, _ := e.DOM.Find("img").Attr("src")
		product.Image = img

		if product.Name == "" || product.Price == "" {
			*id--
		} else {
			*products = append(*products, product)
			info.Article = *products
		}
		// array sort logic
		sort.Slice(info.Article, func(i, j int) bool {


			return strconv.Atoi(info.Article[i].Price) < strconv.Atoi(info.Article[j].Price)
		})

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