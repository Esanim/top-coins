// Package externalcoinclient External Coin Client API
package externalcoinclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Config the client config structure
type Config struct {
	apiKey string
}

// Client holds information necessary to make a request to your API
type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

// Option is a functional option for configuring the API client
type Option func(*Client) error

// BaseURL allows overriding of API client baseURL for testing
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// HTTPClient allows overriding of API http client for testing
func HTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		c.httpClient = client
		return nil
	}
}

// parseOptions parses the supplied options functions and returns a configured
// *Client instance
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// doReq HTTP client
func (c *Client) doReq(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

// makeReq HTTP request helper
func (c *Client) makeReq(url string, headers map[string]interface{}) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}
	resp, err := c.doReq(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// New creates a new API client
func New(cfg Config, opts ...Option) (*Client, error) {

	if cfg.apiKey == "" {
		cfg.apiKey = os.Getenv(apiHeaderKeyName)
	}

	if cfg.apiKey == "" {
		log.Fatal("API Key is required")
	}
	client := &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(5 * time.Second),
		},
		apiKey: cfg.apiKey,
	}

	if err := client.parseOptions(opts...); err != nil {
		return nil, err
	}

	return client, nil
}
