package main

import(
        "fmt"
       ct "github.com/daviddengcn/go-colortext"
       scraper  "github.com/rocky-jaiswal/scrappy/scraper"
)

func main(){
        ct.ChangeColor(ct.Magenta, true, ct.White, false)
        fmt.Println("Welcome to Scrappy!")
        fmt.Println("===================")
        ct.ChangeColor(ct.Blue, true, ct.White, false)

        //for my personal site
        //scraper := scraper.NewScraper("http://rockyj.in")
        //scraper.Find(".summary h3 a")

        //bollywood movies
        scraper := scraper.NewScraper("http://en.wikipedia.org/wiki/List_of_Bollywood_films_of_2014")
        //doc := scraper.GetDocument()
        scraper.Find("table.wikitable i a")
}
