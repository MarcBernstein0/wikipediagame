package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func request() {
	// Make HTTP GET request
	response, err := http.Get("https://www.devdungeon.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}

func requestWithTimeOut() {
	//Make HTTP GET request with timeout
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Get("https://www.devdungeon.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}

func requestWithHTTPHeaders() {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", "https://www.devdungeon.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "Not Firefox")

	//Make request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func requestDownloadURL() {
	response, err := http.Get("https://www.devdungeon.com/archive")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create output file
	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}

// Not the best way to do it
func requestSubstringMatching() {
	response, err := http.Get("https://www.devdungeon.com/archive")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Get response body as string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	pageContent := string(dataInBytes)

	// Find a substr
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No element found")
		return
	}

	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	titleStartIndex += 7

	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleStartIndex == -1 {
		fmt.Println("No closing tag found")
		return
	}

	// (Optional)
	// Copy the substring in to a separate variable so the
	// variables with the full document data can be garbage collected
	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	fmt.Printf("Page title: %s\n", pageTitle)
}

func requestRegex() {
	// Make HTTP request
	response, err := http.Get("https://www.devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// Create a regular expression to find comments
	re := regexp.MustCompile("<!--(.|\n)*?-->")
	comments := re.FindAllString(string(body), -1)
	if comments == nil {
		fmt.Println("No matches found")
	} else {
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}

}

func processElement(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func requestParsingGoquery() {
	// Make HTTP request
	response, err := http.Get("https://www.devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)
}

func parsingURL() {
	// Parse a complex URL
	complexURL := "https://www.example.com/path/to/?query=123&this=that#fragment"
	parsedURL, err := url.Parse(complexURL)
	if err != nil {
		log.Fatal(err)
	}

	// Print out URL pieces
	fmt.Println("Scheme: " + parsedURL.Scheme)
	fmt.Println("Host: " + parsedURL.Host)
	fmt.Println("Path: " + parsedURL.Path)
	fmt.Println("Query string: " + parsedURL.RawQuery)
	fmt.Println("Fragment: " + parsedURL.Fragment)

	// Get the query key/values as a map
	fmt.Println("\nQuery values:")
	queryMap := parsedURL.Query()
	fmt.Println(queryMap)

	// Craft a new URL from scratch
	var customURL url.URL
	customURL.Scheme = "https"
	customURL.Host = "google.com"
	newQueryValues := customURL.Query()
	newQueryValues.Set("key1", "value1")
	newQueryValues.Set("key2", "value2")
	customURL.Fragment = "bookmarkLink"
	customURL.RawQuery = newQueryValues.Encode()

	fmt.Println("\nCustom URL:")
	fmt.Println(customURL.String())
}

func requestGetImages() {
	response, err := http.Get("https://www.devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	document, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exits := element.Attr("src")
		if exits {
			fmt.Println(imgSrc)
		}
	})
}

// func main() {
// 	// request()
// 	// requestWithTimeOut()
// 	// requestWithHTTPHeaders()
// 	// requestDownloadURL()
// 	// requestSubstringMatching()
// 	// requestRegex()
// 	// requestParsingGoquery()
// 	parsingURL()
// 	// requestGetImages()
// }
