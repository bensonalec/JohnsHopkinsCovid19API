package main


import (
	"net/http"
	"io/ioutil"
	"github.com/gocolly/colly"
)

func getNextPage() string{
	c := colly.NewCollector()
	var links []string
	c.OnHTML(".js-navigation-open", func(e *colly.HTMLElement) {
		// fmt.Println(e.Attr("href"))
		links = append(links,e.Attr("href"))
	})
	c.Visit("https://github.com/CSSEGISandData/COVID-19/tree/master/csse_covid_19_data/csse_covid_19_daily_reports")
	//get last, then
	g := colly.NewCollector()
	rawLink := ""
	g.OnHTML("#raw-url", func(e *colly.HTMLElement) {
		rawLink = e.Attr("href")
	})
	g.Visit("https://github.com" + links[len(links)-2])
	return "https://github.com" + rawLink
}

func getPage() string {
	
	url := getNextPage()
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	return string(html)
}

