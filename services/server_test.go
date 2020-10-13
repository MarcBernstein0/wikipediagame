package services

import "testing"

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
