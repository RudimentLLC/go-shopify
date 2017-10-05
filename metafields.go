package shopify

type CreateMetafieldResponse struct {
	Metafield *Metafield `json:"metafield"`
}

type Metafield struct {
	ID        int         `json:"id,omitempty"`
	Namespace string      `json:"namespace"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	ValueType string      `json:"value_type"`
}
