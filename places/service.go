package places

import "net/http"

const baseURL = "https://maps.googleapis.com/maps/api/place"

type Service struct {
	client *http.Client
	key    string
}

// NewService creates a new places service with the given http client and Google Plus Places API key
func NewService(client *http.Client, key string) *Service {
	return &Service{
		client: client,
		key:    key,
	}
}
