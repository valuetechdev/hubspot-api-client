package hubspot

import "fmt"

// Generated. Don't change.

// GET /crm/v3/objects/products
func (c *Client) GetProducts(p *ListParams) (*PaginatedResponse, error) {
	path := "/crm/v3/objects/products"
	return c.list(path, p)
}

// GET /crm/v3/objects/products/{id}
func (c *Client) GetProduct(id string, p *ReadParams) (*HsObject, error) {
	path := fmt.Sprintf("/crm/v3/objects/products/%s", id)
	return c.get(path, p)
}

// POST /crm/v3/objects/products
func (c *Client) CreateProduct(body *CreateBody) (*HsObject, error) {
	path := "/crm/v3/objects/products"
	return c.create(path, body)
}

// PATCH /crm/v3/objects/products/{id}
func (c *Client) UpdateProduct(id string, properties map[string]string) (*HsObject, error) {
	path := "/crm/v3/objects/products"
	return c.update(path, properties)
}

// DELETE /crm/v3/objects/products/{id}
func (c *Client) DeleteProduct(id string) error {
	path := fmt.Sprintf("/crm/v3/objects/products/%s", id)
	return c.delete(path)
}
