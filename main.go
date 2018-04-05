package randomsp

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func GetSPStocks() (stocks []string) {
	res, err := http.Get("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	tbody := doc.Find("tbody").First()
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:first-child")
		stocks = append(stocks, td.Text())
	})
	return
}

func GetRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func GetRandString(ss []string) string {
	randInt := GetRandomInt(1, len(ss))
	return ss[randInt:(randInt + 1)][0]
}

func GetRandomSPStock() string {
	rand.Seed(time.Now().UnixNano())
	stockSlice := GetSPStocks()
	return GetRandString(stockSlice)
}
