package goscrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractTitle(doc *goquery.Document) string {
	title := doc.Find("title").First().Text()

	return title

}

func ExtractParagraphs(doc *goquery.Document) []string {
	paragraphSelections := doc.Find("p")

	var paragraphs []string
	paragraphSelections.Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	return paragraphs
}

func ExtractLinks(doc *goquery.Document, limit int) []string {
	linkSelections := doc.Find("a")

	var links []string
	var limit_so_far int

	linkSelections.Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if limit_so_far == limit {
			return
		}

		if exists && strings.HasPrefix(href, "http") {
			limit_so_far++
			links = append(links, href)
		}
	})

	return links

}
