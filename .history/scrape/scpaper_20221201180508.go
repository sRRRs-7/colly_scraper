package scrape

import (
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func NewScraper()*colly.Collector {
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

	return c
}

