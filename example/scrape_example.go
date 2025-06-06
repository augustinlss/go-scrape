package main

import (
	"fmt"

	"github.com/augustinlss/goscrape"
)

func main() {
	fmt.Println("Starting goscrape example...")

	url := "https://www.bbc.com/news/uk-england-manchester-67124061"

	res := goscrape.Scrape(url, 2, 5)

	fmt.Println(res)
}
