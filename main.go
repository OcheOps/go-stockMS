package main

import (
	"github.com/gocolly/colly"
)


type Stock struct {
	company,price,change string 
}


func main() 
{

	ticker := []string{
		"AAPL",
		"GOOGL",
		"AMZN",
		"MSFT",
		"FB",
		"TSLA",
		"NVDA",
		"NFLX",
		"VZ",
		"AMD",
		"INTC",
		"IBM",
		"ORCL",}
		stocks := []Stock{}
			c := colly.NewCollector()
			c.OnRequest(func(r *colly.Request){
				println("Visiting", r.URL.String())
	
			})
			c.OnError(func(r *colly.Response, err error){
				log.println("something went wrong", err)
	
	
			})
			c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
			
				stock := Stock{}
		
				stock.company = e.ChildText("h1")
				fmt.Println("Company:", stock.company)
				stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
				fmt.Println("Price:", stock.price)
				stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
				fmt.Println("Change:", stock.change)
		
				stocks = append(stocks, stock)
			})
			c.Wait()
			for_, t := range ticker{
				c.Visit("https://finance.yahoo.com/quote/" + t)
	
	
	}
	
}

	