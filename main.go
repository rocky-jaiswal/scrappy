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
        ch := make(chan []string)
        go scrape("http://en.wikipedia.org/wiki/List_of_Bollywood_films_of_2014", "table.wikitable i a", ch)
        selection := <- ch
        fmt.Println(selection)
}

func scrape(url string, selector string, ch chan []string) {
        scraper := scraper.NewScraper(url)
        selection := scraper.Find(selector)
        ch <- selection
}
