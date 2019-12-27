package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	BaseURLGaneshaSpeaks = "https://www.ganeshaspeaks.com/horoscopes/"
)

func FromGaneshSpeaks(sign string, day string) string {
	var prophecy, AstroURL string
	if day == "today" {
		AstroURL = BaseURLGaneshaSpeaks + "daily-horoscope/" + sign + "/"
	} else if day == "yesterday" {
		AstroURL = BaseURLGaneshaSpeaks + "yesterday-horoscope/" + sign + "/"
	} else {
		AstroURL = BaseURLGaneshaSpeaks + "tomorrow-horoscope/" + sign + "/"
	}
	res, err := http.Get(AstroURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".row.card-padding-20.container-fluid-xs.margin-bottom-xs-0").Find("[class = \"margin-top-xs-0\"]").Each(func(i int, s *goquery.Selection) {
		x := s.Text()
		prophecy = strings.TrimSpace(x)
	})
	return prophecy
}
