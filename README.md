# Johns Hopkins Covid 19 API

Using the data from the Repository located at: https://github.com/CSSEGISandData/COVID-19  
All stats returned are as up to date as the CSV in the repository is, on each request it returns the freshly updated response  
Right now only gets daily data
## Goal
The goal of this API is to provide quick and useful, up to date statistics on Covid-19 deaths, confirmed cases, and recovered cases by location.  

## TODO
Not sure, potentially add more endpoints, at this time I'm not sure.  
1. Properly parse each line, right now does naive parsing using ,. Since it's a csv, this is not gonna work for things like Korea, South  
2. Add endpoints for different days data  
3. Add endpoints for the Time Series

## Endpoints

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
