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

// postUpdate is a helper that POSTs an update to the bot's webhook handler.
func postUpdate(t *testing.T, bot *Bot, update models.Update) *httptest.ResponseRecorder {
	t.Helper()
	body, err := json.Marshal(update)
	if err != nil {
		t.Fatalf("failed to marshal update: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	bot.WebhookHandler().ServeHTTP(rec, req)
	return rec
}

func TestRegisterHandler(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandler(OnMessage, "/start", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID: 1,
		Message:  &models.Message{Text: "/start"},
	})

	if called.Load() != 1 {
		t.Fatalf("expected handler called once, got %d", called.Load())
	}
}

func TestRegisterHandlerNoMatch(t *testing.T) {
	var handlerCalled atomic.Int32
	var defaultCalled atomic.Int32

	bot, _ := New("test-token", WithDefaultHandler(func(ctx context.Context, b *Bot, update *models.Update) {
		defaultCalled.Add(1)
	}))

	bot.RegisterHandler(OnMessage, "/start", func(ctx context.Context, b *Bot, update *models.Update) {
		handlerCalled.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID: 1,
		Message:  &models.Message{Text: "/help"},
	})

	if handlerCalled.Load() != 0 {
		t.Fatal("expected /start handler NOT to be called")
	}
	if defaultCalled.Load() != 1 {
		t.Fatalf("expected default handler called once, got %d", defaultCalled.Load())
	}
}

func TestMatchExact(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnMessage, MatchExact, "/start", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	// Exact match should work.
	postUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "/start"}})
	if called.Load() != 1 {
		t.Fatalf("expected exact match to fire, got %d calls", called.Load())
	}

	// Partial should NOT match exact.
	postUpdate(t, bot, models.Update{UpdateID: 2, Message: &models.Message{Text: "/start hello"}})
	if called.Load() != 1 {
		t.Fatalf("expected exact match NOT to fire for partial, got %d calls", called.Load())
	}
}

func TestMatchPrefix(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnMessage, MatchPrefix, "/start", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	// "/start hello" should match prefix.
	postUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "/start hello"}})
	if called.Load() != 1 {
		t.Fatalf("expected prefix match for '/start hello', got %d", called.Load())
	}

	// "/start@bot" should match prefix.
	postUpdate(t, bot, models.Update{UpdateID: 2, Message: &models.Message{Text: "/start@bot"}})
	if called.Load() != 2 {
		t.Fatalf("expected prefix match for '/start@bot', got %d", called.Load())
	}

	// exact "/start" should also match prefix.
	postUpdate(t, bot, models.Update{UpdateID: 3, Message: &models.Message{Text: "/start"}})
	if called.Load() != 3 {
		t.Fatalf("expected prefix match for exact '/start', got %d", called.Load())
	}
}

func TestMatchContains(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnMessage, MatchContains, "hello", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "say hello world"}})
	if called.Load() != 1 {
		t.Fatalf("expected contains match, got %d", called.Load())
	}

	postUpdate(t, bot, models.Update{UpdateID: 2, Message: &models.Message{Text: "goodbye"}})
	if called.Load() != 1 {
		t.Fatalf("expected no contains match for 'goodbye', got %d", called.Load())
	}
}

func TestMatchWildcard(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnMessage, MatchExact, "*", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "anything at all"}})
	if called.Load() != 1 {
		t.Fatalf("expected wildcard to match, got %d", called.Load())
	}

	postUpdate(t, bot, models.Update{UpdateID: 2, Message: &models.Message{Text: "/start"}})
	if called.Load() != 2 {
		t.Fatalf("expected wildcard to match /start, got %d", called.Load())
	}
}

func TestHandlerPriority(t *testing.T) {
	var first atomic.Int32
	var second atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandler(OnMessage, "/start", func(ctx context.Context, b *Bot, update *models.Update) {
		first.Add(1)
	})
	bot.RegisterHandler(OnMessage, "/start", func(ctx context.Context, b *Bot, update *models.Update) {
		second.Add(1)
	})

	postUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "/start"}})

	if first.Load() != 1 {
		t.Fatal("expected first handler to win")
	}
	if second.Load() != 0 {
		t.Fatal("expected second handler NOT to be called")
	}
}

func TestOnCallbackQuery(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandler(OnCallbackQuery, "btn_yes", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID:      1,
		CallbackQuery: &models.CallbackQuery{ID: "1", Data: "btn_yes"},
	})

	if called.Load() != 1 {
		t.Fatalf("expected callback query handler called, got %d", called.Load())
	}
}

func TestOnMessage(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnMessage, MatchExact, "*", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	// Message update should match.
	postUpdate(t, bot, models.Update{UpdateID: 1, Message: &models.Message{Text: "hi"}})
	if called.Load() != 1 {
		t.Fatalf("expected OnMessage handler, got %d", called.Load())
	}

	// CallbackQuery update should NOT match OnMessage.
	postUpdate(t, bot, models.Update{UpdateID: 2, CallbackQuery: &models.CallbackQuery{ID: "1", Data: "x"}})
	if called.Load() != 1 {
		t.Fatalf("expected OnMessage handler NOT to fire for callback, got %d", called.Load())
	}
}

func TestOnPoll(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnPoll, MatchExact, "*", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID: 1,
		Poll:     &models.Poll{ID: "poll1", Question: "?"},
	})

	if called.Load() != 1 {
		t.Fatalf("expected OnPoll handler, got %d", called.Load())
	}
}

func TestOnChatJoinRequest(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnChatJoinRequest, MatchExact, "*", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID:        1,
		ChatJoinRequest: &models.ChatJoinRequest{Chat: models.Chat{ID: 1}, From: models.User{ID: 2}},
	})

	if called.Load() != 1 {
		t.Fatalf("expected OnChatJoinRequest handler, got %d", called.Load())
	}
}

func TestOnInlineQuery(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnInlineQuery, MatchContains, "search", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID:    1,
		InlineQuery: &models.InlineQuery{ID: "1", Query: "search term"},
	})

	if called.Load() != 1 {
		t.Fatalf("expected OnInlineQuery handler, got %d", called.Load())
	}
}

func TestOnShippingQuery(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnShippingQuery, MatchExact, "*", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID:      1,
		ShippingQuery: &models.ShippingQuery{ID: "sq1"},
	})

	if called.Load() != 1 {
		t.Fatalf("expected OnShippingQuery handler, got %d", called.Load())
	}
}

func TestOnChatBoost(t *testing.T) {
	var called atomic.Int32

	bot, _ := New("test-token")
	bot.RegisterHandlerMatch(OnChatBoost, MatchExact, "*", func(ctx context.Context, b *Bot, update *models.Update) {
		called.Add(1)
	})

	postUpdate(t, bot, models.Update{
		UpdateID:  1,
		ChatBoost: &models.ChatBoostUpdated{Chat: models.Chat{ID: 1}},
	})

	if called.Load() != 1 {
		t.Fatalf("expected OnChatBoost handler, got %d", called.Load())
	}
}

func TestNoHandlerNoDefaultNoPanic(t *testing.T) {
	bot, _ := New("test-token")

	// No handlers, no default handler — should not panic.
	postUpdate(t, bot, models.Update{
		UpdateID: 1,
		Message:  &models.Message{Text: "hello"},
	})
}
