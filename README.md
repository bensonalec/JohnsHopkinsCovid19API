# Johns Hopkins Covid 19 API

Using the data from the Repository located at: https://github.com/CSSEGISandData/COVID-19  
All stats returned are as up to date as the CSV in the repository is, on each request it returns the freshly updated response 
## Goal
The goal of this API is to provide quick and useful, up to date statistics on Covid-19 deaths, confirmed cases, and recovered cases by location.  

## Deployment
To deploy this application, simply install golang, then type ``` go get github.com/gorilla/mux ``` and "go get github.com/gocolly/colly". Then build the application with "go build main.go backend.go backendDaily.go backendTimeSeries.go handlers.go" in the terminal after cloning this repo. Then, type "sudo ./main" and the API will be up and listening on port 80 for get requests. 

## TODO  
1. Add endpoints for different days data  

## Endpoints

### /api/Daily
Retrieve all of the nodes for the day  
**EXAMPLE**  
/api/Daily
```
{
    "Nodes": [
        {
            "Province/State": "Hubei",
            "Country/Region": "China",
            "Last_Update": "2020-03-18T12:13:09",
            "Confirmed": 67800,
            "Deaths": 3122,
            "Recovered": 56927,
            "Latitude": 30.9756,
            "Longitude": 112.2707
        },
		...
        {
            "Province/State": "",
            "Country/Region": "The Gambia",
            "Last_Update": "2020-03-18T14:13:56",
            "Confirmed": 0,
            "Deaths": 0,
            "Recovered": 0,
            "Latitude": 13.4667,
            "Longitude": -16.6000
        }
    ]
}
```

### /api/Daily/State/{state}
Based on the state/province that is passed in, will return and reported nodes from that state/province  
**EXAMPLE**  
/api/Daily/State/Washington  
```
{
    "Nodes": [
        {
            "Province/State": "Washington",
            "Country/Region": "US",
            "Last_Update": "2020-03-16T23:53:03",
            "Confirmed": 904,
            "Deaths": 48,
            "Recovered": 1,
            "Latitude": 47.4009,
            "Longitude": -121.4905
        }
    ]
}
```
### /api/Daily/Country/{country}
Based on the country passed in, will return all reported nodes from that country  
**EXAMPLE**  
/api/Daily/Country/US  
```
{
    "Nodes": [
        {
            "Province/State": "New York",
            "Country/Region": "US",
            "Last_Update": "2020-03-16T21:53:03",
            "Confirmed": 967,
            "Deaths": 10,
            "Recovered": 0,
            "Latitude": 42.1657,
            "Longitude": -74.9481
        },
        ...
        {
            "Province/State": "West Virginia",
            "Country/Region": "US",
            "Last_Update": "2020-03-10T02:33:04",
            "Confirmed": 0,
            "Deaths": 0,
            "Recovered": 0,
            "Latitude": 38.4912,
            "Longitude": -80.9545
        }
    ]
}
```  

### /api/Daily/CountryAndState/{state}/{country}

**EXAMPLE**  
/api/Daily/CountryAndState/Washington/US  
```
{
    "Nodes": [
        {
            "Province/State": "Washington",
            "Country/Region": "US",
            "Last_Update": "2020-03-16T23:53:03",
            "Confirmed": 904,
            "Deaths": 48,
            "Recovered": 1,
            "Latitude": 47.4009,
            "Longitude": -121.4905
        }
    ]
}
```
### /api/Timeseries/Confirmed
Get the timeseries for all locations of confirmed cases.  
**EXAMPLE**  
/api/Timeseries/Confirmed  
```
{
    "Nodes": [
        {
            "Province/State": "",
            "Country/Region": "Thailand",
            "Latitude": 5,
            "Longitude": 7,
            "Days": [
                {
                    "1/22/20": 2
                },
				...
                {
                    "3/18/20": 212
                },
                {
                    "3/19/20": 272
                }
            ]
        },
		...
       {
            "Province/State": "",
            "Country/Region": "The Gambia",
            "Latitude": 0,
            "Longitude": 0,
            "Days": [
                {
                    "1/22/20": 0
                },
				...
               {
                    "3/19/20": 0
                }
            ]
        }
    ]
}
```
### /api/Timeseries/Deaths
Get the timeseries for all locations of deaths.  
**EXAMPLE**  
/api/Timeseries/Deaths  
```
{
    "Nodes": [
        {
            "Province/State": "",
            "Country/Region": "Thailand",
            "Latitude": 0,
            "Longitude": 0,
            "Days": [
                {
                    "1/22/20": 0
                },
				...
                {
                    "3/19/20": 1
                }
            ]
        },
		...
        {
            "Province/State": "",
            "Country/Region": "The Gambia",
            "Latitude": 0,
            "Longitude": 0,
            "Days": [
                {
                    "1/22/20": 0
                },
				...
                {
                    "3/19/20": 0
                }
            ]
        }
    ]
}
```
### /api/Timeseries/Recovered
Get the timeseries for all locations of recovered cases.  
**EXAMPLE**  
/api/Timeseries/Recovered  
```
{
    "Nodes": [
        {
            "Province/State": "",
            "Country/Region": "Thailand",
            "Latitude": 0,
            "Longitude": 0,
            "Days": [
                {
                    "1/22/20": 0
                },
				...
                {
                    "3/19/20": 1
                }
            ]
        },
		...
        {
            "Province/State": "",
            "Country/Region": "The Gambia",
            "Latitude": 0,
            "Longitude": 0,
            "Days": [
                {
                    "1/22/20": 0
                },
				...
                {
                    "3/19/20": 0
                }
            ]
        }
    ]
}
```

### /api/Timeseries/Confirmed/State/{state}
Get the timeseries of confirmed cases for all locations that have a matching state/province.
### /api/Timeseries/Recovered/State/{state}
Get the timeseries of recovered cases for all locations that have a matching state/province.
### /api/Timeseries/Deaths/State/{state}
Get the timeseries of confirmed cases for all locations that have a matching state/province.
### /api/Timeseries/Confirmed/Country/{country}
Get the timeseries of confirmed cases for all locations that have a matching country.
### /api/Timeseries/Recovered/Country/{country}
Get the timeseries of recovered cases for all locations that have a matching country.
### /api/Timeseries/Deaths/Country/{country}
Get the timeseries of deaths cases for all locations that have a matching country.
