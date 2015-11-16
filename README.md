# Google Places API Client for Go

[![Build Status](https://travis-ci.org/maxhawkins/google-places-api.png)](https://travis-ci.org/maxhawkins/google-places-api)
[![GoDoc](https://godoc.org/github.com/maxhawkins/google-places-api/places?status.svg)](http://godoc.org/github.com/maxhawkins/google-places-api/places)

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
    "time"

    "github.com/maxhawkins/google-places-api/places"
)

func main() {
    service := places.NewService(http.DefaultClient, "<your api key>")

    call := service.Nearby(37.7833, -122.4167) // San Francisco
    call.Types = append(call.Types, places.Cafe)
    call.Radius = 500

    resp, err := call.Do()
    if places.IsZeroResults(err) {
        fmt.Println("no results")
        return
    }
    if err != nil {
        panic(err)
    }

    results := resp.Results
    token := resp.NextPageToken

    for token != "" {
        time.Sleep(2 * time.Second) // Rate limit

        call.PageToken = token
        resp, err := call.Do()
        if err != nil {
            panic(err)
        }

        token = resp.NextPageToken
        results = append(results, resp.Results...)
    }

    for _, result := range results {
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
* Text Search

#### What's not done

* Place Add
* Place Delete
* Place Photos
* Place Autocomplete
* Query Autocomplete
* More Examples
