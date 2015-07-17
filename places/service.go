// Package places provides a client for the Google Places API
package places

import "net/http"

const baseURL = "https://maps.googleapis.com/maps/api/place"

type Service struct {
	client *http.Client
	key    string
	url    string
}

// NewService creates a new places service with the given http client and Google Plus Places API key
func NewService(client *http.Client, key string) *Service {
	return &Service{
		client: client,
		key:    key,
		url:    baseURL,
	}
}

// SetURL allows overwriting the base url
func (s *Service) SetURL(url string) {
	s.url = url
}
