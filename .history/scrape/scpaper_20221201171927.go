package scrape

import (
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/sRRRs-7/colly_scraper/utils"
)

func Scraper() {
	// initialize
	c := colly.NewCollector
	c.colly.UserAgent("user-agent")
	c.colly.AllowURLRevisit()
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

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", utils.RandomString())
	})


}