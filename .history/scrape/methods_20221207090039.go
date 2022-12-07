package scrape

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/sRRRs-7/colly_scraper/utils"
)

func GetAmazon(c *colly.Collector, url string, products *[]ProductInfo, info *Info) Info {
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
		product.Name = e.DOM.Find("h2").Children().Text()

		url, _ := e.DOM.Find("a").Attr("href")
		product.URL = "https://www.amazon.co.jp" + url

		p := e.DOM.Find("span.a-price-whole").Text()
		product.Price = utils.PriceReplace(p)

		product.Star = e.DOM.Find("span.a-icon-alt").Text()

		img, _ := e.DOM.Find("img").Attr("src")
		product.Image = img

		if product.Name != "" || product.Price != "" {
			*products = append(*products, product)
			info.ProductList = *products
		}
		// distinct logic
		uniq := distinct(info.ProductList)
		*products = uniq
		// array sort logic
		sort.Slice(info.ProductList, func(i, j int) bool {
			i1, _ := strconv.Atoi(info.ProductList[i].Price)
			i2, _ := strconv.Atoi(info.ProductList[j].Price)
			return i1 < i2
		})
		// product count
		info.ProductCount = len(*products)
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

// distinct function
func distinct(arr []ProductInfo) []ProductInfo {
	list := make(map[string]struct{})
	uniq := []ProductInfo{}
	// id := 0
	for _, e := range arr {
		list[e.Name] = struct{}{}	// object distinct
	}
	for i := range list {
		fmt.Println(i)
	}
	return uniq
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