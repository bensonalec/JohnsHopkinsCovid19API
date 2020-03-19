package main


import (
    "net/http"
	"github.com/gorilla/mux"
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

func TSallConfirmedGET(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var toWrite string

	toWrite = getTSConfirmedAll()
	w.WriteHeader(200)
	w.Write([]byte((toWrite)))	
	
}

func TSallRecoveredGET(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var toWrite string

	toWrite = getTSRecoveredAll()
	w.WriteHeader(200)
	w.Write([]byte((toWrite)))	
	
}

func TSallDeathsGET(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var toWrite string

	toWrite = getTSDeathsAll()
	w.WriteHeader(200)
	w.Write([]byte((toWrite)))	
	
}

func TSConfirmedstateGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var state string
	var toWrite string

	if val,ok := pathParams["state"]; ok {
		state = val

		toWrite = getTSConfirmedState(state)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))

	}


}

func TSRecoveredstateGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var state string
	var toWrite string

	if val,ok := pathParams["state"]; ok {
		state = val

		toWrite = getTSRecoveredState(state)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))

	}


}

func TSDeathsstateGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var state string
	var toWrite string

	if val,ok := pathParams["state"]; ok {
		state = val

		toWrite = getTSDeathsState(state)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))

	}


}

func TSConfirmedcountryGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var country string
	var toWrite string

	if val,ok := pathParams["country"]; ok {
		country = val

		toWrite = getTSConfirmedCountry(country)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))

	}


}

func TSRecoveredcountryGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var country string
	var toWrite string

	if val,ok := pathParams["country"]; ok {
		country = val

		toWrite = getTSRecoveredCountry(country)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))

	}


}

func TSDeathscountryGET(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var country string
	var toWrite string

	if val,ok := pathParams["country"]; ok {
		country = val

		toWrite = getTSDeathsCountry(country)
		w.WriteHeader(200)
		w.Write([]byte((toWrite)))

	}


}
