package network

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gram-kit/gramkit-go/models"
)

type Client struct {
	Token      string
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(token string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second}
	}
	return &Client{
		Token:      token,
		BaseURL:    fmt.Sprintf("https://api.telegram.org/bot%s", token),
		HTTPClient: httpClient,
	}
}

func (c *Client) Execute(ctx context.Context, method string, params interface{}, v interface{}) error {
	url := fmt.Sprintf("%s/%s", c.BaseURL, method)

	body, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp struct {
		models.Response[json.RawMessage]
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !apiResp.Ok {
		return fmt.Errorf("api error: %s (code %d)", apiResp.Description, apiResp.ErrorCode)
	}

	if v != nil {
		if err := json.Unmarshal(apiResp.Result, v); err != nil {
			return fmt.Errorf("failed to unmarshal result: %w", err)
		}
	}

	return nil
}

func UseMethod[T any](c *Client, ctx context.Context, method string, params any) (T, error) {
	var result T
	if err := c.Execute(ctx, method, params, &result); err != nil {
		return result, err
	}
	return result, nil
}
