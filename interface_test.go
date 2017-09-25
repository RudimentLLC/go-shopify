package shopify

import "testing"

func TestClientSatisfiesInterface(t *testing.T) {
	var _ ShopifyClient = &Client{}
}
