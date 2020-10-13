package services

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetAllLinks function
func GetAllLinks(urlString string) ([]string, error) {
	var urls []string

	response, err := http.Get(urlString)
	if err != nil {
		return nil, fmt.Errorf("Get request failed:\n%v", err)
	}
	document, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, fmt.Errorf("Go query failed:\n%v", err)
	}
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			fmt.Println(href)
			urls = append(urls, href)
		}
	})

	return urls, nil
}
