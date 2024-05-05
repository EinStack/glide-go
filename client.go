package glide

import (
	"context"
	"net/http"
	"net/url"
)

// Client is a minimal 'Glide' client.
type Client struct {
	ApiKey     string
	UserAgent  string
	BaseURL    *url.URL
	httpClient *http.Client

	Language LanguageSvc
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

	client := &Client{}
	client.Language = &language{client}

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// WithApiKey replaces the api key.
// Default value: 'development'.
// Env variable 'GLIDE_API_KEY'.
func WithApiKey(apiKey string) ClientOption {
	return func(client *Client) error {
		client.ApiKey = apiKey
		return nil
	}
}

// WithUserAgent replaces the 'User-Agent' header.
// Default value: 'glide-go/0.1.0'.
// Env variable: 'GLIDE_USER_AGENT'.
func WithUserAgent(userAgent string) ClientOption {
	return func(client *Client) error {
		client.UserAgent = userAgent
		return nil
	}
}

// WithBaseURL replaces the 'base' Url.
// Default value: 'https://api.einstack.com'.
// Env variable: 'GLIDE_BASE_URL'.
func WithBaseURL(baseURL string) ClientOption {
	return func(client *Client) error {
		parsed, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		client.BaseURL = parsed
		return nil
	}
}

// WithHttpClient replaces the 'HTTP' client.
// Default value: 'http.DefaultClient'.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = httpClient
		return nil
	}
}

// Health returns nil if the service is healthy.
func (c *Client) Health(ctx context.Context) error {
	// TODO.
	return nil
}
