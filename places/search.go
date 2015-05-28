package places

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	errInvalidByProminence = errors.New("radius must be specified when RankByProminence is used")
	errInvalidByDistance   = errors.New("when RankByDistance is specified, one or more of keyword, name, or types is required")
)

// Nearby lets you search for places within a specified area. You can refine your search request by supplying keywords or specifying the type of place you are searching for.
func (p *Service) Nearby(lat, lng float64) *NearbyCall {
	return &NearbyCall{
		service: p,
		lat:     lat,
		lng:     lng,
	}
}

type NearbyCall struct {
	service *Service

	// The latitude/longitude around which to retrieve place information
	lat, lng float64

	// A term to be matched against all content that Google has indexed for this place, including but not limited to name, type, and address, as well as customer reviews and other third-party content.
	Keyword string
	// The language code, indicating in which language the results should be returned, if possible.
	Language string
	// Restricts results to only those places within the specified price level.
	MinPrice, MaxPrice *PriceLevel
	// One or more terms to be matched against the names of places, separated with a space character. Results will be restricted to those containing the passed name values. Note that a place may have additional names associated with it, beyond its listed name. The API will try to match the passed name value against all of these names. As a result, places may be returned in the results whose listed names do not match the search term, but whose associated names do.
	Name string
	// Returns only those places that are open for business at the time the query is sent. Places that do not specify opening hours in the Google Places database will not be returned if you include this parameter in your query.
	OpenNow bool
	// Defines the distance (in meters) within which to return place results. The maximum allowed radius is 50 000 meters. Note that radius must not be included if rankby=distance is specified.
	Radius float64
	// Specifies the order in which results are listed
	RankBy RankBy
	// Restricts the results to places matching at least one of the specified types.
	Types []FeatureType
	// Restricts the search to locations that are Zagat selected businesses.
	ZagatSelected bool
	// Returns the next 20 results from a previously run search. Setting a pagetoken parameter will execute a search with the same parameters used previously — all parameters other than pagetoken will be ignored.
	PageToken string
}

func (n *NearbyCall) validate() error {
	if n.PageToken != "" {
		return nil
	}
	switch n.RankBy {
	case RankByDefault, RankByProminence:
		if n.Radius == 0 {
			return errInvalidByProminence
		}
	case RankByDistance:
		if n.Types == nil && n.Name == "" && n.Keyword == "" {
			return errInvalidByDistance
		}
	}
	return nil
}

func (n *NearbyCall) Do() (*SearchResponse, error) {
	if err := n.validate(); err != nil {
		return nil, err
	}

	searchURL := baseURL + "/nearbysearch/json?" + n.query()

	resp, err := n.service.client.Get(searchURL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad resp %d: %s", resp.StatusCode, body)
	}

	data := &SearchResponse{}
	if err := json.Unmarshal(body, data); err != nil {
		return nil, err
	}

	if data.Status != "OK" {
		return nil, &apiError{
			Status:  data.Status,
			Message: data.ErrorMessage,
		}
	}

	return data, nil
}

func (r *NearbyCall) query() string {
	query := make(url.Values)
	query.Add("key", r.service.key)
	query.Add("location", fmt.Sprintf("%f,%f", r.lat, r.lng))

	if r.PageToken != "" {
		query.Add("pagetoken", r.PageToken)
		return query.Encode()
	}

	if r.Keyword != "" {
		query.Add("keyword", r.Keyword)
	}
	if r.Language != "" {
		query.Add("language", r.Language)
	}
	if r.MinPrice != nil {
		query.Add("minprice", fmt.Sprint(*r.MinPrice))
	}
	if r.MaxPrice != nil {
		query.Add("maxprice", fmt.Sprint(*r.MaxPrice))
	}
	if r.Name != "" {
		query.Add("name", r.Name)
	}
	if r.OpenNow {
		query.Add("opennow", fmt.Sprint(1))
	}
	if r.RankBy != RankByDistance && r.Radius > 0 {
		query.Add("radius", fmt.Sprint(r.Radius))
	}
	if r.RankBy != "" {
		query.Add("rankby", string(r.RankBy))
	}
	if r.ZagatSelected {
		query.Add("zagatselected", "")
	}

	var typeNames []string
	for _, t := range r.Types {
		typeNames = append(typeNames, string(t))
	}
	if len(typeNames) > 0 {
		query.Add("types", strings.Join(typeNames, "|"))
	}

	return query.Encode()
}

// RadarSearch returns results from up to 200 places, but with less detail than is typically returned from a Text Search or Nearby Search request.
func (p *Service) RadarSearch(radius, lat, lng float64) *RadarSearchCall {
	return &RadarSearchCall{
		service: p,
		radius:  radius,
		lat:     lat,
		lng:     lng,
	}
}

type RadarSearchCall struct {
	service *Service

	// The latitude/longitude around which to retrieve place information
	lat, lng float64
	// The distance (in meters) within which to return place results. The maximum allowed radius is 50 000 meters.
	radius float64

	// A term to be matched against all content that Google has indexed for this place, including but not limited to name, type, and address, as well as customer reviews and other third-party content.
	Keyword string
	// Restricts results to only those places within the specified price level.
	MinPrice, MaxPrice *PriceLevel
	// Returns only those places that are open for business at the time the query is sent. Places that do not specify opening hours in the Google Places database will not be returned if you include this parameter in your query.
	OpenNow bool
	// Restricts the results to places matching at least one of the specified types.
	Types []FeatureType
	// Restricts the search to locations that are Zagat selected businesses.
	ZagatSelected bool
	// Returns the next 20 results from a previously run search. Setting a pagetoken parameter will execute a search with the same parameters used previously — all parameters other than pagetoken will be ignored.
	PageToken string
}

func (r *RadarSearchCall) query() string {
	query := make(url.Values)
	query.Add("key", r.service.key)
	if r.Keyword != "" {
		query.Add("keyword", r.Keyword)
	}
	query.Add("location", fmt.Sprintf("%f,%f", r.lat, r.lng))
	if r.MinPrice != nil {
		query.Add("minprice", fmt.Sprint(*r.MinPrice))
	}
	if r.MaxPrice != nil {
		query.Add("maxprice", fmt.Sprint(*r.MaxPrice))
	}
	if r.OpenNow {
		query.Add("opennow", fmt.Sprint(1))
	}
	query.Add("radius", fmt.Sprint(r.radius))

	var typeNames []string
	for _, t := range r.Types {
		typeNames = append(typeNames, string(t))
	}
	if len(typeNames) > 0 {
		query.Add("types", strings.Join(typeNames, "|"))
	}

	if r.ZagatSelected {
		query.Add("zagatselected", "")
	}
	if r.PageToken != "" {
		query.Add("pagetoken", r.PageToken)
	}

	return query.Encode()
}

func (r *RadarSearchCall) Do() (*SearchResponse, error) {
	searchURL := baseURL + "/radarsearch/json?" + r.query()

	resp, err := r.service.client.Get(searchURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad resp %d: %s", resp.StatusCode, body)
	}

	data := &SearchResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Status != "OK" {
		return nil, &apiError{
			Status:  data.Status,
			Message: data.ErrorMessage,
		}
	}

	return data, nil
}

type SearchResponse struct {
	// A list of results matching the query
	Results []PlaceDetails `json:"results"`
	// Contains debugging information to help you track down why the request failed
	Status string `json:"status"`
	// More detailed information about the reasons behind the given status code.
	ErrorMessage string `json:"error_message,omitempty"`
	// A set of attributions about this listing which must be displayed to the user.
	HTMLAttributions []string `json:"html_attributions"`
	// A token that can be used to return up to 20 additional results. A next_page_token will not be returned if there are no additional results to display. The maximum number of results that can be returned is 60. There is a short delay between when a next_page_token is issued, and when it will become valid.
	NextPageToken *string `json:"next_page_token"`
}

// RankBy specifies the order in which results are listed.
type RankBy string

const (
	// RankByDefault is an alias for RankByProminence
	RankByDefault RankBy = ""
	// RankByProminence sorts results based on their importance. Ranking will favor prominent places within the specified area. Prominence can be affected by a place's ranking in Google's index, global popularity, and other factors.
	RankByProminence RankBy = "prominence"
	// RankByDistance sorts results in ascending order by their distance from the specified location.
	RankByDistance RankBy = "distance"
)
