package shopify

import (
	"fmt"
	"net/url"

	"github.com/zpatrick/rclient"
)

type CreateCustomerMetafieldRequest struct {
	Customer CustomerMetafields `json:"customer"`
}

type CustomerMetafields struct {
	ID         int         `json:"id"`
	Metafields []Metafield `json:"metafields"`
}

type Customer struct {
	ID                  int               `json:"id"`
	Email               string            `json:"email"`
	AcceptsMarketing    bool              `json:"accepts_marketing"`
	CreatedAt           string            `json:"created_at"`
	UpdatedAt           string            `json:"updated_at"`
	FirstName           string            `json:"first_name"`
	LastName            string            `json:"last_name"`
	OrdersCount         int               `json:"orders_count"`
	State               string            `json:"state"`
	TotalSpent          string            `json:"total_spent"`
	LastOrderID         int               `json:"last_order_id"`
	Note                string            `json:"note"`
	VerifiedEmail       bool              `json:'verified_email"`
	MultipassIdentifier string            `json:"multipass_identifier"`
	TaxExempt           bool              `json:"tax_exempt"`
	Phone               string            `json:"phone"`
	Tags                string            `json:"tags"`
	LastOrderName       string            `json:"last_order_name"`
	Addresses           []CustomerAddress `json:"addresses"`
	DefaultAddress      CustomerAddress   `json:"default_address"`
}

type CustomerAddress struct {
	ID           int    `json:"id"`
	CustomerID   int    `json:"customer_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Company      string `json:"company"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Zip          string `json:"zip"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	Default      bool   `json:"default"`
}

type GetCustomerResponse struct {
	Customer *Customer `json:"customer"`
}

type GetCustomersResponse struct {
	Customers []*Customer `json:"customers"`
}

func (c *Client) CreateCustomerMetafield(customerID int, metafield Metafield) (*Metafield, error) {
	req := CreateCustomerMetafieldRequest{
		Customer: CustomerMetafields{
			ID:         customerID,
			Metafields: []Metafield{metafield},
		},
	}

	var resp CreateMetafieldResponse
	path := fmt.Sprintf("/admin/customers/%d.json", customerID)
	if err := c.Client.Put(path, req, &resp); err != nil {
		return nil, err
	}

	return resp.Metafield, nil
}

func (c *Client) GetCustomer(customerID int) (*Customer, error) {
	var resp GetCustomerResponse
	path := fmt.Sprintf("/admin/customers/%d.json", customerID)
	if err := c.Client.Get(path, &resp); err != nil {
		return nil, err
	}

	return resp.Customer, nil
}

/*
Query Options:
	ids		A comma-separated list of customer ids
	since_id	Restrict results to after the specified ID
	created_at_min	Show customers created after date (format: 2014-04-25T16:15:47-04:00)
	created_at_max	Show customers created before date (format: 2014-04-25T16:15:47-04:00)
	updated_at_min	Show customers last updated after date (format: 2014-04-25T16:15:47-04:00)
	updated_at_max	Show customers last updated before date (format: 2014-04-25T16:15:47-04:00)
	limit		Amount of results (default: 50) (maximum: 250)
	page		Page to show (default: 1)
	fields		comma-separated list of fields to include in the response
*/
func (c *Client) GetCustomers(query url.Values) ([]*Customer, error) {
	var resp GetCustomersResponse
	if err := c.Client.Get("/admin/customers.json", &resp, rclient.Query(query)); err != nil {
		return nil, err
	}

	return resp.Customers, nil
}

func (c *Client) GetCustomerMetafields(customerID int) ([]Metafield, error) {
	return nil, nil
}
