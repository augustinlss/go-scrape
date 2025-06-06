package goscrape

import (
	"github.com/PuerkitoBio/goquery"
)

func FuzzyParse(doc *goquery.Document, url string, linkLimit int) *ScrapedPage {
	return &ScrapedPage{
		URL:        url,
		Title:      ExtractTitle(doc),
		Paragraphs: ExtractParagraphs(doc),
		Links:      ExtractLinks(doc, linkLimit),
		Depth:      0,
	}
}
