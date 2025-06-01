package goscrape

type ScrapedPage struct {
	URL        string   `json:"url"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Links      []string `json:"links"`
	StatusCode int      `json:"status_code"`
	Error      string   `jsonL:"error"`
	Depth      int      `json:"depth"`
	MetaData   string   `json:"metadata,omitempty"`
}
