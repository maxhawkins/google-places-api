package places

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Details returns more comprehensive information about the indicated place such as its complete address, phone number, user rating and reviews.
func (p *Service) Details(placeid string) *DetailsCall {
	return &DetailsCall{
		service: p,
		placeID: placeid,
	}
}

type DetailsCall struct {
	service *Service
	placeID string

	Extensions string
	Language   string
}

func (d *DetailsCall) query() string {
	query := make(url.Values)

	query.Add("key", d.service.key)
	if d.Extensions != "" {
		query.Add("extensions", d.Extensions)
	}
	if d.Language != "" {
		query.Add("language", d.Language)
	}
	query.Add("placeid", d.placeID)

	return query.Encode()
}

func (d *DetailsCall) Do() (*DetailsResponse, error) {
	searchURL := d.service.url + "/details/json?" + d.query()

	resp, err := d.service.client.Get(searchURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad resp %d: %s", resp.StatusCode, body)
	}

	data := &DetailsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
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

type DetailsResponse struct {
	Result           PlaceDetails `json:"result"`
	Status           string       `json:"status"`
	ErrorMessage     string       `json:"error_message"`
	HTMLAttributions []string     `json:"html_attributions"`
}

// DayTime is used in Period to specify opening and closing times.
type DayTime struct {
	// A number from 0–6, corresponding to the days of the week, starting on Sunday. For example, 2 means Tuesday.
	Day int `json:"day"`
	// May contain a time of day in 24-hour hhmm format. Values are in the range 0000–2359. The time will be reported in the place’s time zone.
	Time string `json:"time"`
}

// Period describes a time period when the place is open.
//
// Note: If a place is always open, the close section will be missing from the response. Clients can rely on always-open being represented as an open period containing day with value 0 and time with value 0000, and no close.
type Period struct {
	// A pair of day and time objects describing when the place opens
	Open DayTime `json:"open"`
	// May contain a pair of day and time objects describing when the place closes.
	Close DayTime `json:"close,omitempty"`
	// An array of seven strings representing the formatted opening hours for each day of the week. If a language parameter was specified in the Place Details request, the Places Service will format and localize the opening hours appropriately for that language. The ordering of the elements in this array depends on the language parameter. Some languages start the week on Monday while others start on Sunday.
	WeekdayText []string `json:"weekday_text"`
}

// An AddressComponent is a component used to compose a given address
type AddressComponent struct {
	// An array indicating the type of the address component.
	Types []string `json:"types"`
	// The full text description or name of the address component.
	LongName string `json:"long_name"`
	// An abbreviated textual name for the address component, if available. For example, an address component for the state of Alaska may have a long_name of "Alaska" and a short_name of "AK" using the 2-letter postal abbreviation.
	ShortName string `json:"short_name"`
}

// LatLng contains the geocoded latitude and longitude value for a place.
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Geometry contains a place's location
type Geometry struct {
	Location LatLng `json:"location"`
}

// OpeningHours describes when a place is open.
type OpeningHours struct {
	// A boolean value indicating if the place is open at the current time.
	OpenNow bool `json:"open_now"`
	// An array of opening periods covering seven days, starting from Sunday, in chronological order.
	Periods []Period `json:"periods"`
}

// Photo contains related photographic content to a place
type Photo struct {
	// A string used to identify the photo when you perform a Photo request.
	PhotoReference string `json:"photo_reference"`
	// The maximum height of the image.
	Height int `json:"height"`
	// The maximum width of the image.
	Width int `json:"width"`
	// Contains any required attributions. This field will always be present, but may be empty.
	HTMLAttributions []string `json:"html_attributions"`
}

// AspectRating provides a rating of a single attribute of an establishment
type AspectRating struct {
	// The name of the aspect that is being rated. The following types are supported: appeal, atmosphere, decor, facilities, food, overall, quality and service.
	Type string `json:"type"`
	// The user's rating for this particular aspect, from 0 to 3.
	Rating int `json:"rating"`
}

// Review written about a place
type Review struct {
	// Contains a collection of AspectRating objects, each of which provides a rating of a single attribute of the establishment. The first object in the collection is considered the primary aspect.
	Aspects []AspectRating `json:"aspects"`
	// The name of the user who submitted the review. Anonymous reviews are attributed to "A Google user".
	AuthorName string `json:"author_name"`
	// The URL to the users Google+ profile, if available.
	AuthorURL string `json:"author_url"`
	// An IETF language code indicating the language used in the user's review. This field contains the main language tag only, and not the secondary tag indicating country or region. For example, all the English reviews are tagged as 'en', and not 'en-AU' or 'en-UK' and so on.
	Language string `json:"language"`
	// The user's overall rating for this place. This is a whole number, ranging from 1 to 5.
	Rating int `json:"rating"`
	// The user's review. When reviewing a location with Google Places, text reviews are considered optional. Therefore, this field may by empty. Note that this field may include simple HTML markup. For example, the entity reference &amp; may represent an ampersand character.
	Text string `json:"text"`
	// The time that the review was submitted, measured in the number of seconds since since midnight, January 1, 1970 UTC.
	Time int `json:"time"`
}

// AltID is an alternative place ID for a place, with a scope related to each alternative ID.
type AltID struct {
	// The most likely reason for a place to have an alternative place ID is if your application adds a place and receives an application-scoped place ID, then later receives a Google-scoped place ID after passing the moderation process.
	PlaceID string `json:"place_id"`
	// The scope of an alternative place ID will always be APP, indicating that the alternative place ID is recognised by your application only.
	Scope string `json:"scope"`
}

// PriceLevel is the price level of a place, on a scale of 0 to 4.
type PriceLevel int

// The exact amount indicated by a specific PriceLevel value will vary from region to region.
const (
	Free          PriceLevel = 0
	Inexpensive   PriceLevel = 1
	Moderate      PriceLevel = 2
	Expensive     PriceLevel = 3
	VeryExpensive PriceLevel = 4
)

// PlaceDetails is the information returned by a place details request
type PlaceDetails struct {
	// An array of separate address components used to compose a given address
	AddressComponents []AddressComponent `json:"address_components"`
	// A string containing the human-readable address of this place. Often this address is equivalent to the "postal address," which sometimes differs from country to country.
	FormattedAddress string `json:"formatted_address"`
	// The place's phone number in its local format.
	FormattedPhoneNumber string `json:"formatted_phone_number"`
	// Geometry contains a place's location
	Geometry Geometry `json:"geometry"`
	// The URL of a suggested icon which may be displayed to the user when indicating this result on a map
	Icon string `json:"icon"`
	// The place's phone number in international format
	InternationalPhoneNumber string `json:"international_phone_number"`
	// Contains the human-readable name for the returned result. For establishment results, this is usually the canonicalized business name.
	Name string `json:"name"`
	// Contains information about when the place is open.
	OpeningHours OpeningHours `json:"opening_hours"`
	// A boolean flag indicating whether the place has permanently shut down (value true).
	PermanentlyClosed bool `json:"permenantly_closed"`
	// An array of photo objects, each containing a reference to an image. A Place Details request may return up to ten photos.
	Photos []Photo `json:"photos"`
	// A textual identifier that uniquely identifies a place.
	PlaceID string `json:"place_id"`
	// Indicates the scope of the place_id
	Scope string `json:"scope"`
	// An array of zero, one or more alternative place IDs for the place, with a scope related to each alternative ID.
	AltIDs []AltID `json:"alt_ids"`
	// The price level of the place, on a scale of 0 to 4. The exact amount indicated by a specific value will vary from region to region.
	PriceLevel *PriceLevel `json:"price_level"`
	// The place's rating, from 1.0 to 5.0, based on aggregated user reviews.
	Rating float64 `json:"rating"`
	// Array of up to five reviews. If a language parameter was specified in the Place Details request, the Places Service will bias the results to prefer reviews written in that language.
	Reviews []*Review `json:"reviews"`
	// An array of feature types describing the given result.
	Types []FeatureType `json:"types"`
	// The URL of the official Google page for this place. This will be the establishment's Google+ page if the Google+ page exists, otherwise it will be the Google-owned page that contains the best available information about the place. Applications must link to or embed this page on any screen that shows detailed results about the place to the user.
	URL string `json:"url"`
	// The number of minutes this place’s current timezone is offset from UTC.
	UTCOffset int `json:"utc_offset"`
	// A simplified address for the place, including the street name, street number, and locality, but not the province/state, postal code, or country.
	Vicinity string `json:"vicinity"`
	// The authoritative website for this place, such as a business' homepage.
	Website string `json:website"`
	// Contains a single AspectRating object, for the primary rating of that establishment. (Only available to Google Places API for Work customers.)
	Aspects []AspectRating `json:"aspects"`
	// Indicates that the place has been selected as a Zagat quality location. The Zagat label identifies places known for their consistently high quality or that have a special or unique character. (Only available to Google Places API for Work customers.)
	ZagatSelected bool `json:"zagat_selected"`
}

// FeatureType is a feature type describing a place.
type FeatureType string

// You can use the following values in the types filter for place searches and when adding a place.
const (
	Accounting            FeatureType = "accounting"
	Airport               FeatureType = "airport"
	AmusementPark         FeatureType = "amusement_park"
	Aquarium              FeatureType = "aquarium"
	ArtGallery            FeatureType = "art_gallery"
	Atm                   FeatureType = "atm"
	Bakery                FeatureType = "bakery"
	Bank                  FeatureType = "bank"
	Bar                   FeatureType = "bar"
	BeautySalon           FeatureType = "beauty_salon"
	BicycleStore          FeatureType = "bicycle_store"
	BookStore             FeatureType = "book_store"
	BowlingAlley          FeatureType = "bowling_alley"
	BusStation            FeatureType = "bus_station"
	Cafe                  FeatureType = "cafe"
	Campground            FeatureType = "campground"
	CarDealer             FeatureType = "car_dealer"
	CarRental             FeatureType = "car_rental"
	CarRepair             FeatureType = "car_repair"
	CarWash               FeatureType = "car_wash"
	Casino                FeatureType = "casino"
	Cemetery              FeatureType = "cemetery"
	Church                FeatureType = "church"
	CityHall              FeatureType = "city_hall"
	ClothingStore         FeatureType = "clothing_store"
	ConvenienceStore      FeatureType = "convenience_store"
	Courthouse            FeatureType = "courthouse"
	Dentist               FeatureType = "dentist"
	DepartmentStore       FeatureType = "department_store"
	Doctor                FeatureType = "doctor"
	Electrician           FeatureType = "electrician"
	ElectronicsStore      FeatureType = "electronics_store"
	Embassy               FeatureType = "embassy"
	Establishment         FeatureType = "establishment"
	Finance               FeatureType = "finance"
	FireStation           FeatureType = "fire_station"
	Florist               FeatureType = "florist"
	Food                  FeatureType = "food"
	FuneralHome           FeatureType = "funeral_home"
	FurnitureStore        FeatureType = "furniture_store"
	GasStation            FeatureType = "gas_station"
	GeneralContractor     FeatureType = "general_contractor"
	GroceryOrSupermarket  FeatureType = "grocery_or_supermarket"
	Gym                   FeatureType = "gym"
	HairCare              FeatureType = "hair_care"
	HardwareStore         FeatureType = "hardware_store"
	Health                FeatureType = "health"
	HinduTemple           FeatureType = "hindu_temple"
	HomeGoodsStore        FeatureType = "home_goods_store"
	Hospital              FeatureType = "hospital"
	InsuranceAgency       FeatureType = "insurance_agency"
	JewelryStore          FeatureType = "jewelry_store"
	Laundry               FeatureType = "laundry"
	Lawyer                FeatureType = "lawyer"
	Library               FeatureType = "library"
	LiquorStore           FeatureType = "liquor_store"
	LocalGovernmentOffice FeatureType = "local_government_office"
	Locksmith             FeatureType = "locksmith"
	Lodging               FeatureType = "lodging"
	MealDelivery          FeatureType = "meal_delivery"
	MealTakeaway          FeatureType = "meal_takeaway"
	Mosque                FeatureType = "mosque"
	MovieRental           FeatureType = "movie_rental"
	MovieTheater          FeatureType = "movie_theater"
	MovingCompany         FeatureType = "moving_company"
	Museum                FeatureType = "museum"
	NightClub             FeatureType = "night_club"
	Painter               FeatureType = "painter"
	Park                  FeatureType = "park"
	Parking               FeatureType = "parking"
	PetStore              FeatureType = "pet_store"
	Pharmacy              FeatureType = "pharmacy"
	Physiotherapist       FeatureType = "physiotherapist"
	PlaceOfWorship        FeatureType = "place_of_worship"
	Plumber               FeatureType = "plumber"
	Police                FeatureType = "police"
	PostOffice            FeatureType = "post_office"
	RealEstateAgency      FeatureType = "real_estate_agency"
	Restaurant            FeatureType = "restaurant"
	RoofingContractor     FeatureType = "roofing_contractor"
	RvPark                FeatureType = "rv_park"
	School                FeatureType = "school"
	ShoeStore             FeatureType = "shoe_store"
	ShoppingMall          FeatureType = "shopping_mall"
	Spa                   FeatureType = "spa"
	Stadium               FeatureType = "stadium"
	Storage               FeatureType = "storage"
	Store                 FeatureType = "store"
	SubwayStation         FeatureType = "subway_station"
	Synagogue             FeatureType = "synagogue"
	TaxiStand             FeatureType = "taxi_stand"
	TrainStation          FeatureType = "train_station"
	TravelAgency          FeatureType = "travel_agency"
	University            FeatureType = "university"
	VeterinaryCare        FeatureType = "veterinary_care"
	Zoo                   FeatureType = "zoo"
)
