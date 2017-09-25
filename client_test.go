package shopify

import (
	"os"
	"testing"
)

func NewTestClient(t *testing.T) *Client {
	domain := fetchEnvVar(t, "SHOPIFY_TEST_DOMAIN")
	key := fetchEnvVar(t, "SHOPIFY_TEST_KEY")
	pass := fetchEnvVar(t, "SHOPIFY_TEST_PASS")

	return NewClient(domain, key, pass)
}

func fetchEnvVar(t *testing.T, name string) string {
	v := os.Getenv(name)
	if v == "" {
		t.Fatalf("Missing required environment variable '%s'!", name)
	}

	return v
}
