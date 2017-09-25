package shopify

import "net/url"

type ShopifyClient interface {
	GetCustomer(customerID int) (*Customer, error)
	GetCustomers(query url.Values) ([]*Customer, error)
}
