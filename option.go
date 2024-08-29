package tg

import (
	"log/slog"
	"net/http"
)

type Option func(*Client)

type HTTPClientDoer interface {
	Do(*http.Request) (*http.Response, error)
}

func WithHTTPClient(httpClient HTTPClientDoer) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		if baseURL == "" {
			baseURL = BaseURL
		}

		c.baseURL = baseURL
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		c.log = logger
	}
}
