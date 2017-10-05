package shopify

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrders(t *testing.T) {
	client := NewTestClient(t)
	orders, err := client.GetOrders(url.Values{})
	if err != nil {
		t.Fatal(err)
	}

	if len(orders) == 0 {
		t.Fatal("There are 0 orders, cannot test GetOrder")
	}

	order, err := client.GetOrder(orders[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, orders[0], order)
}
