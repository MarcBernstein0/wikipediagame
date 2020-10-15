package main

import "github.com/MarcBernstein0/wikipediagame/services"

func main() {
	// fmt.Println(links)
	services.GetLinks("https://en.wikipedia.org/wiki/Go_(programming_language)")
}
