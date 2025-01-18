package main

import (
	"github.com/gocolly/colly"
)

type Stock struct{
company, price, change string
}

func main(){
	ticker := []string{
		"NVDA",
		"AVGO",
		"IBM",
		"DIS",
		"PANW",
		"AAPL",
		"MSFT",
		"GOOG",
		"TSLA",
		"AB",
		"ETH",
		"BTC",

	}

	stocks := []Stock{}
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement){
		stock := Stock{}
		stock.company = e.ChildText('h1')
		fmt.Println("Company:", stock.company)
		stock.price = e.ChildTExt("fin-streamer[data-field='regularMarketPrice]")
		fmt.Println("Price:", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change:", stock.change)

		stocks = append(stocks, stock)
	})
	c.Visit("https://finance.yahoo.com/quote/" + t + "/")

}