package shopify

import "net/url"

type ShopifyClient interface {
	CreateCustomerMetafield(customerID int, metafield Metafield) (*Metafield, error)
	CreatePage(p Page) (*Page, error)
	CreatePageMetafield(pageID int, metafield Metafield) (*Page, error)
	DeleteCustomerMetafield(customerID, metafieldID int) error
	DeletePage(pageID int) error
	GetCustomer(customerID int) (*Customer, error)
	GetCustomerMetafields(customerID int) ([]Metafield, error)
	GetCustomers(query url.Values) ([]*Customer, error)
	GetOrder(orderID int) (*Order, error)
	GetOrders(query url.Values) ([]*Order, error)
	GetProduct(productID int) (*Product, error)
	GetProducts(query url.Values) ([]*Product, error)
}
