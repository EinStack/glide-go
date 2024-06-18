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

// Health returns nil if the service is healthy.
func (c *Client) Health(ctx context.Context) error {
	req, err := c.cfg.Build(ctx, http.MethodGet, "/v1/health/", nil)
	if err != nil {
		return err
	}

	if _, err := c.cfg.Send(req, nil); err != nil {
		return err
	}

	return nil
}
