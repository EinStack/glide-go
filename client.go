package glide

import (
	"context"
	"net/http"
	"net/url"
)

// Client is a minimal 'Glide' client.
type Client struct {
	cfg  *config
	Lang Language
}

type ClientOption func(*Client) error

// NewClient instantiates a new Client.
func NewClient(options ...ClientOption) (*Client, error) {
	options = append([]ClientOption{
		WithApiKey(envApiKey),
		WithUserAgent(envUserAgent),
		WithBaseURL(envBaseUrl),
		WithHttpClient(http.DefaultClient),
	}, options...)

	client := &Client{cfg: &config{}}
	client.Lang = &languageSvc{client.cfg}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// WithApiKey attaches the api key.
// Use environment variable 'GLIDE_API_KEY' to override.
func WithApiKey(apiKey string) ClientOption {
	return func(client *Client) error {
		client.cfg.apiKey = apiKey
		return nil
	}
}

// WithUserAgent replaces the 'User-Agent' header.
// Default value: 'Glide/0.1.0 (Go; Ver. 1.22.2)'.
// Use environment variable 'GLIDE_USER_AGENT' to override.
func WithUserAgent(userAgent string) ClientOption {
	return func(client *Client) error {
		client.cfg.userAgent = userAgent
		return nil
	}
}

// WithBaseURL replaces the 'base' Url.
// Default value: 'http://127.0.0.1:9099/'.
// Use environment variable 'GLIDE_BASE_URL' to override.
func WithBaseURL(baseURL string) ClientOption {
	return func(client *Client) error {
		parsed, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		client.cfg.baseURL = parsed
		return nil
	}
}

// WithHttpClient replaces the 'HTTP' client.
// Default value: 'http.DefaultClient'.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(client *Client) error {
		client.cfg.httpClient = httpClient
		return nil
	}
}

// ApiKey returns the provided API key, empty string otherwise.
func (c *Client) ApiKey() string {
	return c.cfg.apiKey
}

// UserAgent returns the used 'User-Agent' header value.
func (c *Client) UserAgent() string {
	return c.cfg.userAgent
}

// BaseURL returns the used 'base url.URL'.
func (c *Client) BaseURL() url.URL {
	return *c.cfg.baseURL
}

// HttpClient returns the underlying http.Client.
func (c *Client) HttpClient() *http.Client {
	return c.cfg.httpClient
}

// Health returns true if the service is healthy.
func (c *Client) Health(ctx context.Context) (*bool, error) {
	type Health struct {
		Healthy bool `json:"healthy"`
	}

	req, err := c.cfg.Build(ctx, http.MethodGet, "/v1/health/", nil)
	if err != nil {
		return nil, err
	}

	var resp Health
	if _, err := c.cfg.Send(req, resp); err != nil {
		return nil, err
	}

	return &resp.Healthy, nil
}
