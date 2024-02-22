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
}
		stocks := []Stock{}
		c := colly.NewCollector()
		c.OnRequest(func(r *colly.Request){

		})

}
