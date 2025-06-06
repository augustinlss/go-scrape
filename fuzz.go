package goscrape

import (
	"github.com/PuerkitoBio/goquery"
)

func FuzzyParse(doc *goquery.Document, url string) *ScrapedPage {
	return &ScrapedPage{
		URL:        url,
		Title:      ExtractTitle(doc),
		Paragraphs: ExtractParagraphs(doc),
		Links:      ExtractLinks(doc),
		Depth:      0,
	}
}
