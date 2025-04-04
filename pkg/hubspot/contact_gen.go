package hubspot

import "fmt"

// Generated. Don't change.

// GET /crm/v3/objects/contacts
func (c *Client) GetContacts(p *ListParams) (*PaginatedResponse, error) {
	path := "/crm/v3/objects/contacts"
	return c.list(path, p)
}

// GET /crm/v3/objects/contacts/{id}
func (c *Client) GetContact(id string, p *ReadParams) (*HsObject, error) {
	path := fmt.Sprintf("/crm/v3/objects/contacts/%s", id)
	return c.get(path, p)
}

// POST /crm/v3/objects/contacts
func (c *Client) CreateContact(body *CreateBody) (*HsObject, error) {
	path := "/crm/v3/objects/contacts"
	return c.create(path, body)
}

// PATCH /crm/v3/objects/contacts/{id}
func (c *Client) UpdateContact(id string, properties map[string]string) (*HsObject, error) {
	path := "/crm/v3/objects/contacts"
	return c.update(path, properties)
}

// DELETE /crm/v3/objects/contacts/{id}
func (c *Client) DeleteContact(id string) error {
	path := fmt.Sprintf("/crm/v3/objects/contacts/%s", id)
	return c.delete(path)
}
