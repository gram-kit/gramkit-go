package gramkit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/gram-kit/gramkit-go/models"
)

func TestNew(t *testing.T) {
	bot, err := New("123456:ABC-DEF")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if bot == nil {
		t.Fatal("expected bot to be non-nil")
	}
}

func TestNewEmptyToken(t *testing.T) {
	bot, err := New("")
	if err == nil {
		t.Fatal("expected error for empty token")
	}
	if bot != nil {
		t.Fatal("expected bot to be nil on error")
	}
	if !strings.Contains(err.Error(), "token is empty") {
		t.Fatalf("unexpected error message: %v", err)
	}
}

func TestWithPollTimeout(t *testing.T) {
	bot, err := New("123456:ABC-DEF", WithPollTimeout(60))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if bot.pollTimeout != 60 {
		t.Fatalf("expected pollTimeout 60, got %d", bot.pollTimeout)
	}
}

func TestWithDefaultHandler(t *testing.T) {
	var called atomic.Int32

	handler := func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	}

	bot, err := New("123456:ABC-DEF", WithDefaultHandler(handler))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Send an update via the webhook handler to trigger the default handler.
	update := models.Update{
		UpdateID: 1,
		Message:  &models.Message{Text: "hello"},
	}
	body, _ := json.Marshal(update)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	bot.WebhookHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if called.Load() != 1 {
		t.Fatalf("expected handler to be called once, got %d", called.Load())
	}
}

func TestWithHTTPClient(t *testing.T) {
	custom := &http.Client{}
	bot, err := New("123456:ABC-DEF", WithHTTPClient(custom))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if bot.client.HTTPClient != custom {
		t.Fatal("expected custom HTTP client to be set")
	}
}

func TestWithServerURL(t *testing.T) {
	bot, err := New("123456:ABC-DEF", WithServerURL("http://localhost:8081"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := "http://localhost:8081/bot123456:ABC-DEF"
	if bot.client.BaseURL != expected {
		t.Fatalf("expected BaseURL %q, got %q", expected, bot.client.BaseURL)
	}
}

func TestDefaultPollTimeout(t *testing.T) {
	bot, err := New("123456:ABC-DEF")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if bot.pollTimeout != 30 {
		t.Fatalf("expected default pollTimeout 30, got %d", bot.pollTimeout)
	}
}
