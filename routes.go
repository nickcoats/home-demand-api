package main

import (
    "net/http"

    "toddler-travel-api/handlers/api"
    "toddler-travel-api/handlers/person"
    "toddler-travel-api/handlers/country"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        authorize(api.Index,basicAuth),
    },
    Route{
        "Traveler/",
        "GET",
        "/traveler",
        authorize(person.Get, basicAuth),
    },
    Route{
        "Country",
        "GET",
        "/country",
        authorize(country.GetAll, basicAuth),
    },
    Route{
        "Country/ID",
        "GET",
        "/country/{id}",
        authorize(country.Get, basicAuth),
    },


}
