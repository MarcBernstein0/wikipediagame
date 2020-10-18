package services

import (
	"testing"
)

func TestStart(t *testing.T) {
	res, err := Start("https://en.wikipedia.org/wiki/Go_(programming_language)")
	if err != nil {
		t.Errorf("Error start function, %v", err)
	}
	if res.pastURLs[0].String() != "https://en.wikipedia.org/wiki/Go_(programming_language)" {
		t.Errorf("Not correct first past url: expected: %s but got %s", "https://en.wikipedia.org/wiki/Go_(programming_language)", res.pastURLs[0])
	}
}

func TestWrongURLForStart(t *testing.T) {
	_, err := Start("https://google.com")
	if err.Error() != "Not a wikipedia link" {
		t.Errorf("Another error occured: exptected: %v got: %v", "Not a wikipedia link", err)
	} else if err == nil {
		t.Error("Expected error but no error occured")
	}
}
