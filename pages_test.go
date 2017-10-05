package shopify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePage(t *testing.T) {
	client := NewTestClient(t)

	req := Page{
		Title:     "title",
		BodyHTML:  "body_html",
		Published: false,
		Metafields: []Metafield{
			{
				Key:       "key",
				Value:     "val",
				ValueType: "string",
				Namespace: "namespace",
			},
		},
	}

	output, err := client.CreatePage(req)
	if err != nil {
		t.Fatal(err)
	}
	defer client.DeletePage(output.ID)

	assert.Equal(t, req.Title, output.Title)
	assert.Equal(t, req.BodyHTML, output.BodyHTML)
	assert.Equal(t, req.Published, output.Published)
}
