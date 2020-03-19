package main


import (
    "net/http"
	"github.com/gorilla/mux"
	"log"
	
	
)

func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"status code": "400"}`))
}

func stateGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var state string
	var toWrite string

	if val,ok := pathParams["state"]; ok {
		state = val

		toWrite = getState(state)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))

	}


}

func countryGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var country string
	var toWrite string

	if val,ok := pathParams["country"]; ok {
		country = val
		toWrite = getCountry(country)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))	
	}


}

func countrystateGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var country string
	var state string
	var toWrite string

	if val,ok := pathParams["state"]; ok {
		state = val
		if val,ok := pathParams["country"]; ok {
			country = val
			toWrite = getCountryState(state,country)
			w.WriteHeader(200)
			w.Write([]byte((toWrite)))	
	
		}
	}


}

func dailyallGET(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var toWrite string

	toWrite = getAll()
	w.WriteHeader(200)
	w.Write([]byte((toWrite)))	
	


}


func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/Daily/State/{state}",stateGET).Methods(http.MethodGet)
	api.HandleFunc("/Daily/Country/{country}",countryGET).Methods(http.MethodGet)
	api.HandleFunc("/Daily/CountryAndState/{state}/{country}",countrystateGET).Methods(http.MethodGet)
	api.HandleFunc("/Daily",dailyallGET).Methods(http.MethodGet)
	// api.HandleFunc("/Timeseries/Confirmed",stateGET).Methods(http.MethodGet)
	// api.HandleFunc("/Timeseries/Recovered",stateGET).Methods(http.MethodGet)
	// api.HandleFunc("/Timeseries/Deaths",stateGET).Methods(http.MethodGet)
	api.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":80", r))

}

