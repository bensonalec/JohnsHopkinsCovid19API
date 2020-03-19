package main


import (
	"net/http"
	"io/ioutil"
	"strings"
)

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
	url := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/03-16-2020.csv"
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