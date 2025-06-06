# goscrape

This library is a simple web scraping tool. The main function is `Scrape`, which allows you to extract information from web pages.

## Usage

To use the scraper, call the `Scrape` function with the following parameters:

- `url` (string): The target URL to start scraping from.
- `depth` (int): The maximum depth to traverse links from the initial URL. A depth of 0 means only the initial URL will be scraped, a depth of 1 means the initial URL and all links found on it will be scraped, and so on.
- `linkLimit` (int): The maximum number of links to extract from each page. This helps control the scope of the crawl.

Example:

```go
package main

import (
	"fmt"

	"github.com/augustinlss/goscrape"
)

func main() {
	fmt.Println("Starting goscrape example...")

	url := "https://www.bbc.com/news/uk-england-manchester-67124061"

	res := goscrape.Scrape(url, 2, 5)

	fmt.Println(res)
}
```

Keep in mind this is a Fuzzy web scraper, meaning it will output unstructured, general information about any page you want to learn about.
