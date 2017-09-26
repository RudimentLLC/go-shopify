package shopify

import (
	"fmt"
	"net/url"

	"github.com/zpatrick/rclient"
)

type Product struct {
	ID             int              `json:"id"`
	Title          string           `json:"title"`
	BodyHTML       string           `json:"body_html"`
	Vendor         string           `json:"vendor"`
	ProductType    string           `json:"product_type"`
	CreatedAt      string           `json:"created_at"`
	Handle         string           `json:"handle"`
	UpdatedAt      string           `json:"updated_at"`
	PublishedAt    string           `json:"published_at"`
	TemplateSuffix string           `json:"template_suffix"`
	PublishedScope string           `json:"published_scope"`
	Tags           string           `json:"tags"`
	Variants       []ProductVariant `json:"variants"`
	Options        []ProductOption  `json:"options"`
	Images         []ProductImage   `json:"images"`
	Image          ProductImage     `json:"image"`
}

type ProductImage struct {
	ID         int    `json:"id"`
	ProductID  int    `json:"product_id"`
	Position   int    `json:"position"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	SRC        string `json:"src"`
	VariantIDs []int  `json:"variant_ids"`
}

type ProductOption struct {
	ID        int      `json:"id"`
	ProductID int      `json:"product_id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}

type ProductVariant struct {
	ID                   int     `json:"id"`
	ProductID            int     `json:"product_id"`
	Title                string  `json:"title"`
	Price                string  `json:"price"`
	SKU                  string  `json:"sku"`
	Position             int     `json:"position"`
	Grams                int64   `json:"grams"`
	InventoryPolicy      string  `json:"inventory_policy"`
	CompareAtPrice       string  `json:"compare_at_price"`
	FulfillmentService   string  `json:"fulfillment_service"`
	InventoryManagement  string  `json:"inventory_management"`
	Option1              string  `json:"option1"`
	Option2              string  `json:"option2"`
	Option3              string  `json:"option3"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	Taxable              bool    `json:"taxable"`
	Barcode              string  `json:"barcode"`
	ImageID              int     `json:"image_id"`
	InventoryQuantity    int64   `json:"inventory_quantity"`
	Weight               float64 `json:"weight"`
	WeightUnit           string  `json:"weight_unit"`
	OldInventoryQuantity int64   `json:"old_inventory_quantity"`
	RequiresShipping     bool    `json:"requires_shipping"`
}

type GetProductResponse struct {
	Product *Product `json:"product"`
}

type GetProductsResponse struct {
	Products []*Product `json:"products"`
}

func (c *Client) GetProduct(productID int) (*Product, error) {
	var resp GetProductResponse
	path := fmt.Sprintf("/admin/products/%d.json", productID)
	if err := c.Client.Get(path, &resp); err != nil {
		return nil, err
	}

	return resp.Product, nil
}

/*
Query Options:
	ids			Comma-separated list of product ids
	limit			Amount of results (default: 50) (maximum: 250)
	page			Page to show (default: 1)
	since_id		Restrict results to after the specified ID
	title			Filter by product title
	vendor			Filter by product vendor
	handle			Filter by product handle
	product_type		Filter by product type
	collection_id		Filter by collection id
	created_at_min		Show products created after date (format: 2014-04-25T16:15:47-04:00)
	created_at_max		Show products created before date (format: 2014-04-25T16:15:47-04:00)
	updated_at_min		Show products last updated after date (format: 2014-04-25T16:15:47-04:00)
	updated_at_max		Show products last updated before date (format: 2014-04-25T16:15:47-04:00)
	published_at_min	Show products published after date (format: 2014-04-25T16:15:47-04:00)
	published_at_max	Show products published before date (format: 2014-04-25T16:15:47-04:00)
	published_status	published 	Show only published products
				unpublished  	Show only unpublished products
				any 		Show all products (default)
	fields	Comma-separated list of fields to include in the response
*/

func (c *Client) GetProducts(query url.Values) ([]*Product, error) {
	var resp GetProductsResponse
	if err := c.Client.Get("/admin/products.json", &resp, rclient.Query(query)); err != nil {
		return nil, err
	}

	return resp.Products, nil
}
