package otxapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	DefaultBaseURL   = "https://otx.alienvault.com"
	DefaultUserAgent = "go-otx-api/0.2"

	SubscriptionsURLPath = "api/v1/pulses/subscribed"
	PulseDetailURLPath   = "api/v1/pulses/"
	UserURLPath          = "api/v1/user/"
)

// ClientOption is a function that configures a *Client.
type ClientOption func(*Client) error

// APIKey sets the API key used by a *Client for authorizing requests to the
// OpenThreatExchange HTTP API.
func APIKey(key string) ClientOption {
	return func(c *Client) error {
		if key == "" {
			return errors.New("empty api key")
		}
		c.apiKey = key
		return nil
	}
}

// APIKeyFromEnv sets the API key used by the *Client from one of the named
// environment variables. If no environment variable names are specified,
// the X_OTX_API_KEY, and ALIENVAULT_OTXAPI_KEY environment variables will be
// checked.
//
// The returned ClientOption will return a non-nil error if it cannot find a
// non-empty environment variable.
func APIKeyFromEnv(names ...string) ClientOption {
	return func(c *Client) error {
		if len(names) == 0 {
			names = []string{
				"X_OTX_API_KEY",
				"ALIENVAULT_OTXAPI_KEY",
			}
		}
		for _, name := range names {
			if v := os.Getenv(name); v != "" {
				c.apiKey = v
				return nil
			}
		}
		return errors.New("api key not set in environment")
	}
}

// HTTPClient sets a *Client's internal *http.Client. If hc is nil, then
// http.DefaultClient is used.
func HTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) error {
		if hc == nil {
			c.client = http.DefaultClient
			return nil
		}
		c.client = hc
		return nil
	}
}

// UserAgent sets the User-Agent HTTP header for requests made by a *Client.
func UserAgent(ua string) ClientOption {
	return func(c *Client) error {
		c.userAgent = ua
		return nil
	}
}

// NewClient returns a new OpenThreatExchange HTTP API client.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{
		client:    http.DefaultClient,
		userAgent: DefaultUserAgent,
		baseURL:   DefaultBaseURL,
	}

	// Apply any client options.
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, fmt.Errorf("applying client option: %v", err)
		}
	}

	if c.apiKey == "" {
		return nil, ErrNoAPIKey
	}

	c.User = &OTXUserDetailService{client: c}
	c.Pulses = &OTXPulseDetailService{client: c}

	return c, nil
}

// A Client manages communication with the OTX API.
type Client struct {
	User   *OTXUserDetailService
	Pulses *OTXPulseDetailService

	baseURL   string
	userAgent string
	client    *http.Client
	apiKey    string
}

func (c Client) String() string {
	return fmt.Sprintf("%s; host=%q", c.userAgent, c.baseURL)
}

var ErrNoAPIKey = errors.New("api key not set in client")

// newRequest returns a new *http.Request, with the X-OTX-API-KEY and
// User-Agent headers set.
func (c *Client) newRequest(method, urlPath string, body io.Reader) (*http.Request, error) {
	// Sanitize the URL path, by removing any leading "/" characters.
	start := 0
	for i := start; urlPath[i] == '/'; i++ {
		start = i
	}
	urlPath = urlPath[start:]

	req, err := http.NewRequest(method, c.baseURL+"/"+urlPath, body)
	if err != nil {
		return nil, fmt.Errorf("new request: %v", err)
	}

	// Set the user agent header.
	req.Header.Set("User-Agent", c.userAgent)

	if c.apiKey == "" {
		return nil, ErrNoAPIKey
	}
	req.Header.Set("X-OTX-API-KEY", c.apiKey)

	return req, nil
}

// do executes the *http.Request.
//
// On a successful request where the server responds with a 200-series HTTP
// status code, this method will attempt to unmarshal the JSON-encoded
// response body into v.
//
// If the server responds with an unexpected status code, this method will attempt
// to unmarshal the JSON response body into an *APIError, and return it.
func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("cannot decode response body: %v", err)
		}
		return nil
	}

	// If we manage to fall through to here, we have an unexpected status
	// code. Attempt to unmarshal the response to an *APIError.
	var apiErr APIError
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return fmt.Errorf("json unmarshal failed: %v", err)
	}
	apiErr.statusCode = resp.StatusCode

	return &apiErr
}

// IsAPIError returns a bool indicating whether or not the given error is a
// *APIError.
func IsAPIError(err error) bool {
	_, ok := err.(*APIError)
	return ok
}

type APIError struct {
	Message string `json:"detail"`

	statusCode int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s (status %d)", e.Message, e.statusCode)
}
