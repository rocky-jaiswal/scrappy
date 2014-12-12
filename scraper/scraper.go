package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Scraper struct {
	url      string
	document *goquery.Document
}

func NewScraper(url string) *Scraper {
	s := new(Scraper)
	s.url = url
	s.document = s.getDocument()
	return s
}

func (s *Scraper) Find(selector string) []string {
	selection := make([]string, 10)
	s.document.Find(selector).Each(func(i int, s *goquery.Selection) {
		selection = append(selection, s.Text())
	})
	return selection
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
