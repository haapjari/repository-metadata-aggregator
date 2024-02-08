// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for GetApiV1RepositoriesSearchParamsOrder.
const (
	Asc  GetApiV1RepositoriesSearchParamsOrder = "asc"
	Desc GetApiV1RepositoriesSearchParamsOrder = "desc"
)

// Repository defines model for Repository.
type Repository struct {
	ClosedIssuesCount *int       `json:"closed_issues_count,omitempty"`
	ClosedPullsCount  *int       `json:"closed_pulls_count,omitempty"`
	CommitsCount      *int       `json:"commits_count,omitempty"`
	ContributorsCount *int       `json:"contributors_count,omitempty"`
	CreationDate      *time.Time `json:"creation_date,omitempty"`
	Description       *string    `json:"description,omitempty"`
	EventsCount       *int       `json:"events_count,omitempty"`
	ForksCount        *int       `json:"forks_count,omitempty"`
	FullName          string     `json:"full_name"`
	HtmlUrl           string     `json:"html_url"`
	Id                int        `json:"id"`
	Language          string     `json:"language"`
	LatestRelease     *struct {
		PublishedAt *time.Time `json:"published_at,omitempty"`
		TagName     *string    `json:"tag_name,omitempty"`
	} `json:"latest_release,omitempty"`
	LibraryLoc         *int   `json:"library_loc,omitempty"`
	Name               string `json:"name"`
	OpenIssuesCount    *int   `json:"open_issues_count,omitempty"`
	OpenPullsCount     *int   `json:"open_pulls_count,omitempty"`
	SelfWrittenLoc     *int   `json:"self_written_loc,omitempty"`
	StargazersCount    int    `json:"stargazers_count"`
	SubscribersCount   *int   `json:"subscribers_count,omitempty"`
	TotalReleasesCount *int   `json:"total_releases_count,omitempty"`
	WatchersCount      *int   `json:"watchers_count,omitempty"`
}

// GetApiV1RepositoriesSearchParams defines parameters for GetApiV1RepositoriesSearch.
type GetApiV1RepositoriesSearchParams struct {
	// FirstCreationDate The first creation date of the range, in YYYY-MM-DD format.
	FirstCreationDate string `form:"firstCreationDate" json:"firstCreationDate"`

	// LastCreationDate The last creation date of the range, in YYYY-MM-DD format.
	LastCreationDate string `form:"lastCreationDate" json:"lastCreationDate"`

	// Language The primary programming language of the repositories to search for.
	Language string `form:"language" json:"language"`

	// Stars The number of stars a repository must have, e.g., ">100" for more than 100 stars.
	Stars string `form:"stars" json:"stars"`

	// Order The order of the results, either ascending (asc) or descending (desc). Defaults to descending.
	Order *GetApiV1RepositoriesSearchParamsOrder `form:"order,omitempty" json:"order,omitempty"`
}

// GetApiV1RepositoriesSearchParamsOrder defines parameters for GetApiV1RepositoriesSearch.
type GetApiV1RepositoriesSearchParamsOrder string

// GetApiV1RepositoriesSearchFirstCreationDateParams defines parameters for GetApiV1RepositoriesSearchFirstCreationDate.
type GetApiV1RepositoriesSearchFirstCreationDateParams struct {
	// Language The primary programming language of the repositories to search for.
	Language string `form:"language" json:"language"`

	// Stars Minimum number of stars a repository must have. e.g., ">100" for more than 100 stars.
	Stars string `form:"stars" json:"stars"`
}

// GetApiV1RepositoriesSearchLastCreationDateParams defines parameters for GetApiV1RepositoriesSearchLastCreationDate.
type GetApiV1RepositoriesSearchLastCreationDateParams struct {
	// Language The primary programming language of the repositories to search for.
	Language string `form:"language" json:"language"`

	// Stars Minimum number of stars a repository must have. e.g., ">100" for more than 100 stars.
	Stars string `form:"stars" json:"stars"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetApiV1RepositoriesSearch request
	GetApiV1RepositoriesSearch(ctx context.Context, params *GetApiV1RepositoriesSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1RepositoriesSearchFirstCreationDate request
	GetApiV1RepositoriesSearchFirstCreationDate(ctx context.Context, params *GetApiV1RepositoriesSearchFirstCreationDateParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1RepositoriesSearchLastCreationDate request
	GetApiV1RepositoriesSearchLastCreationDate(ctx context.Context, params *GetApiV1RepositoriesSearchLastCreationDateParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetApiV1RepositoriesSearch(ctx context.Context, params *GetApiV1RepositoriesSearchParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1RepositoriesSearchRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1RepositoriesSearchFirstCreationDate(ctx context.Context, params *GetApiV1RepositoriesSearchFirstCreationDateParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1RepositoriesSearchFirstCreationDateRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1RepositoriesSearchLastCreationDate(ctx context.Context, params *GetApiV1RepositoriesSearchLastCreationDateParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1RepositoriesSearchLastCreationDateRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetApiV1RepositoriesSearchRequest generates requests for GetApiV1RepositoriesSearch
func NewGetApiV1RepositoriesSearchRequest(server string, params *GetApiV1RepositoriesSearchParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/repositories/search")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "firstCreationDate", runtime.ParamLocationQuery, params.FirstCreationDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "lastCreationDate", runtime.ParamLocationQuery, params.LastCreationDate); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "language", runtime.ParamLocationQuery, params.Language); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "stars", runtime.ParamLocationQuery, params.Stars); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if params.Order != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "order", runtime.ParamLocationQuery, *params.Order); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1RepositoriesSearchFirstCreationDateRequest generates requests for GetApiV1RepositoriesSearchFirstCreationDate
func NewGetApiV1RepositoriesSearchFirstCreationDateRequest(server string, params *GetApiV1RepositoriesSearchFirstCreationDateParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/repositories/search/firstCreationDate")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "language", runtime.ParamLocationQuery, params.Language); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "stars", runtime.ParamLocationQuery, params.Stars); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1RepositoriesSearchLastCreationDateRequest generates requests for GetApiV1RepositoriesSearchLastCreationDate
func NewGetApiV1RepositoriesSearchLastCreationDateRequest(server string, params *GetApiV1RepositoriesSearchLastCreationDateParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/repositories/search/lastCreationDate")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "language", runtime.ParamLocationQuery, params.Language); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "stars", runtime.ParamLocationQuery, params.Stars); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetApiV1RepositoriesSearchWithResponse request
	GetApiV1RepositoriesSearchWithResponse(ctx context.Context, params *GetApiV1RepositoriesSearchParams, reqEditors ...RequestEditorFn) (*GetApiV1RepositoriesSearchResponse, error)

	// GetApiV1RepositoriesSearchFirstCreationDateWithResponse request
	GetApiV1RepositoriesSearchFirstCreationDateWithResponse(ctx context.Context, params *GetApiV1RepositoriesSearchFirstCreationDateParams, reqEditors ...RequestEditorFn) (*GetApiV1RepositoriesSearchFirstCreationDateResponse, error)

	// GetApiV1RepositoriesSearchLastCreationDateWithResponse request
	GetApiV1RepositoriesSearchLastCreationDateWithResponse(ctx context.Context, params *GetApiV1RepositoriesSearchLastCreationDateParams, reqEditors ...RequestEditorFn) (*GetApiV1RepositoriesSearchLastCreationDateResponse, error)
}

type GetApiV1RepositoriesSearchResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Items []Repository `json:"items"`

		// TotalCount The total number of repositories found.
		TotalCount int `json:"total_count"`
	}
	JSON400 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON403 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON500 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON503 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetApiV1RepositoriesSearchResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1RepositoriesSearchResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1RepositoriesSearchFirstCreationDateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// FirstCreationDate Earliest creation date of from set of the repositories returned from the GitHub Search API.
		FirstCreationDate *string `json:"firstCreationDate,omitempty"`
	}
	JSON400 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON403 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON500 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON503 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetApiV1RepositoriesSearchFirstCreationDateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1RepositoriesSearchFirstCreationDateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1RepositoriesSearchLastCreationDateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// LastCreationDate Last creation date of from set of the repositories returned from the GitHub Search API.
		LastCreationDate *string `json:"lastCreationDate,omitempty"`
	}
	JSON400 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON403 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON500 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
	JSON503 *struct {
		// Error Error Message.
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetApiV1RepositoriesSearchLastCreationDateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1RepositoriesSearchLastCreationDateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetApiV1RepositoriesSearchWithResponse request returning *GetApiV1RepositoriesSearchResponse
func (c *ClientWithResponses) GetApiV1RepositoriesSearchWithResponse(ctx context.Context, params *GetApiV1RepositoriesSearchParams, reqEditors ...RequestEditorFn) (*GetApiV1RepositoriesSearchResponse, error) {
	rsp, err := c.GetApiV1RepositoriesSearch(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1RepositoriesSearchResponse(rsp)
}

// GetApiV1RepositoriesSearchFirstCreationDateWithResponse request returning *GetApiV1RepositoriesSearchFirstCreationDateResponse
func (c *ClientWithResponses) GetApiV1RepositoriesSearchFirstCreationDateWithResponse(ctx context.Context, params *GetApiV1RepositoriesSearchFirstCreationDateParams, reqEditors ...RequestEditorFn) (*GetApiV1RepositoriesSearchFirstCreationDateResponse, error) {
	rsp, err := c.GetApiV1RepositoriesSearchFirstCreationDate(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1RepositoriesSearchFirstCreationDateResponse(rsp)
}

// GetApiV1RepositoriesSearchLastCreationDateWithResponse request returning *GetApiV1RepositoriesSearchLastCreationDateResponse
func (c *ClientWithResponses) GetApiV1RepositoriesSearchLastCreationDateWithResponse(ctx context.Context, params *GetApiV1RepositoriesSearchLastCreationDateParams, reqEditors ...RequestEditorFn) (*GetApiV1RepositoriesSearchLastCreationDateResponse, error) {
	rsp, err := c.GetApiV1RepositoriesSearchLastCreationDate(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1RepositoriesSearchLastCreationDateResponse(rsp)
}

// ParseGetApiV1RepositoriesSearchResponse parses an HTTP response from a GetApiV1RepositoriesSearchWithResponse call
func ParseGetApiV1RepositoriesSearchResponse(rsp *http.Response) (*GetApiV1RepositoriesSearchResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1RepositoriesSearchResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Items []Repository `json:"items"`

			// TotalCount The total number of repositories found.
			TotalCount int `json:"total_count"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetApiV1RepositoriesSearchFirstCreationDateResponse parses an HTTP response from a GetApiV1RepositoriesSearchFirstCreationDateWithResponse call
func ParseGetApiV1RepositoriesSearchFirstCreationDateResponse(rsp *http.Response) (*GetApiV1RepositoriesSearchFirstCreationDateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1RepositoriesSearchFirstCreationDateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// FirstCreationDate Earliest creation date of from set of the repositories returned from the GitHub Search API.
			FirstCreationDate *string `json:"firstCreationDate,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetApiV1RepositoriesSearchLastCreationDateResponse parses an HTTP response from a GetApiV1RepositoriesSearchLastCreationDateWithResponse call
func ParseGetApiV1RepositoriesSearchLastCreationDateResponse(rsp *http.Response) (*GetApiV1RepositoriesSearchLastCreationDateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1RepositoriesSearchLastCreationDateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// LastCreationDate Last creation date of from set of the repositories returned from the GitHub Search API.
			LastCreationDate *string `json:"lastCreationDate,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest struct {
			// Error Error Message.
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Abstraction for GitHub Search API, accepts Search API Queries with a Date Range.
	// (GET /api/v1/repositories/search)
	GetApiV1RepositoriesSearch(c *gin.Context, params GetApiV1RepositoriesSearchParams)
	// Returns first creation date from the set of repositories.
	// (GET /api/v1/repositories/search/firstCreationDate)
	GetApiV1RepositoriesSearchFirstCreationDate(c *gin.Context, params GetApiV1RepositoriesSearchFirstCreationDateParams)
	// Returns first creation date from the set of repositories.
	// (GET /api/v1/repositories/search/lastCreationDate)
	GetApiV1RepositoriesSearchLastCreationDate(c *gin.Context, params GetApiV1RepositoriesSearchLastCreationDateParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetApiV1RepositoriesSearch operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RepositoriesSearch(c *gin.Context) {

	var err error

	c.Set(ApiKeyAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiV1RepositoriesSearchParams

	// ------------- Required query parameter "firstCreationDate" -------------

	if paramValue := c.Query("firstCreationDate"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument firstCreationDate is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "firstCreationDate", c.Request.URL.Query(), &params.FirstCreationDate)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter firstCreationDate: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "lastCreationDate" -------------

	if paramValue := c.Query("lastCreationDate"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument lastCreationDate is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "lastCreationDate", c.Request.URL.Query(), &params.LastCreationDate)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter lastCreationDate: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "language" -------------

	if paramValue := c.Query("language"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument language is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "stars" -------------

	if paramValue := c.Query("stars"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument stars is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "stars", c.Request.URL.Query(), &params.Stars)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter stars: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "order" -------------

	err = runtime.BindQueryParameter("form", true, false, "order", c.Request.URL.Query(), &params.Order)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter order: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1RepositoriesSearch(c, params)
}

// GetApiV1RepositoriesSearchFirstCreationDate operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RepositoriesSearchFirstCreationDate(c *gin.Context) {

	var err error

	c.Set(ApiKeyAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiV1RepositoriesSearchFirstCreationDateParams

	// ------------- Required query parameter "language" -------------

	if paramValue := c.Query("language"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument language is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "stars" -------------

	if paramValue := c.Query("stars"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument stars is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "stars", c.Request.URL.Query(), &params.Stars)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter stars: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1RepositoriesSearchFirstCreationDate(c, params)
}

// GetApiV1RepositoriesSearchLastCreationDate operation middleware
func (siw *ServerInterfaceWrapper) GetApiV1RepositoriesSearchLastCreationDate(c *gin.Context) {

	var err error

	c.Set(ApiKeyAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiV1RepositoriesSearchLastCreationDateParams

	// ------------- Required query parameter "language" -------------

	if paramValue := c.Query("language"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument language is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "language", c.Request.URL.Query(), &params.Language)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter language: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "stars" -------------

	if paramValue := c.Query("stars"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument stars is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "stars", c.Request.URL.Query(), &params.Stars)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter stars: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetApiV1RepositoriesSearchLastCreationDate(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/v1/repositories/search", wrapper.GetApiV1RepositoriesSearch)
	router.GET(options.BaseURL+"/api/v1/repositories/search/firstCreationDate", wrapper.GetApiV1RepositoriesSearchFirstCreationDate)
	router.GET(options.BaseURL+"/api/v1/repositories/search/lastCreationDate", wrapper.GetApiV1RepositoriesSearchLastCreationDate)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xXXW/bxhL9KwPe+3AD0JIc37zozfe6+UBjNJXTAkZiCCNyJG5C7jKzs06VQP+9mKVo",
	"UdGH3ToG3FZv4n7NmbOzc46+JpmramfJik+GXxOfFVRh/Dmi2nkjjuf6VbOricVQnMtK5ykfG+8D+XHm",
	"ghUdlnlNyTAxVmhGnCzSdmEdynL/OldVRm5ZYoXNJIjj/euYUIyz4xyFdMnUcYWSDBMdOBJTUZK2+7yw",
	"sTPdlpPP2NS6s3Puap6ulaN9gaeOP+5fEMpybLGirQEKqcpx4HLrpMm3H1minQWcbT+xRCEvY6aS0NPm",
	"LdZhUhpfUD5GuTtRgrNdSSxulrvJB8okgjATRp6PS5dtT2EnIa4me4cSi8tuLTBP5XT8mY0I2d1YvCDP",
	"8AvtrzAfJlork1uWiRMsW/b3rvyMkhV7T1ukCdOnYJjyZPhOy2FJXLeqOjW0JZVOsVxt3FNkKAtsZH6h",
	"LaApkNPa/Ejz0yCFfq29kOSn+ANLeGHkZZjAG2Ifv0+zjLyHt+4jWRAHGKQgKyZDIdAkyIuHqWMwVt+q",
	"pxxY50qjLaCXpInRAAVhTtzmOUwUhmPzJT7uVWFixJgsNAVjpy6yZ6TUuRcleg+nb14laXJN7Bvkx71B",
	"b9BWDtYmGSYncShNapQiZt7H2vSvj/vcNkFDvu8JOYtUzEg2GTmdeGHM9CumtyTmIu5SFClgllEtvjMG",
	"PwfSw+GzkQIQzpSJEdoZ9aCluJyn+1nO0MKEoEavZBrrTU6wxhe8jGzqfZD1gSkFKVDaY0ca9XXkH3JH",
	"HqwTwLou5z0YkQS2HqaGvcDxYDCALivNOZVWMEhB8CkQz/UStdPE2K9yvQqS09r8ejzqbG1YiLQzViTE",
	"Phm++5bWtwUtY7edHbQ/gZvGeKxcpWAsXF5eXh6dnx+dnUHTyRQF/YZVHYvh6eD45Gjw7Ghw3FZYhLoq",
	"sBjk/8sYeg9J99EJB0qX+ri19W3DXeLDw9YY3xd1zaZCnkPNbsZYVcbOoO0dN/jXSsBB8zY0g3X4L9xO",
	"2MtmdG+4NlQTYgWmPc8DrsDNoQpeoMBrSoF6s14K75P3YTA4oePB4H0S32nlmLSKrRZ3c8Z6DjcbdqQS",
	"t9w/D8d5k0bDrw+l+BTISEEM6DOyud7Ef9BnT8Ax6AntmP5+0oMzmqJu0xtZTa9no+M7EokIki7wvDlw",
	"tY1sqFSBMH7FwasNk7C4UjJ87axvhOTpYBBto7NCjcBpc1FJMM72P/jGc63CrvsUI1St//g30zQZJv/q",
	"r9xrf2ld+x3fuvIjyIzzlSjf6OzmLcQFnZpaK/SpCzbvrcRnlz53o6RL2FtE91vbmVyE2NanoYTRkj8F",
	"/d970UfMjjeT/UGH4Zy8xxn1Np3e4g54/4c5jBpJb3CePE6czx1PTJ6TVZTPHiubr6wQq7hfEF8TQzyz",
	"QfxIeVWgJiP4xeI1mhInJa15ySjnXRf57kp7gw+VCswDeKYYfY9562+K/C47NyJhQ9fR4dC6iPu2SxNy",
	"achLV3EUlLGAkLmypCa1xiMRSdzla8rM1FDeqmbGRogN9uBtYTyQzWtnrNykHts0rFxS096nxlJz3vIU",
	"Z3Oj4XwKHE2bSsNu8I2pQps3NmWPqfPUCf5HvN3zLZbqVrN3f/vxcH7j3FhTheqOnqP3mDzH99Xkre/o",
	"mwbTPo4NAzxlV4En2XqbTelS3qzS6Y2m8Ge71UFdD+r6T1LX9b/u68/w5nUt32H3Dd4uoxt/Og8q+nAq",
	"+hoPIvq3FNFtr2idqtd40M+Dfh7086+nn4vF7wEAAP//hQMmLlwdAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
