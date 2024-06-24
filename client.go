package glide

import (
	"context"
	"net/http"
	"net/url"

	"github.com/einstack/glide-go/config"
	"github.com/einstack/glide-go/lang"
)

// Client is an Einstack Glide client.
type Client struct {
	config *config.Config
	Lang   lang.Language
}

// ClientOption is a functional option type for Client.
// Also see NewClient.
type ClientOption func(*Client) error

// NewClient instantiates a new Client.
func NewClient(options ...ClientOption) (*Client, error) {
	client := &Client{config: config.NewConfig()}
	client.Lang = lang.NewLanguage(client.config)

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// NewClientFromConfig instantiates a new Client.
func NewClientFromConfig(config *config.Config) (*Client, error) {
	client := &Client{config: config}
	client.Lang = lang.NewLanguage(client.config)
	return client, nil
}

// WithApiKey attaches the api key.
// Use environment variable 'GLIDE_API_KEY' to override.
func WithApiKey(apiKey string) ClientOption {
	return func(client *Client) error {
		client.config.ApiKey = apiKey
		return nil
	}
}

// WithUserAgent replaces the 'User-Agent' header.
// Default value: 'Glide/0.1.0 (Go; Ver. 1.22.2)'.
// Use environment variable 'GLIDE_USER_AGENT' to override.
func WithUserAgent(userAgent string) ClientOption {
	return func(client *Client) error {
		client.config.UserAgent = userAgent
		return nil
	}
}

// WithRawBaseURL parses and replaces the base URL.
// Default value: 'http://127.0.0.1:9099/'.
// Use environment variable 'GLIDE_BASE_URL' to override.
func WithRawBaseURL(rawBaseURL string) ClientOption {
	return func(client *Client) error {
		baseURL, err := url.Parse(rawBaseURL)
		if err != nil {
			return err
		}

		client.config.BaseURL = baseURL
		return nil
	}
}

// WithBaseURL replaces the base URL.
// Also see WithRawBaseURL.
func WithBaseURL(baseURL url.URL) ClientOption {
	return func(client *Client) error {
		client.config.BaseURL = &baseURL
		return nil
	}
}

// WithHttpClient replaces the 'HTTP' client.
// Default value: 'http.DefaultClient'.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(client *Client) error {
		client.config.HttpClient = httpClient
		return nil
	}
}

// ApiKey returns the provided API key, empty string otherwise.
func (c *Client) ApiKey() string {
	return c.config.ApiKey
}

// UserAgent returns the used 'User-Agent' header value.
func (c *Client) UserAgent() string {
	return c.config.UserAgent
}

// BaseURL returns the used 'base url.URL'.
func (c *Client) BaseURL() url.URL {
	return *c.config.BaseURL
}

// HttpClient returns the underlying http.Client.
func (c *Client) HttpClient() *http.Client {
	return c.config.HttpClient
}

// Health returns true if the service is healthy.
func (c *Client) Health(ctx context.Context) (*bool, error) {
	type Health struct {
		Healthy bool `json:"healthy"`
	}

	req, err := c.config.Build(ctx, http.MethodGet, "/v1/health/", nil)
	if err != nil {
		return nil, err
	}

	var resp *Health
	if _, err := c.config.Send(req, resp); err != nil {
		return nil, err
	}

	return &resp.Healthy, nil
}
