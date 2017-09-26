package shopify

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	client := NewTestClient(t)
	products, err := client.GetProducts(url.Values{})
	if err != nil {
		t.Fatal(err)
	}

	if len(products) == 0 {
		t.Fatal("There are 0 products, cannot test GetProduct")
	}

	product, err := client.GetProduct(products[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, products[0], product)
}
