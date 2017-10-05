package shopify

import (
	"fmt"
	"net/url"

	"github.com/zpatrick/rclient"
)

type ClientDetail struct {
	BrowserIP      string      `json:"browser_ip"`
	AcceptLanguage interface{} `json:"accept_language"`
	UserAgent      interface{} `json:"user_agent"`
	SessionHash    interface{} `json:"session_hash"`
	BrowserWidth   interface{} `json:"browser_width"`
	BrowserHeight  interface{} `json:"browser_height"`
}

type DiscountCode struct {
	Code   string `json:"code"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type Fulfillment struct {
	ID              int          `json:"id"`
	OrderID         int          `json:"order_id"`
	Status          string       `json:"status"`
	CreatedAt       string       `json:"created_at"`
	Service         string       `json:"service"`
	UpdatedAt       string       `json:"upated_at"`
	TrackingCompany string       `json:"tracking_company"`
	ShipmentStatus  string       `json:"shipment_status"`
	TrackingNumber  string       `json:"tracking_number"`
	TrackingNumbers []string     `json:"tracking_numbers"`
	TrackingURL     string       `json:"tracking_url"`
	TrackingURLs    []string     `json:"tracking_urls"`
	Receipt         Receipt      `json:"receipt"`
	LineItems       []LineItem   `json:"line_items"`
	ClientDetails   ClientDetail `json:"client_detail"`
	Refunds         []Refund     `json:"refunds"`
}

type GetOrderResponse struct {
	Order *Order `json:"order"`
}

type GetOrdersResponse struct {
	Orders []*Order `json:"orders"`
}

type LineItem struct {
	ID                         int        `json:"id"`
	VariantID                  int        `json:"variant_id"`
	Title                      string     `json:"title"`
	Quantity                   int        `json:"quantity"`
	Price                      string     `json:"price"`
	Grams                      int        `json:"grams"`
	SKU                        string     `json:"sku"`
	VariantTitle               string     `json:"variant_title"`
	Vendor                     string     `json:"vendor"`
	FulfillmentService         string     `json:"fulfillment_service"`
	ProductID                  int        `json:"product_id"`
	RequiresShipping           bool       `json:"requires_shipping"`
	Taxable                    bool       `json:"taxable"`
	GiftCard                   bool       `json:"gift_card"`
	Name                       string     `json:"name"`
	VariantInventoryManagement string     `json:"variant_inventory_management"`
	Properties                 []Property `json:"properties"`
	ProductExists              bool       `json:"product_exists"`
	FulfillableQuantity        int        `json:"fulfillable_quantity"`
	TotalDiscount              string     `json:"total_discount"`
	FulfillmentStatus          string     `json:"fulfillment_status"`
	TaxLines                   []TaxLine  `json:"tax_lines"`
}

type NoteAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//https://github.com/getconversio/go-shopify/blob/00f23dd9acd384af91576477e3d152497f475c56/order.go#L29
type Order struct {
	ID                    int               `json:"id,omitempty"`
	Email                 string            `json:"email"`
	ClosedAt              string            `json:"closed_at"`
	CreatedAt             string            `json:"created_at"`
	UpdatedAt             string            `json:"updated_at"`
	Number                int               `json:"number"`
	Note                  string            `json:"note"`
	Token                 string            `json:"token"`
	Gateway               string            `json:"gateway"`
	Test                  bool              `json:"test"`
	TotalPrice            string            `json:"total_price"`
	SubtotalPrice         string            `json:"subtotal_price"`
	TotalWeight           float64           `json:"total_weight"`
	TotalTax              string            `json:"total_tax"`
	TaxIncluded           bool              `json:"tax_included"`
	Currency              string            `json:"currency"`
	FinancialStatus       string            `json:"financial_status"`
	Confirmed             bool              `json:"confirmed"`
	TotalDiscounts        string            `json:"total_discounts"`
	TotalLineItemsPrice   string            `json:"total_line_items_price"`
	CartToken             string            `json:"cart_token"`
	BuyerAcceptsMarketing bool              `json:"buyer_accepts_marketing"`
	Name                  string            `json:"name"`
	ReferringSite         string            `json:"referring_site"`
	LandingSite           string            `json:"landing_site"`
	CancelledAt           string            `json:"cancelled_at"`
	CancelReason          string            `json:"cancel_reason"`
	TotalPriceUSD         string            `json:"total_price_usd"`
	CheckoutToken         string            `json:"checkout_token"`
	Reference             string            `json:"reference"`
	UserID                string            `json:"user_id"`
	LocationID            string            `json:"location_id"`
	SourceIdentifier      string            `json:"source_identifier"`
	SourceURL             string            `json:"source_url"`
	ProcessedAt           string            `json:"processed_at"`
	DeviceID              string            `json:"device_id"`
	Phone                 string            `json:"phone"`
	CustomerLocale        string            `json:"customer_locale"`
	AppID                 int               `json:"app_id"`
	BrowserIP             string            `json:"browser_ip"`
	LandingSiteRef        string            `json:"landing_site_ref"`
	OrderNumber           int               `json:"order_number"`
	DiscountCodes         []DiscountCode    `json:"discount_code"`
	NoteAttributes        []NoteAttribute   `json:"note_attributes"`
	PaymentGatewayNames   []string          `json:"payment_gateway_names"`
	ProcessingMethod      string            `json:"processing_method"`
	CheckoutID            int               `json:"checkout_id"`
	SourceName            string            `json:"source_name"`
	FullfillmentStatus    string            `json:"fulfillment_status"`
	TaxLines              []TaxLine         `json:"tax_lines"`
	Tags                  string            `json:"tags"`
	ContactEmail          string            `json:"contact_email"`
	OrderStatusURL        string            `json:"order_status_url"`
	LineItems             []LineItem        `json:"line_items"`
	ShippingLines         []ShippingLine    `json:"shipping_lines"`
	BillingAddress        OrderAddress      `json:"billing_address"`
	ShippingAddress       OrderAddress      `json:"shipping_address"`
	Fullfillments         []Fulfillment     `json:"fullfillments"`
	ClientDetails         ClientDetail      `json:"client_details"`
	Refunds               []Refund          `json:"refunds"`
	Transactions          []Transaction     `json:"transactions"`
	OrderAdjustments      []OrderAdjustment `json:"order_adjustments"`
}

type OrderAddress struct {
	FirstName    string  `json:"first_name"`
	Address1     string  `json:"address1"`
	Phone        string  `json:"phone"`
	City         string  `json:"city"`
	Zip          string  `json:"zip"`
	Province     string  `json:"province"`
	Country      string  `json:"country"`
	LastName     string  `json:"last_name"`
	Address2     string  `json:"address2"`
	Company      string  `json:"company"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"country_code"`
	ProvinceCode string  `json:"province_code"`
}

type OrderAdjustment interface{}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Receipt struct {
	TestCase      bool   `json:"testcase"`
	Authorization string `json:"authorization"`
}

type Refund struct {
	ID              int           `json:"id"`
	OrderID         int           `json:"order_id"`
	CreatedAt       string        `json:"created_at"`
	Note            string        `json:"note"`
	Restock         bool          `json:"restock"`
	UserID          int           `json:"user_id"`
	ProcessedAt     string        `json:"processed_at"`
	RefundLineItems []LineItem    `json:"refund_line_items"`
	Transactions    []Transaction `json:"transaction"`
}

type ShippingLine struct {
	ID                            int       `json:"id"`
	Title                         string    `json:"title"`
	Price                         string    `json:"price"`
	Code                          string    `json:"code"`
	Source                        string    `json:"source"`
	Phone                         string    `json:"phone"`
	RequestedFulfillmentServiceID string    `json:"requested_fulfillment_service_id"`
	DeliveryCategory              string    `json:"delivery_category"`
	CarrierIdentifier             string    `json:"carrier_identifier"`
	TaxLines                      []TaxLine `json:"tax_lines"`
}

type TaxLine struct {
	Title string  `json:"title"`
	Price string  `json:"price"`
	Rate  float64 `json:"float"`
}

type Transaction struct {
	ID            int         `json:"id"`
	OrderID       int         `json:"order_id"`
	Amount        string      `json:"amount"`
	Kind          string      `json:"kind"`
	Gateway       string      `json:"gateway"`
	Status        string      `json:"status"`
	Message       interface{} `json:"message"`
	CreatedAt     string      `json:"created_at"`
	Test          bool        `json:"test"`
	Authorization string      `json:"authorization"`
	Currency      string      `json:"currency"`
	LocationID    interface{} `json:"location_id"`
	UserID        interface{} `json:"user_id"`
	ParentID      int         `json:"parent_id"`
	DeviceID      interface{} `json:"device_id"`
	Receipt       Receipt     `json:"receipt"`
	ErrorCode     interface{} `json:"error_code"`
	SourceName    string      `json:"source_name"`
}

func (c *Client) GetOrder(orderID int) (*Order, error) {
	var resp GetOrderResponse
	path := fmt.Sprintf("/admin/orders/%d.json", orderID)
	if err := c.Client.Get(path, &resp); err != nil {
		return nil, err
	}

	return resp.Order, nil
}

/*
Query Options:
	ids			A comma-separated list of order ids
	limit			Amount of results(default: 50) (maximum: 250)
	page			Page to show(default: 1)
	since_id		Restrict results to after the specified ID
	created_at_min		Show orders created after date (format: 2014-04-25T16:15:47-04:00)
	created_at_max		Show orders created before date (format: 2014-04-25T16:15:47-04:00)
	updated_at_min		Show orders last updated after date (format: 2014-04-25T16:15:47-04:00)
	updated_at_max		Show orders last updated before date (format: 2014-04-25T16:15:47-04:00)
	processed_at_min	Show orders imported after date (format: 2014-04-25T16:15:47-04:00)
	processed_at_max	Show orders imported before date (format: 2014-04-25T16:15:47-04:00)
	attribution_app_id	Show orders attributed to a specific app.
	status
		open		All open orders (default)
		closed		Show only closed orders
		cancelled	Show only cancelled orders
		any		Any order status
	financial_status
		authorized		Show only authorized orders
		pending			Show only pending orders
		paid			Show only paid orders
		partially_paid		Show only partially paid orders
		refunded		Show only refunded orders
		voided			Show only voided orders
		partially_refunded	Show only partially_refunded orders
		any			Show all authorized, pending, and paid orders (default). This is a filter, not a value.
		unpaid 			Show all authorized, or partially_paid orders. This is a filter, not a value.
	fulfillment_status
		shipped		Show orders that have been shipped
		partial		Show partially shipped orders
		unshipped	Show orders that have not yet been shipped
		any		Show orders with any fulfillment_status. (default)
	fields	comma-separated list of fields to include in the response
*/

func (c *Client) GetOrders(query url.Values) ([]*Order, error) {
	var resp GetOrdersResponse
	if err := c.Client.Get("/admin/orders.json", &resp, rclient.Query(query)); err != nil {
		return nil, err
	}

	return resp.Orders, nil
}
