package main

import(
        "fmt"
        "github.com/daviddengcn/go-colortext"
)

func main(){
        ct.ChangeColor(ct.Magenta, true, ct.White, false)
        fmt.Println("Welcome to Scrappy!")
        fmt.Println("===================")
        ct.ChangeColor(ct.Blue, true, ct.White, false)

        scraper := NewScraper("http://rockyj.in")
        doc := scraper.GetDocument()
        scraper.Find(doc, ".summary h3 a")
}
