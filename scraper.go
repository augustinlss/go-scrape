package goscrape

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/augustinlss/goscrape/utils"
)

func Scrape(url string, depth int16) *ScrapedPage {
	resp, status, err := utils.Fetch(url)

	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error fetching url content...")
		return nil
	}

	if status != http.StatusOK {
		log.Fatalf("Unexpected status code when fetching url content: %d", status)
		return nil
	}

	//TODO need to ingest the retuned response package...
	//
	// not really sure how i am going to do this.
	// do i just store all the entire response in a single buffer?
	// then parse it through some smart html parser?

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatalf("Error fetching url content...")
		return nil
	}

	return FuzzyParse(doc, url)

}
