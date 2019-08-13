package cep

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"

	"fmt"
)

type Address struct {
	Street       string
	Neighborhood string
	City         string
	Zipcode      string
}

func GetAddressByZipcode(zipcode string) (address Address) {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
	})
	c.OnHTML(".ctrlcontent", func(e *colly.HTMLElement) {

		e.ForEach(".tmptabela tbody tr", func(_ int, e *colly.HTMLElement) {
			var productName string

			goquerySelection := e.DOM
			if goquerySelection.Find("td").Last().Text() == zipcode {
				goquerySelection.Find("td").Siblings().Each(func(i int, s *goquery.Selection) {
					fmt.Printf("%d, Sibling text: %s\n", i, s.Text())
					switch i {
					case 0:
						address.Neighborhood = s.Text()
					case 1:
						address.City = s.Text()
					case 2:
						address.Zipcode = s.Text()
					case 3:
						address.Street = RemoveRange(s.Text())
					}
				})
				productName = e.ChildText("td")

				if productName == "" {
					// If we can't get any name, we return and go directly to the next element
					return
				}

				fmt.Printf("Product Name: %s \n", address)
			}
		})
	})
	c.Post("http://www.buscacep.correios.com.br/sistemas/buscacep/resultadoBuscaCepEndereco.cfm", map[string]string{"relaxation": zipcode, "tipoCEP": "ALL"})
	return
}

func RemoveRange(street string) string {
	return strings.Split(street, " - ")[0]
}
