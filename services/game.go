package services

import (
	"errors"
	"fmt"
	"net/url"
)

// Link struct tells what link you are on
// and what is the pervious links so we know how we got here
type Link struct {
	*url.URL
	pastURLs []*url.URL
}

// Start func starts the game
func Start(urlString string) (Link, error) {
	startingURL, err := url.Parse(urlString)
	if err != nil {
		return Link{}, err
	} else if startingURL.Scheme != "https" {
		return Link{}, errors.New("Initial url scheme not https")
	} else if startingURL.Host != "en.wikipedia.org" {
		return Link{}, errors.New("Not a wikipedia link")
	}
	startingLink := Link{
		URL: startingURL,
	}
	links, err := GetLinks(startingLink)
	if err != nil {
		return Link{}, err
	}
	for i, link := range links {
		fmt.Printf("Pass link: %+v\n", link.pastURLs)
		fmt.Printf("Link %d:\n\tscheme: %s\n\thost: %s\n\tpath: %s\n",
			i,
			link.Scheme,
			link.Host,
			link.Path,
		)
	}
	return links[0], nil
}
