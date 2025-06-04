package goscrape

import "github.com/PuerkitoBio/goquery"

func ExtractTitle(doc *goquery.Document) string {
	title := doc.Find("title").First().Text()

	return title

}

func ExtractParagraphs(doc *goquery.Document) []string
