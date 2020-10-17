package services

import (
	"net/http"
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
	urlString := "https://en.wikipedia.org/wiki/Go_(programming_language)"
	result, err := GetLinks(urlString)
	if err != nil {
		t.Errorf("Got back error: %v\n", err)
	}
	if result == nil {
		t.Errorf("No links found")
	}
}
