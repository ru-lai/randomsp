package randomsp

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getDaxStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/DAX")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tables := doc.Find(".mw-parser-output > table")
	tables.Filter("#constituents").Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(4)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getFinancialTimesStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/FTSE_100_Index")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("#constituents > tbody")
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(2)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getItalianFinancialTimesStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/FTSE_MIB")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("#constituents > tbody")
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(2)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getNasdaqStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/NASDAQ-100")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("#constituents > tbody")
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:nth-child(2)")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getStandardPoorsStocks() (stocks []string, err error) {
	res, err := http.Get("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	tbody := doc.Find("tbody").First()
	tbody.Find("tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td:first-child")
		if td.Text() != "" {
			stocks = append(stocks, td.Text())
		}
	})
	return
}

func getRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getRandomString(ss []string) string {
	randInt := getRandomInt(1, len(ss))
	return ss[randInt:(randInt + 1)][0]
}

func GetRandomDaxStock() (stock Stock, err error) {
	rand.Seed(time.Now().UnixNano())
	stockSlice, err := getDaxStocks()
	if err != nil {
		return
	}

	stock = Stock{getRandomString(stockSlice), "Dax"}
	return
}

func GetRandomFinancialTimesStock() (stock Stock, err error) {
	rand.Seed(time.Now().UnixNano())
	stockSlice, err := getFinancialTimesStocks()
	if err != nil {
		return
	}

	stock = Stock{getRandomString(stockSlice), "Financial Times"}
	return
}

func GetRandomItalianFinancialTimesStock() (stock Stock, err error) {
	rand.Seed(time.Now().UnixNano())
	stockSlice, err := getItalianFinancialTimesStocks()
	if err != nil {
		return
	}

	stock = Stock{getRandomString(stockSlice), "Italian Financial Times"}
	return
}

func GetRandomNasdaqStock() (stock Stock, err error) {
	rand.Seed(time.Now().UnixNano())
	stockSlice, err := getNasdaqStocks()
	if err != nil {
		return
	}

	stock = Stock{getRandomString(stockSlice), "Nasdaq"}

	return
}

func GetRandomSPStock() (stock Stock, err error) {
	rand.Seed(time.Now().UnixNano())
	stockSlice, err := getStandardPoorsStocks()
	if err != nil {
		return
	}

	stock = Stock{getRandomString(stockSlice), "S&P 500"}

	return
}

func GetRandomIndexStock() (stock Stock, err error) {
	rand.Seed(time.Now().UnixNano())
	stockFuncs := []func() (Stock, error){GetRandomNasdaqStock, GetRandomSPStock, GetRandomFinancialTimesStock, GetRandomItalianFinancialTimesStock, GetRandomDaxStock}
	stock, err = stockFuncs[getRandomInt(0, len(stockFuncs))]()
	return
}
