package places

import (
	"net/http"
	"testing"
)

func TestNewService(t *testing.T) {
	for _, test := range []struct {
		Client *http.Client
		Key    string
		Want   Service
	}{
		{
			Client: http.DefaultClient,
			Key:    "key",
			Want: Service{
				client: http.DefaultClient,
				key:    "key",
				url:    "https://maps.googleapis.com/maps/api/place",
			},
		},
	} {
		service := NewService(test.Client, test.Key)

		if *service != test.Want {
			t.Errorf("NewService() %#v = %#v", service, test.Want)
		}
	}
}

func TestSetURL(t *testing.T) {
	url := "https://localhost/maps/api/place"

	service := NewService(http.DefaultClient, "key")
	service.SetURL(url)

	if service.url != url {
		t.Errorf("Service{}.SetURL() %#v = %#v", service.url, url)
	}
}
