package klook

type LoginResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        uint   `json:"expires_in"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type SKU struct {
	SKUID string  `json:"sku_id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

type SKUOrder struct {
	SKUID    uint   `json:"sku_id"`
	Quantity uint   `json:"quantity"`
	SKUPrice string `json:"sku_price"`
	Currency string `json:"currency"`
}

type SKUInventory struct {
	SKUID           string `json:"sku_id"`
	InventoryStatus int    `json:"inventory_status"`
}

type SKUProduct struct {
	SKUID     uint   `json:"sku_id"`
	Title     string `json:"title"`
	SKUMinPax uint   `json:"sku_min_pax"`
	SKUMaxPax uint   `json:"sku_max_pax"`
	Required  bool   `json:"required"`
}

type Schedule struct {
	ProductID      string         `json:"product_id"`
	Date           string         `json:"date"`
	TimeSlot       string         `json:"time_slot"`
	SKUInventories []SKUInventory `json:"sku_inventory"`
}

type Schedules struct {
	Schedules []Schedule    `json:"schedules"`
	SKUS      []SKU         `json:"skus"`
	Error     ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Availablilty struct {
	Success bool          `json:"success"`
	Error   ErrorResponse `json:"error"`
}

type Cancellation struct {
	Success        bool          `json:"success"`
	CancelationID  uint          `json:"cancelation_id"`
	Timestamp      string        `json:"timestamp"`
	RefundAmount   string        `json:"refund_amount"`
	TransactionFee string        `json:"transaction_fee"`
	Currency       string        `json:"currency"`
	Error          ErrorResponse `json:"error"`
}

type CancelationRequest struct {
	RefundID       uint   `json:"refund_id"`
	Timestamp      uint   `json:"timestamp"`
	RefundAmount   string `json:"refund_amount"`
	TransactionFee string `json:"transaction_fee"`
	Currency       string `json:"currency"`
	Status         string `json:"status"`
}

type OrderBooking struct {
	ActivityName     string     `json:"activity_name"`
	PackageName      string     `json:"package_name"`
	ConfirmStatus    string     `json:"confirm_status"`
	VoucherUrl       string     `json:"voucher_url"`
	FileType         uint       `json:"file_type"`
	SKUS             []SKUOrder `json:"skus"`
	BookingRefNumber string     `json:"booking_ref_number"`
}

type OrderDetail struct {
	Success            bool               `json:"success"`
	AgentOrderID       string             `json:"agent_order_id"`
	KLKTechOrderID     string             `json:"klktech_order_id"`
	TotalAmount        string             `json:"total_amount"`
	Currency           string             `json:"currency"`
	CancelationRequest CancelationRequest `json:"cancelation_request"`
	Bookings           []OrderBooking     `json:"bookings"`
	Error              ErrorResponse      `json:"error"`
}

type Order struct {
	Success        bool          `json:"success"`
	AgentOrderID   string        `json:"agent_order_id"`
	KLKTechOrderID string        `json:"klktech_order_id"`
	TotalAmount    string        `json:"total_amount"`
	Currency       string        `json:"currency"`
	ConfirmStatus  string        `json:"confirm_status"`
	SKUS           []SKUOrder    `json:"skus"`
	VoucherURL     string        `json:"voucher_url"`
	Error          ErrorResponse `json:"error"`
}

type City struct {
	CityID uint   `json:"city_id"`
	Name   string `json:"name"`
}

type Country struct {
	CountryID uint   `json:"country_id"`
	Name      string `json:"name"`
	Cities    []City `json:"cities"`
}

type Countries struct {
	Success   bool      `json:"success"`
	Countries []Country `json:"countries"`
}

type ActivityThumbnail struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Width       uint   `json:"width"`
	Height      uint   `json:"height"`
}

type ActivitySimple struct {
	ActivityID  uint              `json:"activity_id"`
	Title       string            `json:"title"`
	Subtitle    string            `json:"subtitle"`
	Description string            `json:"description"`
	CityID      uint              `json:"city_id"`
	City        string            `json:"city"`
	CountryID   uint              `json:"country_id"`
	Country     string            `json:"country"`
	Thumbnail   ActivityThumbnail `json:"thumbnail"`
	Price       float32           `json:"price"`
}

type SubCategory struct {
	CategoryID  uint   `json:"category_id"​`
	Category    string `json:"category"​`
	Description string `json:"description"​`
}

type Activities struct {
	Success    bool             `json:"success"`
	Total      uint             `json:"total"`
	Page       uint             `json:"page"`
	PageSize   uint             `json:"page_size"`
	Activities []ActivitySimple `json:"activities"`
}

// Extra Info General
type ExtraInfoGeneral struct {
	TypeID      uint          `json:"type_id"`
	Type        string        `json:"type"`
	Content     string        `json:"content"`
	Hint        string        `json:"hint"`
	Required    bool          `json:"required"`
	TypeSpecial uint          `json:"type_special"`
	Options     []interface{} `json:"options"`
}

// Extra Info Traveler
type ExtraInfoTraveler struct {
	TypeID   uint   `json:"type_id"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Hint     string `json:"hint"`
	Required bool   `json:"required"`
}

// Extra Info
type ExtraInfo struct {
	General   []ExtraInfoGeneral  `json:"general"`
	Travelers []ExtraInfoTraveler `json:"travelers"`
}

// Product
type Product struct {
	ProductID           string       `json:"product_id"`
	Title               string       `json:"title"`
	Detail              string       `json:"detail"`
	SKUS                []SKUProduct `json:"skus"`
	ExtraInfo           ExtraInfo    `json:"extra_info"`
	ConfirmationType    string       `json:"confirmation_type"`
	CancelationType     string       `json:"cancelation_type"`
	VoucherUsage        string       `json:"voucher_usage"`
	ConfirmationDetails string       `json:"confirmation_details"`
	VoucherTypeDesc     string       `json:"voucher_type_desc"`
	VoucherValidity     string       `json:"voucher_validity"`
	RedemptionProcess   string       `json:"redemption_process"`
	VoucherIdentity     string       `json:"voucher_identity"`
	OpenHours           string       `json:"open_hours"`
	Transportation      string       `json:"transportation"`
	PickUpProcedure     string       `json:"pick_up_procedure"`
	ProductMinPax       uint         `json:"product_min_pax"`
	ProductMaxPax       uint         `json:"product_max_pax"`
}

// Detail Activity
type Activity struct {
	ActivityID          uint              `json:"activity_id"`
	Title               string            `json:"title"`
	Subtitle            string            `json:"subtitle"`
	CityID              uint              `json:"city_id"`
	City                string            `json:"city"`
	CountryID           uint              `json:"country_id"`
	Country             string            `json:"country"`
	CategoryID          uint              `json:"category_id"`
	Category            string            `json:"category"`
	Subcategories       []SubCategory     `json:"subcategories"`
	Duration            string            `json:"duration"`
	Language            string            `json:"language"`
	Travellers          string            `json:"travellers"`
	Transport           string            `json:"transport"`
	Confirmation        string            `json:"confirmation"`
	Refund              string            `json:"refund"`
	TermsConditions     string            `json:"terms_and_conditions"`
	Giudelines          string            `json:"giudelines"`
	Usage               string            `json:"usage"`
	ActivityInformation string            `json:"activity_information"`
	Geolocation         string            `json:"geolocation"`
	VideoURL            string            `json:"video_url"`
	Thumbnail           ActivityThumbnail `json:"thumbnail"`
	Products            []Product         `json:"products"`
}

type DetailActivityResponse struct {
	Success  bool          `json:"success"`
	Activity Activity      `json:"activity"`
	Error    ErrorResponse `json:"error"`
}
