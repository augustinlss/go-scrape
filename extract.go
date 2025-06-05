package goscrape

import "github.com/PuerkitoBio/goquery"

func ExtractTitle(doc *goquery.Document) string {
	title := doc.Find("title").First().Text()

	return title

}

func ExtractParagraphs(doc *goquery.Document) []string {
	paragraphsSelection := doc.Find("p")

	var paragraphs []string
	paragraphsSelection.Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	return paragraphs
}
