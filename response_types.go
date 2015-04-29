package places

type detailsResponse struct {
	Result PlaceDetails `json:"result"`
	Status string       `json:"status"`
}

type radarSearchResponse struct {
	Results []struct {
		PlaceID string `json:"place_id"`
	} `json:"results"`
	Status string `json:"status"`
}

type HoursTime struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

type OpenInterval struct {
	Open  HoursTime `json:"open"`
	Close HoursTime `json:"close"`
}

type PlaceDetails struct {
	AddressComponents []struct {
		Types     []string `json:"types"`
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
	} `json:"address_components"`
	FormattedAddress     string `json:"formatted_address"`
	FormattedPhoneNumber string `json:"formatted_phone_number"`
	Geometry             struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
	Icon                     string `json:"icon"`
	ID                       string `json:"id"`
	InternationalPhoneNumber string `json:"international_phone_number"`
	Name                     string `json:"name"`
	OpeningHours             struct {
		OpenNow bool           `json:"open_now"`
		Periods []OpenInterval `json:"periods"`
	} `json:"opening_hours"`
	PermanentlyClosed bool `json:"permenantly_closed"`
	Photos            []struct {
		PhotoReference   string   `json:"photo_reference"`
		Height           int      `json:"height"`
		Width            int      `json:"width"`
		HTMLAttributions []string `json:"html_attributions"`
	} `json:"photos"`
	PlaceID string `json:"place_id"`
	Scope   string `json:"scope"`
	AltIDs  []struct {
		PlaceID string `json:"place_id"`
		Scope   string `json:"scope"`
	} `json:"alt_ids"`
	PriceLevel int     `json:"price_level"`
	Rating     float64 `json:"rating"`
	Reviews    []struct {
		Aspects struct {
			Type   string `json:"type"`
			Rating int    `json:"rating"`
		} `json:"aspects"`
		AuthorName string `json:"author_name"`
		AuthorURL  string `json:"author_url"`
		Language   string `json:"language"`
		Rating     int    `json:"rating"`
		Text       string `json:"text"`
		Time       string `json:"time"`
	} `json:"ratings"`
	Types     []string `json:"types"`
	URL       string   `json:"url"`
	UTCOffset int      `json:"utc_offset"`
	Vicinity  string   `json:"vicinity"`
	Website   string   `json:website"`
}
