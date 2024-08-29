package tg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const BaseURL = "https://api.telegram.org/bot"

type Client struct {
	log        *slog.Logger
	httpClient HTTPClientDoer
	baseURL    string
	token      string
}

type response struct {
	Success     bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

func New(token string, opts ...Option) *Client {
	c := &Client{token: token}

	defaultOpts := []Option{
		WithLogger(slog.New(slog.NewJSONHandler(io.Discard, nil))),
		WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
		WithBaseURL(BaseURL),
	}

	for _, opt := range append(defaultOpts, opts...) {
		opt(c)
	}

	return c
}

func (c *Client) Call(ctx context.Context, value Sendable) (json.RawMessage, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s%s/%s", c.baseURL, c.token, value.Method()),
		bytes.NewReader(value.Params()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusBadRequest:
	case http.StatusUnauthorized:
	case http.StatusConflict:
	default:
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	return c.parseResponse(resp.Body)
}

func (c *Client) parseResponse(r io.Reader) (json.RawMessage, error) {
	var resp response

	if err := json.NewDecoder(r).Decode(&resp); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}

	if !resp.Success {
		return nil, &Error{
			Code:    resp.ErrorCode,
			Message: resp.Description,
		}
	}

	return resp.Result, nil
}

func DecodeJSON[T any](payload []byte) (T, error) {
	var (
		val = new(T)
		err = json.Unmarshal(payload, val)
	)

	return *val, err
}
