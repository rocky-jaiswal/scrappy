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

        scraper := scraper.NewScraper("http://rockyj.in")
        doc := scraper.GetDocument()
        scraper.Find(doc, ".summary h3 a")
}
