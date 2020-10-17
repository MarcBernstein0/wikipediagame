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
	fmt.Println(links)
}
