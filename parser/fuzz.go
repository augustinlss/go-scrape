package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/augustinlss/goscrape"
)

func FuzzyParse(doc *goquery.Document) *goscrape.ScrapedPage {
	return &goscrape.ScrapedPage{
		URL: extract_url(doc.Url),
	}
}
