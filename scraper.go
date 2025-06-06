package goscrape

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/augustinlss/goscrape/utils"
)

func Scrape(url string, depth int, linkLimit int) *[]ScrapedPage {
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
	firstPage := FuzzyParse(doc, url, linkLimit)
	var res []ScrapedPage
	res = append(res, *firstPage)

	queue := NewQueue[ScrapedPage]()
	queue.Enqueue(*firstPage)

	// TODO do BFS here to search through all the pages at a certain depth
	for queue.Length > 0 {
		current_page, _ := queue.Dequeue()
		current_depth := current_page.Depth

		// if reached max depth then we don't explore this path any futher
		if current_depth == depth || len(current_page.Links) == 0 {
			continue
		}

		for _, link := range current_page.Links {
			resp, status, err := utils.Fetch(link)

			fmt.Println(link)

			if err != nil {
				log.Println("Error crawling link...")
				continue
			}

			if status != http.StatusOK {
				log.Printf("Unexpected status code when crawling link: %s, got code: %d", link, status)
				continue
			}

			doc, err := goquery.NewDocumentFromReader(resp.Body)

			if err != nil {
				log.Println("Error fetching link content...")
				continue
			}
			discoverdPage := FuzzyParse(doc, link, linkLimit)
			discoverdPage.Depth = current_depth + 1
			queue.Enqueue(*discoverdPage)
			res = append(res, *discoverdPage)
		}

	}

	return &res
}
