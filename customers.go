package shopify

import (
	"fmt"
	"net/url"

	"github.com/zpatrick/rclient"
)

type Customer struct {
	ID                  int               `json:"id,omitempty"`
	Email               string            `json:"email,omitempty"`
	AcceptsMarketing    bool              `json:"accepts_marketing,omitempty"`
	CreatedAt           string            `json:"created_at,omitempty"`
	UpdatedAt           string            `json:"updated_at,omitempty"`
	FirstName           string            `json:"first_name,omitempty"`
	LastName            string            `json:"last_name,omitempty"`
	OrdersCount         int               `json:"orders_count"`
	State               string            `json:"state,omitempty"`
	TotalSpent          string            `json:"total_spent,omitempty"`
	LastOrderID         int               `json:"last_order_id,omitempty"`
	Note                string            `json:"note,omitempty"`
	VerifiedEmail       bool              `json:'verified_email,omitempty"`
	MultipassIdentifier string            `json:"multipass_identifier,omitempty"`
	TaxExempt           bool              `json:"tax_exempt,omitempty"`
	Phone               string            `json:"phone,omitempty"`
	Tags                string            `json:"tags"`
	LastOrderName       string            `json:"last_order_name,omitempty"`
	Addresses           []CustomerAddress `json:"addresses,omitempty"`
	DefaultAddress      CustomerAddress   `json:"default_address,omitempty"`
}

type CustomerAddress struct {
	ID           int    `json:"id,omitempty"`
	CustomerID   int    `json:"customer_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Company      string `json:"company,omitempty"`
	Address1     string `json:"address1,omitempty"`
	Address2     string `json:"address2,omitempty"`
	City         string `json:"city,omitempty"`
	Province     string `json:"province,omitempty"`
	Country      string `json:"country,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Name         string `json:"name,omitempty"`
	ProvinceCode string `json:"province_code,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
	CountryName  string `json:"country_name,omitempty"`
	Default      bool   `json:"default,omitempty"`
}

type GetCustomerResponse struct {
	Customer *Customer `json:"customer"`
}

type GetCustomersResponse struct {
	Customers []*Customer `json:"customers"`
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
	page		Page to show(default: 1)
	fields		comma-separated list of fields to include in the response
*/
func (c *Client) GetCustomers(query url.Values) ([]*Customer, error) {
	var resp GetCustomersResponse
	if err := c.Client.Get("/admin/customers.json", &resp, rclient.Query(query)); err != nil {
		return nil, err
	}

	return resp.Customers, nil
}
