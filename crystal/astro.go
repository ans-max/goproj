package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	BaseURLAstroLogyCom = "https://www.astrology.com/horoscope/daily"
)

func FromAstroCom(sign string, day string) string {
	var prophecy string
	AstroURL := BaseURLAstroLogyCom + "/" + day + "/" + sign + ".html"
	res, err := http.Get(AstroURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//Load the html doc to qoquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".grid.grid-right-sidebar.primis-rr").Each(func(i int, s *goquery.Selection) {
		x := s.Find("p").Has("span").Text()
		prophecy = strings.TrimSpace(strings.Split(x, ":")[1])
	})
	return prophecy
}
