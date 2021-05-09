package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
	// colly.AllowedDomains("olx.pl"),
	)

	c.OnHTML("li", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.Visit("https://www.olx.pl/d/oferta/kawalerka-z-kuchnia-poznan-lazarz-36-m-CID3-IDJBpjV.html#198e881dab;promoted")
}
