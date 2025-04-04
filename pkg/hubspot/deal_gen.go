package hubspot

import "fmt"

// Generated. Don't change.

// GET /crm/v3/objects/deals
func (c *Client) GetDeals(p *ListParams) (*PaginatedResponse, error) {
	path := "/crm/v3/objects/deals"
	return c.list(path, p)
}

// GET /crm/v3/objects/deals/{id}
func (c *Client) GetDeal(id string, p *ReadParams) (*HsObject, error) {
	path := fmt.Sprintf("/crm/v3/objects/deals/%s", id)
	return c.get(path, p)
}

// POST /crm/v3/objects/deals
func (c *Client) CreateDeal(body *CreateBody) (*HsObject, error) {
	path := "/crm/v3/objects/deals"
	return c.create(path, body)
}

// PATCH /crm/v3/objects/deals/{id}
func (c *Client) UpdateDeal(id string, properties map[string]string) (*HsObject, error) {
	path := "/crm/v3/objects/deals"
	return c.update(path, properties)
}

// DELETE /crm/v3/objects/deals/{id}
func (c *Client) DeleteDeal(id string) error {
	path := fmt.Sprintf("/crm/v3/objects/deals/%s", id)
	return c.delete(path)
}
