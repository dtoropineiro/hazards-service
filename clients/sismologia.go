package clients

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const sismologiaBaseUrl = "https://www.sismologia.cl/sismicidad/catalogo"
const HTML = ".html"

const SLASH = "/"

type EarthquakeSismologia struct {
	DateLocal string
	Location  string
	DateUTC   string
	LatLong   string
	Depth     string
	Magnitude string
}

func GetEarthquakes() ([]EarthquakeSismologia, error) {
	earthquakes := make([]EarthquakeSismologia, 0)
	c := colly.NewCollector()
	url := buildFullUrl()
	print(url)
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put(buildFullUrl(), r.URL.String())
	})
	const tableSelector = "table.sismologia.detalle"
	c.OnHTML(tableSelector, func(e *colly.HTMLElement) {
		logrus.Info(e.Name)
		e.ForEach("tr", func(i int, row *colly.HTMLElement) {

			columns := row.DOM.Find("td")
			if columns.First().Text() != "" {

				earthquake := EarthquakeSismologia{
					DateLocal: columns.Eq(1).Text(),
					Location:  columns.Eq(0).Text(),
					DateUTC:   columns.Eq(3).Text(),
					LatLong:   columns.Eq(2).Text(),
					Depth:     columns.Eq(3).Text(),
					Magnitude: columns.Eq(4).Text(),
				}
				earthquakes = append(earthquakes, earthquake)
			}

		})
	})
	c.OnResponse(func(r *colly.Response) {

		fmt.Println("-----------------------------")

		fmt.Println(r.StatusCode)

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(earthquakes, "", "    ")
		if err != nil {
			logrus.Error(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("products.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}

	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return earthquakes, nil
}

func buildFullUrl() string {
	chileZone, _ := time.LoadLocation("America/Santiago")
	nowChile := time.Now().In(chileZone)

	currentYear := fmt.Sprintf("%d", nowChile.Year())
	currentMonth := fmt.Sprintf("%02d", nowChile.Month())
	currentDay := fmt.Sprintf("%02d", nowChile.Day())
	mergedCurrentDay := currentYear + currentMonth + currentDay

	return sismologiaBaseUrl +
		SLASH +
		currentYear +
		SLASH +
		currentMonth +
		SLASH +
		mergedCurrentDay +
		HTML
}
