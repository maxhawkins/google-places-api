package places

type RankBy string

const (
	RankByProminence RankBy = "prominence"
	RankByDistance   RankBy = "distance"
)

type Price int

const (
	// PriceUnspecified is used when you don't care about the price
	PriceUnspecified Price = iota
	// Price0 maps to $
	Price0
	// Price1 maps to $$
	Price1
	// Price2 maps to $$$
	Price2
	// Price3 maps to $$$$
	Price3
	// Price4 maps to $$$$$
	Price4
)

// Type is used for the types property in the Google Places API
type Type string

// Types you can use in place searches and place additions
const (
	Accounting            Type = "accounting"
	Airport               Type = "airport"
	AmusementPark         Type = "amusement_park"
	Aquarium              Type = "aquarium"
	ArtGallery            Type = "art_gallery"
	Atm                   Type = "atm"
	Bakery                Type = "bakery"
	Bank                  Type = "bank"
	Bar                   Type = "bar"
	BeautySalon           Type = "beauty_salon"
	BicycleStore          Type = "bicycle_store"
	BookStore             Type = "book_store"
	BowlingAlley          Type = "bowling_alley"
	BusStation            Type = "bus_station"
	Cafe                  Type = "cafe"
	Campground            Type = "campground"
	CarDealer             Type = "car_dealer"
	CarRental             Type = "car_rental"
	CarRepair             Type = "car_repair"
	CarWash               Type = "car_wash"
	Casino                Type = "casino"
	Cemetery              Type = "cemetery"
	Church                Type = "church"
	CityHall              Type = "city_hall"
	ClothingStore         Type = "clothing_store"
	ConvenienceStore      Type = "convenience_store"
	Courthouse            Type = "courthouse"
	Dentist               Type = "dentist"
	DepartmentStore       Type = "department_store"
	Doctor                Type = "doctor"
	Electrician           Type = "electrician"
	ElectronicsStore      Type = "electronics_store"
	Embassy               Type = "embassy"
	Establishment         Type = "establishment"
	Finance               Type = "finance"
	FireStation           Type = "fire_station"
	Florist               Type = "florist"
	Food                  Type = "food"
	FuneralHome           Type = "funeral_home"
	FurnitureStore        Type = "furniture_store"
	GasStation            Type = "gas_station"
	GeneralContractor     Type = "general_contractor"
	GroceryOrSupermarket  Type = "grocery_or_supermarket"
	Gym                   Type = "gym"
	HairCare              Type = "hair_care"
	HardwareStore         Type = "hardware_store"
	Health                Type = "health"
	HinduTemple           Type = "hindu_temple"
	HomeGoodsStore        Type = "home_goods_store"
	Hospital              Type = "hospital"
	InsuranceAgency       Type = "insurance_agency"
	JewelryStore          Type = "jewelry_store"
	Laundry               Type = "laundry"
	Lawyer                Type = "lawyer"
	Library               Type = "library"
	LiquorStore           Type = "liquor_store"
	LocalGovernmentOffice Type = "local_government_office"
	Locksmith             Type = "locksmith"
	Lodging               Type = "lodging"
	MealDelivery          Type = "meal_delivery"
	MealTakeaway          Type = "meal_takeaway"
	Mosque                Type = "mosque"
	MovieRental           Type = "movie_rental"
	MovieTheater          Type = "movie_theater"
	MovingCompany         Type = "moving_company"
	Museum                Type = "museum"
	NightClub             Type = "night_club"
	Painter               Type = "painter"
	Park                  Type = "park"
	Parking               Type = "parking"
	PetStore              Type = "pet_store"
	Pharmacy              Type = "pharmacy"
	Physiotherapist       Type = "physiotherapist"
	PlaceOfWorship        Type = "place_of_worship"
	Plumber               Type = "plumber"
	Police                Type = "police"
	PostOffice            Type = "post_office"
	RealEstateAgency      Type = "real_estate_agency"
	Restaurant            Type = "restaurant"
	RoofingContractor     Type = "roofing_contractor"
	RvPark                Type = "rv_park"
	School                Type = "school"
	ShoeStore             Type = "shoe_store"
	ShoppingMall          Type = "shopping_mall"
	Spa                   Type = "spa"
	Stadium               Type = "stadium"
	Storage               Type = "storage"
	Store                 Type = "store"
	SubwayStation         Type = "subway_station"
	Synagogue             Type = "synagogue"
	TaxiStand             Type = "taxi_stand"
	TrainStation          Type = "train_station"
	TravelAgency          Type = "travel_agency"
	University            Type = "university"
	VeterinaryCare        Type = "veterinary_care"
	Zoo                   Type = "zoo"
)
