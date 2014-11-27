package scrappy

import(
        "fmt"
        "net/http"
        "github.com/PuerkitoBio/goquery"
)

type Scraper struct {
        url string
}

func NewScraper(url string) *Scraper {
        s := Scraper{url}
        return &s
}

func (s *Scraper) GetDocument() *goquery.Document {
        resp := s.getResponse()
        defer resp.Body.Close()
        doc, err := goquery.NewDocumentFromResponse(resp)
        if err != nil {
                panic(err)
        }
        return doc
}

func (s *Scraper) Find(doc *goquery.Document, selector string) {
        doc.Find(selector).Each(func(i int, s *goquery.Selection) {
                fmt.Printf("%d: %s\n", i, s.Text())
        })
}

func (s *Scraper) getResponse() *http.Response {
        resp, err := http.Get(s.url)
        if err != nil {
                panic(err)
        }
        return resp
}
