package hubspot

import "fmt"

// Generated. Don't change.

// GET /crm/v3/objects/line_items
func (c *Client) GetLineItems(p *ListParams) (*PaginatedResponse, error) {
	path := "/crm/v3/objects/line_items"
	return c.list(path, p)
}

// GET /crm/v3/objects/line_items/{id}
func (c *Client) GetLineItem(id string, p *ReadParams) (*HsObject, error) {
	path := fmt.Sprintf("/crm/v3/objects/line_items/%s", id)
	return c.get(path, p)
}

// POST /crm/v3/objects/line_items
func (c *Client) CreateLineItem(body *CreateBody) (*HsObject, error) {
	path := "/crm/v3/objects/line_items"
	return c.create(path, body)
}

// PATCH /crm/v3/objects/line_items/{id}
func (c *Client) UpdateLineItem(id string, properties map[string]string) (*HsObject, error) {
	path := "/crm/v3/objects/line_items"
	return c.update(path, properties)
}

// DELETE /crm/v3/objects/line_items/{id}
func (c *Client) DeleteLineItem(id string) error {
	path := fmt.Sprintf("/crm/v3/objects/line_items/%s", id)
	return c.delete(path)
}
