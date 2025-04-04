package hubspot

import "fmt"

// Generated. Don't change.

// GET /crm/v3/objects/leads
func (c *Client) GetLeads(p *ListParams) (*PaginatedResponse, error) {
	path := "/crm/v3/objects/leads"
	return c.list(path, p)
}

// GET /crm/v3/objects/leads/{id}
func (c *Client) GetLead(id string, p *ReadParams) (*HsObject, error) {
	path := fmt.Sprintf("/crm/v3/objects/leads/%s", id)
	return c.get(path, p)
}

// POST /crm/v3/objects/leads
func (c *Client) CreateLead(body *CreateBody) (*HsObject, error) {
	path := "/crm/v3/objects/leads"
	return c.create(path, body)
}

// PATCH /crm/v3/objects/leads/{id}
func (c *Client) UpdateLead(id string, properties map[string]string) (*HsObject, error) {
	path := "/crm/v3/objects/leads"
	return c.update(path, properties)
}

// DELETE /crm/v3/objects/leads/{id}
func (c *Client) DeleteLead(id string) error {
	path := fmt.Sprintf("/crm/v3/objects/leads/%s", id)
	return c.delete(path)
}
