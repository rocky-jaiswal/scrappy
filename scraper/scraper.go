package scrappy

import(
        "fmt"
        "net/http"
        "github.com/PuerkitoBio/goquery"
)

type Scraper struct {
        url string
        document *goquery.Document
}

func NewScraper(url string) *Scraper {
        s := new(Scraper)
        s.url = url
        s.document = s.getDocument()
        return s
}

func (s *Scraper) Find(selector string) {
        s.document.Find(selector).Each(func(i int, s *goquery.Selection) {
                fmt.Printf("%s\n", s.Text())
        })
}

func (s *Scraper) getDocument() *goquery.Document {
        resp := s.getResponse()
        defer resp.Body.Close()
        doc, err := goquery.NewDocumentFromResponse(resp)
        if err != nil {
                panic(err)
        }
        return doc
}

func (s *Scraper) getResponse() *http.Response {
        resp, err := http.Get(s.url)
        if err != nil {
                panic(err)
        }
        return resp
}
