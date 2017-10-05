package shopify

import "fmt"

type CreatePageRequest struct {
	Page Page `json:"page"`
}

type CreatePageResponse struct {
	Page *Page `json:"page"`
}

type Page struct {
	ID             int         `json:"id,omitempty"`
	Title          string      `json:"title,omitempty"`
	Body           string      `json:"body,omitempty"`
	BodyHTML       string      `json:"body_html,omitempty"`
	Published      bool        `json:"published,omitempty"`
	ShopID         int         `json:"shop_id,omitempty"`
	Handle         string      `json:"handle,omitempty"`
	Author         string      `json:"author,omitempty"`
	CreatedAt      string      `json:"created_at,omitempty"`
	UpdatedAt      string      `json:"updated_at,omitempty"`
	PublishedAt    string      `json:"published_at,omitempty"`
	TemplateSuffix string      `json:"template_suffix,omitempty"`
	Metafields     []Metafield `json:"metafields,omitempty"`
}

func (c *Client) CreatePage(p Page) (*Page, error) {
	req := CreatePageRequest{Page: p}

	var resp CreatePageResponse
	if err := c.Client.Post("/admin/pages.json", req, &resp); err != nil {
		return nil, err
	}

	return resp.Page, nil
}

func (c *Client) DeletePage(pageID int) error {
	path := fmt.Sprintf("/admin/pages/%d.json", pageID)
	if err := c.Client.Delete(path, nil, nil); err != nil {
		return err
	}

	return nil
}
