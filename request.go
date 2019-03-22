package klook

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Token struct {
	GrantType    string `json:"grant_type"`
	ClientId     uint   `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

type SKURequest struct {
	SKUID    uint `json:"sku_id"`
	Quantity uint `json:"quantity"`
}

type Contact struct {
	FirstName  string `json:"first_name"`
	FamilyName string `json:"family_name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Title      string `json:"title"`
	Country    string `json:"country"`
	Language   string `json:"language"`
}

type TravelerInfo struct {
	TypeID uint   `json:"type_id"`
	Value  string `json:"value"`
}

type Traveler struct {
	Traveler uint           `json:"traveler"`
	Info     []TravelerInfo `json:"info"`
}

type AvailabilityRequest struct {
	ProductID       string       `json:"product_id"`
	SKUS            []SKURequest `json:"skus"`
	RequestDate     string       `json:"request_date"`
	RequestTimeslot string       `json:"request_timeslot"`
}

type OrderRequest struct {
	ProductID       string       `json:"product_id"`
	SKUS            []SKURequest `json:"skus"`
	RequestDate     string       `json:"request_date"`
	RequestTimeslot string       `json:"request_timeslot"`
	Travelers       []Traveler   `json:"travelers"`
	Contact         Contact      `json:"contact"`
}

type ScheduleRequest struct {
	ProductID string `json:"product_id"`
	Language  string `json:"lang"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ActivityRequest struct {
	Page      uint `json:"page"`
	PageSize  uint `json:"page_size"`
	CountryID uint `json:"country_id"`
	CityID    uint `json:"city_id"`
}
