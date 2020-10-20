package services

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNonInternalWikiString(t *testing.T) {
	urlString := "https://en.wikipedia.org/wiki/Go_(programming_language)"
	_, valid := getValidWikiURL(urlString)
	if valid == true {
		t.Errorf("Expected false because link is not internal wiki link: initial link: %v expected: %v got: %v", urlString, true, valid)
	}
}

func TestInternalWikiString(t *testing.T) {
	urlString := "/wiki/Go_(programming_language)"
	resultURL, valid := getValidWikiURL(urlString)
	// get if valid internal link
	if valid != true {
		t.Errorf("Expected true because link is internal wiki link: initial link: %v expected: %v got: %v", urlString, true, valid)
	}
	// check correct scheme
	if resultURL.Scheme != "https" {
		t.Errorf("Scheme should be set to https: expected: %v got: %v", "https", resultURL.Scheme)
	}
	// check correct host
	if resultURL.Host != "en.wikipedia.org" {
		t.Errorf("Host should be set to wikipedia: expected: %v got: %v", "en.wikipedia.org", resultURL.Host)
	}
	// check if link works
	res, err := http.Get(resultURL.String())
	if err != nil {
		t.Errorf("Error with connecting to website: link: %s err: %v", resultURL, err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Status code not ok: expected: %d got: %d", http.StatusOK, status)
	}

}

func TestURLList(t *testing.T) {
	address, err := url.Parse("https://en.wikipedia.org/wiki/Go_(programming_language)")
	if err != nil {
		t.Error("URL could not be parsed")
	}
	testLink := Link{
		URL: address,
	}
	result, err := GetLinks(testLink)
	if err != nil {
		t.Errorf("Got back error: %v\n", err)
	}
	if result == nil {
		t.Errorf("No links found")
	}

	// for i, link := range result {
	// 	fmt.Printf("Pass link: %+v\n", link.pastURLs)
	// 	fmt.Printf("Link %d:\n\tscheme: %s\n\thost: %s\n\tpath: %s\n",
	// 		i,
	// 		link.Scheme,
	// 		link.Host,
	// 		link.Path,
	// 	)
	// }
}
