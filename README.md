# Google Places API Client for Go

[![Build Status](https://travis-ci.org/maxhawkins/google-places-api.png)](https://travis-ci.org/maxhawkins/google-places-api)
[![GoDoc](https://godoc.org/github.com/maxhawkins/google-places-api?status.svg)](http://godoc.org/github.com/maxhawkins/google-places-api)

A Go client for the [Google Places API](https://developers.google.com/places/webservice/). A work in progress, contributions welcome.

To install this package, run:

```
go get github.com/maxhawkins/google-places-api/places
```

## Example

``` go
package main

import (
    "fmt"
    "net/http"

    "github.com/maxhawkins/google-places-api/places"
)

func main() {
    service := places.NewService(http.DefaultClient, "<your api key>")

    call := service.Nearby(37.7833, -122.4167) // San Francisco
    call.Types = append(call.Types, places.Cafe)
    call.Radius = 7000

    resp, err := call.Do()
    if places.IsZeroResults(err) {
        fmt.Println("no results")
        return
    }
    if err != nil {
        panic(err)
    }

    for _, result := range resp.Results {
        fmt.Println(result.Name)
    }
}
```

## Status

This package is a work in progress. Not all API calls are implemented.

#### What's done

* Place Details
* Nearby Search
* Radar Search

#### What's not done

* Text Search
* Place Add
* Place Delete
* Place Photos
* Place Autocomplete
* Query Autocomplete
* More Examples
