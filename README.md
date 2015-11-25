# Google Places API Client for Go

[![Build Status](https://travis-ci.org/maxhawkins/google-places-api.png)](https://travis-ci.org/maxhawkins/google-places-api)
[![GoDoc](https://godoc.org/github.com/maxhawkins/google-places-api/places?status.svg)](http://godoc.org/github.com/maxhawkins/google-places-api/places)

A Go client for the [Google Places API](https://developers.google.com/places/webservice/). A work in progress, contributions welcome.

To install this package, run:

```
go get gopkg.in/maxhawkins/google-places-api.v1/places
```

## Example

``` go
package main

import (
    "fmt"
    "net/http"
    "time"

    "gopkg.in/maxhawkins/google-places-api.v1/places"
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

## Gotchas

* This package is a work in progress. Not all API calls are implemented.
* [There is a short delay](https://developers.google.com/places/web-service/search#PlaceSearchPaging) between when a NextPageToken is issued, and when it will become valid. Requesting the next page before it is available will return an `INVALID_REQUEST` response. Retrying the request with the same NextPageToken will return the next page of results.

#### What's done

* Nearby Search
* Place Details
* Radar Search
* Text Search

#### What's not done

* Place Add
* Place Autocomplete
* Place Delete
* Place Photos
* Query Autocomplete
* More Examples
