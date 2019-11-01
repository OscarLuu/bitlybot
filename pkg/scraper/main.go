package scraper

import (
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func ScrapeWebPage(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", err
	}
	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				return "", err
			}
		}
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenType = tokenizer.Next()
				scrape := tokenizer.Token().Data
				return scrape, err
			}
		}
	}
}
