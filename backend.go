package main


import (
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/gocolly/colly"
	
	"encoding/csv"
	"io"
	"log"

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

func getCountry(country string) string{
	
	html := getPage()
	
	spl := strings.Split(string(html),"\n")
	finalFinal := "{\n\"Nodes\":["
	if(len(spl) > 0) {
		var foundList []string
		for _,ele := range spl {
			line := strings.Split(ele,",")
			if(len(line) > 1 && strings.ToLower(line[1]) == strings.ToLower(country)){
				// fmt.Println(line)
				//now build JSON response
				/*
				{
					Province/State:"",
					Country/Region:"",
					Last Update: "",
					Confirmed:"",
					Deaths:"",
					Recovered:"",
					Latitude:"",
					Longitude:""
				}
				*/
				finalReturn := "{" + "\n\"Province/State\":\"" + line[0] + "\",\n\"Country/Region\":\""+line[1]+"\",\n\"Last_Update\":\""+line[2]+"\",\n\"Confirmed\":"+line[3]+",\n\"Deaths\":"+line[4]+",\n\"Recovered\":"+line[5]+",\n\"Latitude\":"+line[6]+",\n\"Longitude\":"+line[7]+"\n},"
				finalFinal += finalReturn
				foundList = append(foundList,finalReturn)
				// fmt.Println(finalReturn)
			}
			
			
			
		}
		if(len(foundList) > 0) {
			finalFinal = finalFinal[:len(finalFinal)-1]
		}
		
		
	}
	finalFinal += "]\n}"
	
	return finalFinal
}

func getState(city string) string{
	html := getPage()
	
	spl := strings.Split(string(html),"\n")
	finalFinal := "{\n\"Nodes\":["
	if(len(spl) > 0) {
		var foundList []string
		for _,ele := range spl {
			line := strings.Split(ele,",")

			if(len(line) > 1 && strings.ToLower(line[0]) == strings.ToLower(city)){
				// fmt.Println(line)
				//now build JSON response
				/*
				{
					Province/State:"",
					Country/Region:"",
					Last Update: "",
					Confirmed:"",
					Deaths:"",
					Recovered:"",
					Latitude:"",
					Longitude:""
				}
				*/
				finalReturn := "{" + "\n\"Province/State\":\"" + line[0] + "\",\n\"Country/Region\":\""+line[1]+"\",\n\"Last_Update\":\""+line[2]+"\",\n\"Confirmed\":"+line[3]+",\n\"Deaths\":"+line[4]+",\n\"Recovered\":"+line[5]+",\n\"Latitude\":"+line[6]+",\n\"Longitude\":"+line[7]+"\n},"
				finalFinal += finalReturn
				foundList = append(foundList,finalReturn)
				// fmt.Println(finalReturn)
			}
			
			
			
		}
		if(len(foundList) > 0) {
			finalFinal = finalFinal[:len(finalFinal)-1]
		}
		
		
	}
	finalFinal += "]\n}"
	
	return finalFinal
}

func getCountryState(state string,country string) string{
	html := getPage()
	
	spl := strings.Split(string(html),"\n")
	finalFinal := "{\n\"Nodes\":["
	if(len(spl) > 0) {
		var foundList []string
		for _,ele := range spl {
			line := strings.Split(ele,",")
			if(len(line) > 1 && strings.ToLower(line[0]) == strings.ToLower(state) && strings.ToLower(line[1]) == strings.ToLower(country)){
				// fmt.Println(line)
				//now build JSON response
				/*
				{
					Province/State:"",
					Country/Region:"",
					Last Update: "",
					Confirmed:"",
					Deaths:"",
					Recovered:"",
					Latitude:"",
					Longitude:""
				}
				*/
				finalReturn := "{" + "\n\"Province/State\":\"" + line[0] + "\",\n\"Country/Region\":\""+line[1]+"\",\n\"Last_Update\":\""+line[2]+"\",\n\"Confirmed\":"+line[3]+",\n\"Deaths\":"+line[4]+",\n\"Recovered\":"+line[5]+",\n\"Latitude\":"+line[6]+",\n\"Longitude\":"+line[7]+"\n},"
				finalFinal += finalReturn
				foundList = append(foundList,finalReturn)
				// fmt.Println(finalReturn)
			}
			
			
			
		}
		if(len(foundList) > 0) {
			finalFinal = finalFinal[:len(finalFinal)-1]
		}
		
		
	}
	finalFinal += "]\n}"
	
	return finalFinal
}



func getAll() string{
	html := getPage()
	
	spl := strings.Split(string(html),"\n")
	finalFinal := "{\n\"Nodes\":["
	if(len(spl) > 0) {
		var foundList []string
		for _,ele := range spl {
			line := strings.Split(ele,",")
			if(len(line) == 8) {

				//now build JSON response
				/*
				{
					Province/State:"",
					Country/Region:"",
					Last Update: "",
					Confirmed:"",
					Deaths:"",
					Recovered:"",
					Latitude:"",
					Longitude:""
				}
				*/
				finalReturn := "{" + "\n\"Province/State\":\"" + line[0] + "\",\n\"Country/Region\":\""+line[1]+"\",\n\"Last_Update\":\""+line[2]+"\",\n\"Confirmed\":"+line[3]+",\n\"Deaths\":"+line[4]+",\n\"Recovered\":"+line[5]+",\n\"Latitude\":"+line[6]+",\n\"Longitude\":"+line[7]+"\n},"
				finalFinal += finalReturn
				foundList = append(foundList,finalReturn)

			}			
		}
		if(len(foundList) > 0) {
			finalFinal = finalFinal[:len(finalFinal)-1]
		}
		
		
	}
	finalFinal += "]\n}"
	
	return finalFinal
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

func getTSConfirmedAll() string{
	resp, err := http.Get("https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Confirmed.csv")
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
	return getTSData(string(html))
}

func getTSDDeathsAll() string{
	resp, err := http.Get("https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Deaths.csv")
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
	return getTSData(string(html))

}

func getTSDRecoveredAll() string{
	resp, err := http.Get("https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-Recovered.csv")
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
	return getTSData(string(html))

}


func getTSData(in string) string{
	r := csv.NewReader(strings.NewReader(in))
	finalFinal := "{\n\"Nodes\":["
	var found []string
	var first []string
	i := 0
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if(i == 0) {
			first = line
			first = first[4:]
		} else {
			finalReturn := "{" + "\n\"Province/State\":\"" + line[0] + "\",\n\"Country/Region\":\""+line[1]+"\",\n\"Latitude\":"+line[6]+",\n\"Longitude\":"+line[7]+",\n\"Days\":["
			for ind,ele := range first {
				finalReturn += "{\"" + ele + "\":" + line[ind+4]+"},\n"
			}
			finalReturn = finalReturn[:len(finalReturn)-2]
			finalReturn += "]\n},"
			finalFinal += finalReturn
			found = append(found,finalReturn)

		}
		i++
	}
	if(len(found) > 0) {
		finalFinal = finalFinal[:len(finalFinal)-1]
	}
	finalFinal += "]\n}"
	return finalFinal
	
}