package main


import (
    "net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/Daily/State/{state}",stateGET).Methods(http.MethodGet)
	api.HandleFunc("/Daily/Country/{country}",countryGET).Methods(http.MethodGet)
	api.HandleFunc("/Daily/CountryAndState/{state}/{country}",countrystateGET).Methods(http.MethodGet)
	api.HandleFunc("/Daily",dailyallGET).Methods(http.MethodGet)

	api.HandleFunc("/Timeseries/Confirmed",TSallConfirmedGET).Methods(http.MethodGet)
	api.HandleFunc("/Timeseries/Recovered",TSallRecoveredGET).Methods(http.MethodGet)
	api.HandleFunc("/Timeseries/Deaths",TSallDeathsGET).Methods(http.MethodGet)

	// api.HandleFunc("/Timeseries/Confirmed/State/{state}",TSallDeathsGET).Methods(http.MethodGet)
	// api.HandleFunc("/Timeseries/Recovered/State/{state}",TSallDeathsGET).Methods(http.MethodGet)
	// api.HandleFunc("/Timeseries/Deaths/State/{state}",TSallDeathsGET).Methods(http.MethodGet)

	// api.HandleFunc("/Timeseries/Confirmed/Country/{country}",TSallConfirmedGET).Methods(http.MethodGet)
	// api.HandleFunc("/Timeseries/Recovered/Country/{country}",TSallConfirmedGET).Methods(http.MethodGet)
	// api.HandleFunc("/Timeseries/Deaths/Country/{country}",TSallConfirmedGET).Methods(http.MethodGet)

	api.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":80", r))

}

