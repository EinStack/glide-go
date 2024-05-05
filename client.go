package glide

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
// Default value: 'Glide/0.1.0 (Go; Ver. 1.22.2)'.
// Env variable: 'GLIDE_USER_AGENT'.
func WithUserAgent(userAgent string) ClientOption {
	return func(client *Client) error {
		client.UserAgent = userAgent
		return nil
	}
}

// WithBaseURL replaces the 'base' Url.
// Default value: 'http://127.0.0.1:9099/'.
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

// Build instantiates a new http.Request.
func (c *Client) Build(ctx context.Context, method, path string, data any) (*http.Request, error) {
	abs, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, abs.String(), nil)
	if err != nil {
		return nil, err
	}

	if data != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(data); err != nil {
			return nil, err
		}

		req.Body = io.NopCloser(buf)
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	return req, nil
}

// Send sends an http.Request and decodes http.Response into ret.
func (c *Client) Send(r *http.Request, ret any) (*http.Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > http.StatusBadRequest {
		// TODO: Decode into ErrorResponse.
		reason := fmt.Sprintf("status code: %d", resp.StatusCode)
		return nil, errors.New(reason)
	}

	if resp.StatusCode != http.StatusNoContent && ret != nil && resp.Body != nil {
		if err = json.NewDecoder(resp.Body).Decode(ret); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// Health returns nil if the service is healthy.
func (c *Client) Health(ctx context.Context) error {
	req, err := c.Build(ctx, http.MethodGet, "/v1/health/", nil)
	if err != nil {
		return err
	}

	if _, err := c.Send(req, nil); err != nil {
		return err
	}

	return nil
}
