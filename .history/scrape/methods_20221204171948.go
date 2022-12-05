package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetHTMLTitle(c *colly.Collector, url string) {
	// info articles struct instance
	i := &info{}
	articles := make([]articleInfo, 0, 4)
	// html element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		i.Title = e.Text
		fmt.Println("element: ", e.Text)
	})
	// article scraping
	cnt := 0
	c.OnHTML("div", func(e *colly.HTMLElement) {
		cnt++
		title := e.ChildText("h2")
		url, _ := e.DOM.Find("a").Attr("href")
		name := e.DOM.Find("h2").Attr("span")
		article := articleInfo {
			Title: title,
			URL: url,
		}
		articles = append(articles, article)
		i.Article = articles
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

	saveFileJson("scrape.json", i)
}

func GetHTML(c *colly.Collector, url string) {
	arr := []string{}

	// html element
	c.OnHTML("a", func(e *colly.HTMLElement) {
		arr = append(arr, e.Text)
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
	// wait all threads
	c.Wait()

	outputFile("scrape.txt", arr)
}