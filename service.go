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

const baseURL = "https://maps.googleapis.com/maps/api/place"

type Service struct {
	Client *http.Client
	Key    string
}

func NewService(client *http.Client, key string) *Service {
	return &Service{
		Client: client,
		Key:    key,
	}
}

type NearbyCall struct {
	service *Service

	Keyword            string
	Language           string
	Lat, Long          float64
	MinPrice, MaxPrice Price
	Name               string
	OpenNow            bool
	Radius             float64
	RankBy             RankBy
	Types              []Type
}

func (p *Service) Nearby(lat, long float64) *NearbyCall {
	return &NearbyCall{
		service: p,
		Lat:     lat,
		Long:    long,
	}
}

// TODO(maxhawkins): add search result struct and return results
// from nearby & radar

func (n *NearbyCall) Do() (placeIDs []string, err error) {
	if n.RankBy == RankByProminence && n.Radius == 0 {
		return nil, errors.New("radius must be specified when RankByProminence is used")
	}
	if n.RankBy == RankByDistance &&
		n.Types == nil &&
		n.Name == "" &&
		n.Keyword == "" {
		return nil, errors.New("when RankByDistance is specified, one or more of keyword, name, or types is required")
	}

	searchURL := baseURL + "/nearbysearch/json?" + n.query()

	resp, err := n.service.Client.Get(searchURL)
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

	var data radarSearchResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	switch data.Status {
	case "OK":
		break
	case "ZERO_RESULTS":
		return nil, nil
	default:
		return nil, fmt.Errorf("search failed: %s", body)
	}

	for _, r := range data.Results {
		placeIDs = append(placeIDs, r.PlaceID)
	}

	return placeIDs, nil
}

func (r *NearbyCall) query() string {
	query := make(url.Values)
	query.Add("key", r.service.Key)
	if r.Keyword != "" {
		query.Add("keyword", r.Keyword)
	}
	if r.Language != "" {
		query.Add("language", r.Language)
	}
	query.Add("location", fmt.Sprintf("%f,%f", r.Lat, r.Long))
	if r.MinPrice != PriceUnspecified {
		n := int(r.MinPrice) - 1
		query.Add("minprice", fmt.Sprint(n))
	}
	if r.MaxPrice != PriceUnspecified {
		n := int(r.MaxPrice) - 1
		query.Add("maxprice", fmt.Sprint(n))
	}
	if r.Name != "" {
		query.Add("name", r.Name)
	}
	if r.OpenNow {
		query.Add("opennow", fmt.Sprint(1))
	}
	if r.RankBy != RankByDistance {
		query.Add("radius", fmt.Sprint(r.Radius))
	}
	query.Add("rankby", string(r.RankBy))

	var typeNames []string
	for _, t := range r.Types {
		typeNames = append(typeNames, string(t))
	}
	query.Add("types", strings.Join(typeNames, ","))

	return query.Encode()
}

func (p *Service) Details(placeid string) *DetailsCall {
	return &DetailsCall{
		service: p,
		PlaceID: placeid,
	}
}

type DetailsCall struct {
	service *Service

	Extensions string
	Language   string
	PlaceID    string
}

func (d *DetailsCall) query() string {
	query := make(url.Values)

	query.Add("key", d.service.Key)
	if d.Extensions != "" {
		query.Add("extensions", d.Extensions)
	}
	if d.Language != "" {
		query.Add("language", d.Language)
	}
	query.Add("placeid", d.PlaceID)

	return query.Encode()
}

func (d *DetailsCall) Do() (*PlaceDetails, error) {
	searchURL := baseURL + "/details/json?" + d.query()

	resp, err := d.service.Client.Get(searchURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad resp %d: %s", resp.StatusCode, body)
	}

	var data detailsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Status != "OK" {
		return nil, fmt.Errorf("lookup failed: %s", data.Status)
	}

	return &data.Result, nil
}

func (p *Service) RadarSearch(radius, lat, long float64) *RadarSearchCall {
	return &RadarSearchCall{
		service: p,
		Radius:  radius,
		Lat:     lat,
		Long:    long,
	}
}

type RadarSearchCall struct {
	service *Service

	Keyword            string
	Lat, Long          float64
	MinPrice, MaxPrice Price
	OpenNow            bool
	Radius             float64
	Types              []Type
}

func (r *RadarSearchCall) query() string {
	query := make(url.Values)
	query.Add("key", r.service.Key)
	if r.Keyword != "" {
		query.Add("keyword", r.Keyword)
	}
	query.Add("location", fmt.Sprintf("%f,%f", r.Lat, r.Long))
	if r.MinPrice != PriceUnspecified {
		n := int(r.MinPrice) - 1
		query.Add("minprice", fmt.Sprint(n))
	}
	if r.MaxPrice != PriceUnspecified {
		n := int(r.MaxPrice) - 1
		query.Add("maxprice", fmt.Sprint(n))
	}
	if r.OpenNow {
		query.Add("opennow", fmt.Sprint(1))
	}
	query.Add("radius", fmt.Sprint(r.Radius))

	var typeNames []string
	for _, t := range r.Types {
		typeNames = append(typeNames, string(t))
	}
	query.Add("types", strings.Join(typeNames, ","))

	return query.Encode()
}

func (r *RadarSearchCall) Do() (placeIDs []string, err error) {
	searchURL := baseURL + "/radarsearch/json?" + r.query()

	resp, err := r.service.Client.Get(searchURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad resp %d: %s", resp.StatusCode, body)
	}

	var data radarSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	switch data.Status {
	case "OK":
		break
	case "ZERO_RESULTS":
		return nil, nil
	default:
		return nil, fmt.Errorf("search failed: %s", data.Status)
	}

	for _, r := range data.Results {
		placeIDs = append(placeIDs, r.PlaceID)
	}

	return placeIDs, nil
}
