package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Stock struct {
	Company     string
	Price       float64
	Change      float64
	PriceChange float64
}

func main() {
	tickers := []string{
		"NVO",
		"TSM",
		"TSLA",
		"NVDA",
		
	}

	var stocks []Stock

	
	collector := colly.NewCollector()

	
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	
	collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	
	collector.OnHTML("body", func(e *colly.HTMLElement) {
		var stock Stock

		
		stock.Company = path.Base(e.Request.URL.Path)
		fmt.Println("Company:", stock.Company)

		
		priceStr := e.ChildText(fmt.Sprintf("fin-streamer[data-symbol='%s'][data-field='regularMarketPrice']", stock.Company))
		priceStr = strings.TrimSpace(priceStr)
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Printf("Error parsing price for %s: %v\n", stock.Company, err)
			return
		}
		stock.Price = price
		fmt.Println("Price:", stock.Price)

		
		changeStr := e.ChildText(fmt.Sprintf("fin-streamer[data-symbol='%s'][data-field='regularMarketChangePercent']", stock.Company))
		changeStr = strings.TrimSpace(changeStr)
		changeStr = strings.Trim(changeStr, "()") 
		changeStr = strings.TrimSuffix(changeStr, "%") 
		change, err := strconv.ParseFloat(changeStr, 64)
		if err != nil {
			fmt.Printf("Error parsing percentage change for %s: %v\n", stock.Company, err)
			return
		}
		stock.Change = change
		fmt.Println("Percentage Change:", stock.Change)

		
		priceChangeStr := e.ChildText(fmt.Sprintf("fin-streamer[data-symbol='%s'][data-field='regularMarketChange']", stock.Company))
		priceChangeStr = strings.TrimSpace(priceChangeStr)
		priceChangeStr = strings.TrimPrefix(priceChangeStr, "+") 
		priceChange, err := strconv.ParseFloat(priceChangeStr, 64)
		if err != nil {
			fmt.Println("Error parsing price change:", err)
			return
		}
		stock.PriceChange = priceChange
		fmt.Println("Price Change:", stock.PriceChange)

		
		stocks = append(stocks, stock)
	})

	
	for _, ticker := range tickers {
		collector.Visit("https://finance.yahoo.com/quote/" + ticker + "/")
	}
	collector.Wait()

	
	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{"Company", "Price", "Percentage Change", "Price Change"}
	writer.Write(headers)

	for _, stock := range stocks {
		record := []string{
			stock.Company,
			fmt.Sprintf("%.2f", stock.Price),
			fmt.Sprintf("%.2f%%", stock.Change),
			fmt.Sprintf("%.2f", stock.PriceChange),
		}
		writer.Write(record)
	}
	writer.Flush()

	fmt.Println("Stock data written to stocks.csv")
}
