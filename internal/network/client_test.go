package network

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c := NewClient("token123", nil)
	if c.Token != "token123" {
		t.Fatalf("expected token 'token123', got %q", c.Token)
	}
	if c.HTTPClient == nil {
		t.Fatal("expected default HTTP client to be set")
	}
	expected := "https://api.telegram.org/bottoken123"
	if c.BaseURL != expected {
		t.Fatalf("expected BaseURL %q, got %q", expected, c.BaseURL)
	}
}

func TestNewClientCustomHTTP(t *testing.T) {
	custom := &http.Client{Timeout: 60 * time.Second}
	c := NewClient("token123", custom)
	if c.HTTPClient != custom {
		t.Fatal("expected custom HTTP client")
	}
}

func TestExecute(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"ok":     true,
			"result": map[string]interface{}{"id": 12345, "is_bot": true, "first_name": "TestBot"},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	c := &Client{
		Token:      "test",
		BaseURL:    srv.URL,
		HTTPClient: srv.Client(),
	}

	var result struct {
		ID        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
	}

	err := c.Execute(context.Background(), "getMe", nil, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ID != 12345 {
		t.Fatalf("expected ID 12345, got %d", result.ID)
	}
	if !result.IsBot {
		t.Fatal("expected is_bot true")
	}
	if result.FirstName != "TestBot" {
		t.Fatalf("expected first_name 'TestBot', got %q", result.FirstName)
	}
}

func TestExecuteAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"ok":          false,
			"description": "Unauthorized",
			"error_code":  401,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	c := &Client{
		Token:      "bad-token",
		BaseURL:    srv.URL,
		HTTPClient: srv.Client(),
	}

	var result json.RawMessage
	err := c.Execute(context.Background(), "getMe", nil, &result)
	if err == nil {
		t.Fatal("expected error for API error response")
	}
	if got := err.Error(); got == "" {
		t.Fatal("expected non-empty error message")
	}
}

func TestExecuteContextCancellation(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response — context should cancel before this returns.
		<-r.Context().Done()
	}))
	defer srv.Close()

	c := &Client{
		Token:      "test",
		BaseURL:    srv.URL,
		HTTPClient: srv.Client(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately.

	var result json.RawMessage
	err := c.Execute(ctx, "getMe", nil, &result)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
}

func TestExecuteWithParams(t *testing.T) {
	var receivedBody map[string]interface{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewDecoder(r.Body).Decode(&receivedBody)
		resp := map[string]interface{}{"ok": true, "result": true}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	c := &Client{
		Token:      "test",
		BaseURL:    srv.URL,
		HTTPClient: srv.Client(),
	}

	params := map[string]interface{}{
		"chat_id": 123,
		"text":    "hello",
	}

	var result bool
	err := c.Execute(context.Background(), "sendMessage", params, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Fatal("expected result to be true")
	}
	if receivedBody["text"] != "hello" {
		t.Fatalf("expected text 'hello', got %v", receivedBody["text"])
	}
}

func TestUseMethod(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"ok":     true,
			"result": map[string]interface{}{"id": 42, "is_bot": true, "first_name": "Bot"},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	c := &Client{
		Token:      "test",
		BaseURL:    srv.URL,
		HTTPClient: srv.Client(),
	}

	type User struct {
		ID        int    `json:"id"`
		IsBot     bool   `json:"is_bot"`
		FirstName string `json:"first_name"`
	}

	result, err := UseMethod[User](c, context.Background(), "getMe", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ID != 42 {
		t.Fatalf("expected ID 42, got %d", result.ID)
	}
}
