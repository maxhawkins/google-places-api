package places

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestDetailsCallDo(t *testing.T) {
	// open a test server and immediatly close it
	// we do this to get a valid test URL later on in the communcation fault test
	cts := httptest.NewServer(http.HandlerFunc(handler))
	cts.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	for _, test := range []struct {
		Name      string
		PlaceID   string
		Language  string
		Extension string
		URL       string
		Want      error
	}{
		{
			Name:      "OK Response",
			PlaceID:   "ChIJLU7jZClu5kcR4PcOOO6p3I0",
			Language:  "en",
			Extension: "review_summary",
			URL:       ts.URL,
			Want:      nil,
		},
		{
			Name:    "Invalid Request",
			PlaceID: "invalid_request",
			URL:     ts.URL,
			Want:    errors.New("INVALID_REQUEST"),
		},
		{
			Name:    "Non-OK Status",
			PlaceID: "notok",
			URL:     ts.URL,
			Want:    errors.New("bad resp 400: "),
		},
		{
			Name:    "Invalid JSON",
			PlaceID: "invalid_json",
			URL:     ts.URL,
			Want:    errors.New("json: cannot unmarshal string into Go value of type places.DetailsResponse"),
		},
		{
			Name:    "Communication Problem",
			PlaceID: "wrong",
			URL:     cts.URL,
			Want: errors.New(
				"Get " + cts.URL + "/details/json?key=testkey&placeid=wrong: dial tcp " + cts.Listener.Addr().String() + ": getsockopt: connection refused",
			),
		},
	} {
		var service = Service{
			client: http.DefaultClient,
			key:    "testkey",
			url:    test.URL,
		}

		var details = DetailsCall{
			service:    &service,
			placeID:    test.PlaceID,
			Language:   test.Language,
			Extensions: test.Extension,
		}
		_, got := details.Do()

		if got != test.Want {
			if got.Error() != test.Want.Error() {
				t.Errorf("DetailsCall{}.Do() %v = %#v, want %#v",
					test.Name, got, test.Want)
			}
		}
	}
}

func handler(writer http.ResponseWriter, reader *http.Request) {
	uri := reader.URL.RequestURI()

	if uri == "/details/json?extensions=review_summary&key=testkey&language=en&placeid=ChIJLU7jZClu5kcR4PcOOO6p3I0" {
		fmt.Fprint(writer, readResponse("ok"))
		return
	}

	if uri == "/details/json?key=testkey&placeid=invalid_request" {
		fmt.Fprint(writer, readResponse("invalid_request"))
		return
	}

	if uri == "/details/json?key=testkey&placeid=invalid_json" {
		fmt.Fprint(writer, readResponse("invalid_json"))
		return
	}

	if uri == "/details/json?key=testkey&placeid=notok" {
		writer.WriteHeader(400)
		fmt.Fprint(writer, "")
		return
	}
}

func readResponse(responseType string) string {
	absPath, err := filepath.Abs("../data/" + responseType + ".json")
	if err != nil {
		panic(err)
	}

	response, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	return string(response)
}
