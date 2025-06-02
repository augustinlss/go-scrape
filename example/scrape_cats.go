package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/augustinlss/goscrape/utils"
)

func main() {
	fmt.Println("Starting goscrape example...")

	url := "https://en.wikipedia.org/wiki/Cat"
	resp, status, err := utils.Fetch(url)

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("!! Error fetching url data !!")
		return
	}

	if status != http.StatusOK {
		log.Fatalf("Received non-200 code: %d", status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	htmlContent := string(bodyBytes)
	log.Print(len(htmlContent))

	if len(htmlContent) > 500 {
		fmt.Println(htmlContent[:2000])
	} else {
		fmt.Println(htmlContent)
	}
}
