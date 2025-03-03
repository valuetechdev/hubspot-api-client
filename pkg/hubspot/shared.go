package hubspot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func toListResult(res *http.Response) (*PaginatedResponse, error) {

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var dst PaginatedResponse
	err = json.Unmarshal(body, &dst)
	if err != nil {
		return nil, err
	}

	return &dst, err
}

func toHsObject(res *http.Response) (*HsObject, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var dst HsObject
	err = json.Unmarshal(body, &dst)
	if err != nil {
		return nil, err
	}

	return &dst, err
}

func (c *Client) doList(method, path string, params map[string][]string) (*PaginatedResponse, error) {
	res, err := c.do(method, path, nil, params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if !c.ok(res) {
		return nil, c.errorResponse(res.StatusCode, nil, res.Status)
	}

	list, err := toListResult(res)
	if err != nil {
		return nil, c.errorResponse(500, nil, "Could not unmarshal result")
	}

	return list, nil

}

// send a request. if the response is in < 400 and didn't error, return HsObject, else return error
func (c *Client) hsDo(method, path string, body []byte, params map[string][]string) (*HsObject, error) {
	res, err := c.do(method, path, body, params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if !c.ok(res) {
		b, _ := io.ReadAll(res.Body)
		return nil, c.errorResponse(res.StatusCode, b, res.Status)
	}

	result, err := toHsObject(res)
	if err != nil {
		return nil, c.errorResponse(500, nil, "Could not unmarshal result")
	}

	return result, nil

}

func (c *Client) list(path string, p *ListParams) (*PaginatedResponse, error) {
	method := http.MethodGet

	return c.doList(method, path, p.ToMap())
}

func (c *Client) get(path string, p *ReadParams) (*HsObject, error) {
	method := http.MethodGet

	return c.hsDo(method, path, nil, p.ToMap())
}

func (c *Client) create(path string, body *CreateBody) (*HsObject, error) {
	method := http.MethodPost

	if body == nil {
		return nil, fmt.Errorf("body can't be nil")
	}

	b, _ := json.Marshal(body)

	return c.hsDo(method, path, b, nil)
}

func (c *Client) update(path string, properties map[string]string) (*HsObject, error) {
	method := http.MethodPatch

	payload := map[string]any{
		"properties": properties,
	}

	body, _ := json.Marshal(payload)

	return c.hsDo(method, path, body, nil)
}

func (c *Client) delete(path string) error {
	method := http.MethodDelete
	res, err := c.do(method, path, nil, nil)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", res.Status)
	}

	return nil
}
