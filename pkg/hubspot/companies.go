package hubspot

import (
	"fmt"
)

// GET /crm/v3/objects/companies
func (c *Client) GetCompanies(p *ListParams) (*PaginatedResponse, error) {
	path := "/crm/v3/objects/companies"

	return c.list(path, p)
}

// GET /crm/v3/objects/companies/{id}
func (c *Client) GetCompany(id string, p *ReadParams) (*HsObject, error) {
	path := fmt.Sprintf("/crm/v3/objects/companies/%s", id)

	return c.get(path, p)
}

// POST /crm/v3/objects/companies
func (c *Client) CreateCompany(body *CreateBody) (*HsObject, error) {
	path := "/crm/v3/objects/companies"

	return c.create(path, body)
}

// PATCH /crm/v3/objects/companies/{id}
func (c *Client) UpdateCompany(id string, properties map[string]string) (*HsObject, error) {
	path := fmt.Sprintf("/crm/v3/objects/companies/%s", id)

	return c.update(path, properties)
}

// DELETE /crm/v3/objects/companies/{id}
func (c *Client) DeleteCompany(id string) error {
	path := fmt.Sprintf("/crm/v3/objects/companies/%s", id)

	return c.delete(path)
}
