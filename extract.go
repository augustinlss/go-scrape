package goscrape

import "github.com/PuerkitoBio/goquery"

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

func ExtractLinks(doc goquery.Document) []string {
	linkSelections := doc.Find("a")

	var links []string

	linkSelections.Each(func(i int, s *goquery.Selection) {
		links = append(links, s.Text())
	})

	return links
}
