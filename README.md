# Johns Hopkins Covid 19 API

Using the data from the Repository located at: https://github.com/CSSEGISandData/COVID-19  
All stats returned are as up to date as the CSV in the repository is, on each request it returns the freshly updated response 
## Goal
The goal of this API is to provide quick and useful, up to date statistics on Covid-19 deaths, confirmed cases, and recovered cases by location.  

## TODO  
1. Properly parse each line, right now does naive parsing using ,. Since it's a csv, this is not gonna work for things like Korea, South  (This is implemented in time series backend)
2. Add endpoints for different days data  
3. Add endpoints for Time Series State and Country
4. Add documentation for Time Series Endpoints

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
