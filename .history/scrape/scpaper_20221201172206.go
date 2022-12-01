package scrape

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func Scraper() {
	// initialize
	c := colly.NewCollector(
		// colly.UserAgent("user-agent"),
		// colly.AllowURLRevisit(),
	)
	// extensions
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	// http client configuration
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer {
			Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns: 100,
		IdleConnTimeout: 90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	c.OnResponse(func(r *colly.Request) {
		log.Println(string(r.Body))
	})


}