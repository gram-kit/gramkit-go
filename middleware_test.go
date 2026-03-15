package gramkit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/gram-kit/gramkit-go/models"
)

func sendUpdate(t *testing.T, bot *Bot, update models.Update) {
	t.Helper()
	body, _ := json.Marshal(update)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	bot.WebhookHandler().ServeHTTP(rec, req)
}

func TestMiddlewareCallsNext(t *testing.T) {
	var handlerCalled atomic.Int32
	var mwCalled atomic.Int32

	bot, _ := New("test-token",
		WithDefaultHandler(func(ctx context.Context, b *Bot, update *models.Update) {
			handlerCalled.Add(1)
		}),
		WithMiddleware(func(ctx context.Context, b *Bot, update *models.Update, next HandlerFunc) {
			mwCalled.Add(1)
			next(ctx, b, update)
		}),
	)

	sendUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "hi"}})

	if mwCalled.Load() != 1 {
		t.Fatalf("expected middleware called once, got %d", mwCalled.Load())
	}
	if handlerCalled.Load() != 1 {
		t.Fatalf("expected handler called once, got %d", handlerCalled.Load())
	}
}

func TestMiddlewareBlocksNext(t *testing.T) {
	var handlerCalled atomic.Int32

	bot, _ := New("test-token",
		WithDefaultHandler(func(ctx context.Context, b *Bot, update *models.Update) {
			handlerCalled.Add(1)
		}),
		WithMiddleware(func(ctx context.Context, b *Bot, update *models.Update, next HandlerFunc) {
			// Intentionally do NOT call next.
		}),
	)

	sendUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "hi"}})

	if handlerCalled.Load() != 0 {
		t.Fatal("expected handler NOT to be called when middleware blocks")
	}
}

func TestMiddlewareOrder(t *testing.T) {
	var mu sync.Mutex
	var order []int

	appendOrder := func(n int) {
		mu.Lock()
		order = append(order, n)
		mu.Unlock()
	}

	mw1 := func(ctx context.Context, b *Bot, update *models.Update, next HandlerFunc) {
		appendOrder(1)
		next(ctx, b, update)
		appendOrder(4)
	}
	mw2 := func(ctx context.Context, b *Bot, update *models.Update, next HandlerFunc) {
		appendOrder(2)
		next(ctx, b, update)
		appendOrder(3)
	}

	bot, _ := New("test-token",
		WithDefaultHandler(func(ctx context.Context, b *Bot, update *models.Update) {
			appendOrder(0) // innermost
		}),
		WithMiddleware(mw1, mw2),
	)

	sendUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "hi"}})

	mu.Lock()
	defer mu.Unlock()

	expected := []int{1, 2, 0, 3, 4}
	if len(order) != len(expected) {
		t.Fatalf("expected %d calls, got %d: %v", len(expected), len(order), order)
	}
	for i, v := range expected {
		if order[i] != v {
			t.Fatalf("expected order %v, got %v", expected, order)
		}
	}
}

func TestUseAddsMiddleware(t *testing.T) {
	var mwCalled atomic.Int32
	var handlerCalled atomic.Int32

	bot, _ := New("test-token",
		WithDefaultHandler(func(ctx context.Context, b *Bot, update *models.Update) {
			handlerCalled.Add(1)
		}),
	)

	// Add middleware after creation via Use().
	bot.Use(func(ctx context.Context, b *Bot, update *models.Update, next HandlerFunc) {
		mwCalled.Add(1)
		next(ctx, b, update)
	})

	sendUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "hi"}})

	if mwCalled.Load() != 1 {
		t.Fatalf("expected Use middleware called, got %d", mwCalled.Load())
	}
	if handlerCalled.Load() != 1 {
		t.Fatalf("expected handler called, got %d", handlerCalled.Load())
	}
}
