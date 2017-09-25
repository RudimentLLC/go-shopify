package shopify

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
client.Client.ResponseReader = func(resp *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Body: %v\n", string(body))
	return nil
}
*/

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
