package main

import (
	scraper "./scraper"
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
)

func main() {
	printWelcome()

	years := []string{"2009", "2010", "2011", "2012", "2014"}
	for _, year := range years {
		ch := make(chan []string)
		go scrape("http://en.wikipedia.org/wiki/List_of_Bollywood_films_of_"+year, "table.wikitable i a", ch)
		selection := <-ch
		printMovies(selection)
	}
}

func printWelcome() {
	ct.ChangeColor(ct.Magenta, true, ct.White, true)
	fmt.Println("Welcome to Scrappy!")
	fmt.Println("===================")
}

func scrape(url string, selector string, ch chan []string) {
	scraper := scraper.NewScraper(url)
	selection := scraper.Find(selector)
	ch <- selection
}

func printMovies(movies []string) {
	ct.ChangeColor(ct.Blue, true, ct.White, false)
	for _, movie := range movies {
		fmt.Println(movie)
	}
	fmt.Println("___________________")
}
