package main

import(
        "fmt"
        "net/http"
        // "io/ioutil"
        "github.com/daviddengcn/go-colortext"
        "github.com/PuerkitoBio/goquery"
)

func main(){
        ct.ChangeColor(ct.Magenta, true, ct.White, false)
        fmt.Println("Welcome to Scrappy!")
        fmt.Println("===================")
        ct.ChangeColor(ct.Blue, true, ct.White, false)

        resp, err := http.Get("http://rockyj.in")
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()

        //for debugging
        // body, _ := ioutil.ReadAll(resp.Body)
        // fmt.Println(string(body))

        doc, err := goquery.NewDocumentFromResponse(resp)
        if err != nil {
                panic(err)
        }

        doc.Find(".summary h3 a").Each(func(i int, s *goquery.Selection) {
                fmt.Printf("%d: %s\n", i, s.Text())
        })
}
