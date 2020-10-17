package main

import (
	"fmt"
	"log"

	"github.com/MarcBernstein0/wikipediagame/services"
)

func main() {
	// fmt.Println(links)
	links, err := services.GetLinks("https://en.wikipedia.org/wiki/Go_(programming_language)")
	if err != nil {
		log.Fatal(err)
	}
	for i, link := range links {
		fmt.Printf("Link %d:\n\tscheme: %s\n\thost: %s\n\tpath: %s\n",
			i,
			link.Scheme,
			link.Host,
			link.Path,
		)
	}
}
