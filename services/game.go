package services

import (
	"fmt"
	"net/url"
)

// Link struct tells what link you are on
// and what is the pervious links so we know how we got here
type Link struct {
	*url.URL
	pastURLs []*url.URL
}

func isValidWikipediaURL(url *url.URL) bool {
	if url.Scheme != "https" || url.Host != "en.wikipedia.org" {
		return false
	}
	return true
}

// Start func starts the game
func Start(startingRawURL, endingRawURL string) (Link, error) {
	startingURL, err := url.Parse(startingRawURL)
	if err != nil {
		return Link{}, err
	}
	endingURL, err := url.Parse(endingRawURL)
	if err != nil {
		return Link{}, err
	}
	if !isValidWikipediaURL(startingURL) {
		return Link{}, fmt.Errorf("startingURL not valid wikipedia url:\t%v", startingURL)
	}
	if !isValidWikipediaURL(endingURL) {
		return Link{}, fmt.Errorf("startingURL not valid wikipedia url:\t%v", startingURL)
	}

	startingLink := Link{
		URL: startingURL,
	}
	fmt.Println(len(startingLink.pastURLs))
	if *startingURL == *endingURL {
		return startingLink, nil
	}

	links, err := GetLinks(startingLink)
	if err != nil {
		return Link{}, err
	}
	// for i, link := range links {
	// 	fmt.Printf("Pass link: %+v\n", link.pastURLs)
	// 	fmt.Printf("Link %d:\n\tscheme: %s\n\thost: %s\n\tpath: %s\n",
	// 		i,
	// 		link.Scheme,
	// 		link.Host,
	// 		link.Path,
	// 	)
	// }
	return links[0], nil
}
