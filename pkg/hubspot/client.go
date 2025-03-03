package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/valuetechdev/hubspot-api-client/pkg/httpclient"
)

const defaultHost = "https://api.hubapi.com"

type Client struct {
	// most likely https://api.hubapi.com
	host string

	httpClient *http.Client

	privateKey string
}

type ListParams struct {
	Limit        int
	After        string
	Properties   []string
	Associations []string
}

func (p *ListParams) ToMap() map[string][]string {
	m := map[string][]string{}

	if p.Limit != 0 {
		m["limit"] = []string{fmt.Sprint(p.Limit)}
	}
	if p.After != "" {
		m["after"] = []string{fmt.Sprint(p.After)}
	}
	if p.Properties != nil {
		m["properties"] = p.Properties
	}
	if p.Associations != nil {
		m["associations"] = p.Associations
	}

	return m
}

type ReadParams struct {
	Properties   []string
	Associations []string
}

func (p *ReadParams) ToMap() map[string][]string {
	m := map[string][]string{}

	if p.Properties != nil {
		m["properties"] = p.Properties
	}
	if p.Associations != nil {
		m["associations"] = p.Associations
	}

	return m
}

type PaginatedResponse struct {
	Paging  Paging     `json:"paging"`
	Results []HsObject `json:"results"`
}

func (pr *PaginatedResponse) String() string {
	j, _ := json.MarshalIndent(pr, "", "  ")
	return string(j)
}

type Paging struct {
	Next *Next `json:"next"`
	Prev *Prev `json:"prev"`
}

type Next struct {
	Link  string `json:"link"`
	After string `json:"after"`
}
type Prev struct {
	Link   string `json:"link"`
	Before string `json:"after"`
}

type HsObject struct {
	// record id
	Id string `json:"id"`

	Associations any `json:"associations"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Archived   bool      `json:"archived"`
	ArchivedAt time.Time `json:"archivedAt"`

	// some values must be parsed to be used correctly. the values can be:
	//	- int
	//	- float
	//	- string
	//	- bool
	//	- date
	Properties            map[string]string          `json:"properties"`
	PropertiesWithHistory map[string]PropWithHistory `json:"propertiesWithHistory"`
}

func (r *HsObject) String() string {
	j, _ := json.MarshalIndent(r, "", "  ")
	return string(j)
}

type PropWithHistory struct {
	SourceId        string `json:"sourceId"`
	SourceType      string `json:"sourceType"`
	SourceLabel     string `json:"sourceLabel"`
	UpdatedByUserId int    `json:"updatedByUserId"`
	Value           string `json:"value"`
	Timestamp       string `json:"timestamp"`
}

type Association struct {
	Paging Paging
	Result AssociationResult
}
type AssociationResult struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type ErrorResponse struct {
	Message string
	Status  int
	Body    []byte
}

type CreateBody struct {
	Properties map[string]string `json:"properties"`

	// TODO implement
	Associations map[string]any `json:"associations"`
}

func (er *ErrorResponse) Error() string {
	if er.Body != nil {
		return fmt.Sprintf("%s\n%s", er.Message, er.Body)
	}
	return fmt.Sprintf("%d: %s", er.Status, er.Message)
}

func New(privateKey string) *Client {
	httpClient := httpclient.WithRetry()

	httpClient.Timeout = time.Second * 10

	return &Client{
		host:       defaultHost,
		httpClient: httpClient,
		privateKey: privateKey,
	}
}

func (c *Client) auth(r *http.Request) {
	r.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.privateKey))
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s%s", c.host, path)
}

func (c *Client) ok(r *http.Response) bool {
	return r.StatusCode >= 200 && r.StatusCode < 400
}

func (c *Client) errorResponse(statusCode int, body []byte, message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  statusCode,
		Message: message,
		Body:    body,
	}
}

func toErrorResponse(err error) (e *ErrorResponse, ok bool) {
	e, ok = err.(*ErrorResponse)
	return e, ok
}

func (c *Client) do(method, path string, body []byte, params map[string][]string) (*http.Response, error) {
	r, err := http.NewRequest(method, c.url(path), bytes.NewReader(body))

	c.auth(r)
	r.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	if params != nil {
		q := r.URL.Query()

		for key, val := range params {
			for _, v := range val {
				q.Add(key, v)
			}
		}
		r.URL.RawQuery = q.Encode()
	}

	return c.httpClient.Do(r)
}
