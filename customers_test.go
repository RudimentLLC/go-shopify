package shopify

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomers(t *testing.T) {
	client := NewTestClient(t)
	customers, err := client.GetCustomers(url.Values{})
	if err != nil {
		t.Fatal(err)
	}

	if len(customers) == 0 {
		t.Fatal("There are 0 customers, cannot test GetCustomer")
	}

	customer, err := client.GetCustomer(customers[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, customers[0], customer)
}
