package main


import (
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/csv"
	"io"
	"log"

)


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

func getTSConfirmedState(state string) string{
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
	return getTSDataState(string(html),state)
}

func getTSConfirmedCountry(country string) string{
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
	return getTSDataCountry(string(html),country)
}


func getTSDeathsAll() string{
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

func getTSDeathsState(state string) string{
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
	return getTSDataState(string(html),state)

}

func getTSDeathsCountry(country string) string{
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
	return getTSDataCountry(string(html),country)

}


func getTSRecoveredAll() string{
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

func getTSRecoveredState(state string) string{
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
	return getTSDataState(string(html),state)

}

func getTSRecoveredCountry(country string) string{
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
	return getTSDataCountry(string(html),country)

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

func getTSDataState(in string, state string) string{
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
			if(strings.ToLower(line[0]) == strings.ToLower(state)) {
				finalReturn := "{" + "\n\"Province/State\":\"" + line[0] + "\",\n\"Country/Region\":\""+line[1]+"\",\n\"Latitude\":"+line[6]+",\n\"Longitude\":"+line[7]+",\n\"Days\":["
				for ind,ele := range first {
					finalReturn += "{\"" + ele + "\":" + line[ind+4]+"},\n"
				}
				finalReturn = finalReturn[:len(finalReturn)-2]
				finalReturn += "]\n},"
				finalFinal += finalReturn
				found = append(found,finalReturn)
	
			}

		}
		i++
	}
	if(len(found) > 0) {
		finalFinal = finalFinal[:len(finalFinal)-1]
	}
	finalFinal += "]\n}"
	return finalFinal
	
}

func getTSDataCountry(in string, country string) string{
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
			if(strings.ToLower(line[1]) == strings.ToLower(country)) {
				finalReturn := "{" + "\n\"Province/State\":\"" + line[0] + "\",\n\"Country/Region\":\""+line[1]+"\",\n\"Latitude\":"+line[6]+",\n\"Longitude\":"+line[7]+",\n\"Days\":["
				for ind,ele := range first {
					finalReturn += "{\"" + ele + "\":" + line[ind+4]+"},\n"
				}
				finalReturn = finalReturn[:len(finalReturn)-2]
				finalReturn += "]\n},"
				finalFinal += finalReturn
				found = append(found,finalReturn)
	
			}

		}
		i++
	}
	if(len(found) > 0) {
		finalFinal = finalFinal[:len(finalFinal)-1]
	}
	finalFinal += "]\n}"
	return finalFinal
	
}