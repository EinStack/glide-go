package glide

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type config struct {
	apiKey     string
	userAgent  string
	baseURL    *url.URL
	httpClient *http.Client
}

// Build instantiates a new http.Request.
func (c *config) Build(ctx context.Context, method, path string, data any) (*http.Request, error) {
	abs, err := c.baseURL.Parse(path)
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
	req.Header.Set("User-Agent", c.userAgent)

	if len(c.apiKey) > 0 {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	return req, nil
}

// Send sends an http.Request and decodes http.Response into ret.
func (c *config) Send(r *http.Request, ret any) (*http.Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		if resp.Body == nil {
			return nil, NewError()
		}

		var errorResp *Error
		if err := json.NewDecoder(resp.Body).Decode(errorResp); err != nil {
			return nil, err
		}

		errorResp.Status = resp.StatusCode
		return nil, errorResp
	}

	if resp.StatusCode != http.StatusNoContent && ret != nil && resp.Body != nil {
		if err = json.NewDecoder(resp.Body).Decode(ret); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// Upgrade establishes the WebSocket connection.
func (c *config) Upgrade(ctx context.Context, path string) (*websocket.Conn, error) {
	wsBaseURL := c.baseURL
	if c.baseURL.Scheme == "https" {
		wsBaseURL.Scheme = "wss"
	} else if c.baseURL.Scheme == "http" {
		wsBaseURL.Scheme = "ws"
	}

	abs, err := wsBaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	if len(c.apiKey) > 0 {
		header.Set("Authorization", "Bearer "+c.apiKey)
	}

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, abs.String(), header)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
