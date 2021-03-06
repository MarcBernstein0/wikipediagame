package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	scheme = "https"
	host   = "en.wikipedia.org"
)

func getValidWikiURL(urlString string) (*url.URL, bool) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, false
	}
	if url.Host == "" && url.Scheme == "" && strings.Contains(url.Path, "/wiki/") {
		url.Scheme = scheme
		url.Host = host
		return url, true
	}
	return nil, false
}

func getInternalURLs(body io.Reader) ([]*url.URL, error) {
	var urls []*url.URL
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("Error with creating doc %v", err)
	}
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if url, isValid := getValidWikiURL(href); isValid {
				urls = append(urls, url)
			}
		}
	})
	return urls, err
}

// GetLinks takes in a url string and returns a list of *url.URL or err if
// an error occured
func GetLinks(link Link) ([]Link, error) {
	urlString := link.String()
	pastURLs := append(link.pastURLs, link.URL)
	response, err := http.Get(urlString)
	if err != nil {
		return nil, fmt.Errorf("Get request failed:\n%v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("status code error for link %s: %d %s", urlString, response.StatusCode, response.Status)
	}

	urls, err := getInternalURLs(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Go query failed:\n%v", err)
	}
	if urls == nil {
		return nil, fmt.Errorf("No links found on page %v", urlString)
	}
	var links []Link
	for _, parsedURL := range urls {
		newLink := Link{
			URL:      parsedURL,
			pastURLs: pastURLs,
		}
		links = append(links, newLink)
	}

	return links, nil
}
